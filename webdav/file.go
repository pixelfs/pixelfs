package webdav

import (
	"bytes"
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"connectrpc.com/connect"
	"github.com/pixelfs/pixelfs/config"
	pb "github.com/pixelfs/pixelfs/gen/pixelfs/v1"
	"github.com/pixelfs/pixelfs/rpc/core"
	"github.com/pixelfs/pixelfs/util"
	"golang.org/x/net/webdav"
)

type fileInfo struct {
	os.FileInfo
	file *pb.File
}

func newFileInfo(file *pb.File) *fileInfo {
	return &fileInfo{file: file}
}

func (fi *fileInfo) Name() string {
	return fi.file.Name
}

func (fi *fileInfo) Size() int64 {
	return fi.file.Size
}

func (fi *fileInfo) Mode() os.FileMode {
	return os.FileMode(fi.file.Perm)
}

func (fi *fileInfo) ModTime() time.Time {
	return fi.file.ModifiedAt.AsTime()
}

func (fi *fileInfo) IsDir() bool {
	return fi.file.Type == pb.FileType_NODE || fi.file.Type == pb.FileType_LOCATION || fi.file.Type == pb.FileType_DIR
}

func (fi *fileInfo) Sys() any {
	return fi.file
}

type file struct {
	webdav.File
	request  *http.Request
	cfg      *config.Config
	core     *core.GrpcV1Client
	context  *pb.FileContext
	location *pb.Location
	info     *pb.File
	offset   int64
	lock     sync.Map
}

func newFile(request *http.Request, cfg *config.Config, core *core.GrpcV1Client, context *pb.FileContext, info *pb.File, location *pb.Location) *file {
	return &file{request: request, cfg: cfg, core: core, context: context, info: info, location: location}
}

func (f *file) Readdir(count int) ([]os.FileInfo, error) {
	response, err := f.core.FileSystemService.List(
		context.Background(),
		connect.NewRequest(&pb.FileListRequest{
			Context: f.context,
		}),
	)
	if err != nil {
		return nil, err
	}

	files := make([]os.FileInfo, 0, len(response.Msg.Files))
	for _, file := range response.Msg.Files {
		files = append(files, newFileInfo(file))
	}

	return files, nil
}

func (f *file) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.offset = offset
	case io.SeekCurrent:
		f.offset += offset
	case io.SeekEnd:
		f.offset = f.info.Size + offset
	}

	if f.offset < 0 {
		return 0, fmt.Errorf("invalid offset")
	}

	return f.offset, nil
}

func (f *file) Stat() (os.FileInfo, error) {
	return newFileInfo(f.info), nil
}

func (f *file) Read(p []byte) (n int, err error) {
	length := int64(len(p))
	remaining := f.info.Size - f.offset

	if remaining <= 0 {
		return 0, io.EOF
	}

	if length > remaining {
		p = p[:remaining]
	}

	blockStart := f.offset / f.location.BlockSize
	blockEnd := (f.offset + length - 1) / f.location.BlockSize

	var buffer bytes.Buffer
	for bloxkIndex := blockStart; bloxkIndex <= blockEnd; bloxkIndex++ {
		data, err := f.read(length, blockStart, blockEnd, bloxkIndex, f.location.BlockSize)
		if err != nil {
			return 0, err
		}

		buffer.Write(data)
	}

	n = copy(p, buffer.Bytes())
	f.offset += int64(n)

	return n, nil
}

func (f *file) Write(p []byte) (n int, err error) {
	if f.request.ContentLength == 0 {
		return 0, nil
	}

	tmpDir := filepath.Join(os.TempDir(), "pixelfs", "webdav")
	cacheFile := filepath.Join(tmpDir, fmt.Sprintf("%x", md5.Sum([]byte(f.request.URL.Path))))
	_ = util.EnsureDir(tmpDir)

	file, err := os.OpenFile(cacheFile, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	if _, err := file.Seek(f.offset, io.SeekStart); err != nil {
		return 0, err
	}

	n, err = file.Write(p)
	if err != nil {
		return 0, err
	}

	f.offset += int64(n)
	if f.offset >= f.request.ContentLength {
		if err := f.write(cacheFile); err != nil {
			return 0, err
		}
	}

	return n, nil
}

func (f *file) Close() error {
	return nil
}

func (f *file) read(length, blockStart, blockEnd, blockIndex, blockSize int64) ([]byte, error) {
	f.lockBlock(blockIndex, blockSize)
	defer f.unlockBlock(blockIndex, blockSize)

	if f.info.Hash == "" {
		return nil, errors.New("file hash is empty")
	}

	cacheFile := filepath.Join(f.cfg.Webdav.Cache.Path, f.info.Hash, fmt.Sprintf("%d_%d", blockSize/1000, blockIndex))
	if _, err := os.Stat(cacheFile); err != nil {
		if os.IsNotExist(err) {
			var read *connect.Response[pb.FileReadResponse]
			for retries := 0; retries < 20; retries++ {
				read, err = f.core.FileSystemService.Read(
					context.Background(),
					connect.NewRequest(&pb.FileReadRequest{
						Context:    f.context,
						BlockType:  pb.BlockType_SIZE,
						BlockIndex: blockIndex,
					}),
				)
				if err != nil {
					return nil, err
				}

				if read.Msg.BlockStatus != pb.BlockStatus_PENDING {
					break
				}
				time.Sleep(5 * time.Second)
			}

			if read == nil || read.Msg.BlockStatus == pb.BlockStatus_PENDING {
				return nil, fmt.Errorf("block %d is still pending after retries", blockIndex)
			}

			_, err = util.Resty.R().SetOutput(cacheFile).Get(read.Msg.Url)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	file, err := os.Open(cacheFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	startOffset := int64(0)
	if blockIndex == blockStart {
		startOffset = f.offset % blockSize
	}

	endOffset := blockSize
	if blockIndex == blockEnd {
		endOffset = (f.offset + length) % blockSize

		// if length is zero or aligns with block boundaries, treat it as full block size
		if endOffset == 0 {
			endOffset = blockSize
		}
	}
	readLength := endOffset - startOffset

	if _, err := file.Seek(startOffset, io.SeekStart); err != nil {
		return nil, err
	}

	buf := make([]byte, readLength)
	n, err := file.Read(buf)
	if err != nil && err != io.EOF {
		return nil, err
	}

	return buf[:n], nil
}

func (f *file) write(input string) error {
	file, err := os.Open(input)
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	hash, err := util.GetFileHash(input)
	if err != nil {
		return err
	}

	blockCount := fileInfo.Size() / f.location.BlockSize
	for index := int64(0); index <= blockCount; index++ {
		offset := index * f.location.BlockSize
		if offset >= fileInfo.Size() {
			offset = fileInfo.Size() - 1
		}

		blockSize := f.location.BlockSize
		if offset+blockSize > fileInfo.Size() {
			blockSize = fileInfo.Size() - offset
		}

		storage, err := f.core.StorageService.Upload(
			context.Background(),
			connect.NewRequest(&pb.StorageUploadRequest{
				Context:    f.context,
				Hash:       hash,
				BlockType:  pb.BlockType_SIZE,
				BlockIndex: index,
				BlockSize:  blockSize,
			}),
		)
		if err != nil {
			return err
		}

		buffer := make([]byte, blockSize)
		if _, err = file.ReadAt(buffer, offset); err != nil && err != io.EOF {
			return err
		}

		if _, err = util.Resty.R().SetBody(buffer).Put(storage.Msg.Url); err != nil {
			return err
		}

		_, err = f.core.FileSystemService.Write(
			context.Background(),
			connect.NewRequest(&pb.FileWriteRequest{
				Context:    f.context,
				Hash:       hash,
				BlockType:  pb.BlockType_SIZE,
				BlockIndex: index,
			}),
		)
		if err != nil {
			return err
		}
	}

	// remove tmp file
	if err := os.Remove(input); err != nil {
		return err
	}

	return nil
}

func (f *file) lockBlock(blockIndex, blockSize int64) {
	lockKey := fmt.Sprintf("%s:%d:%d", f.info.Hash, blockSize, blockIndex)
	val, _ := f.lock.LoadOrStore(lockKey, make(chan struct{}, 1))
	lock := val.(chan struct{})

	lock <- struct{}{}
}

func (f *file) unlockBlock(blockIndex, blockSize int64) {
	lockKey := fmt.Sprintf("%s:%d:%d", f.info.Hash, blockSize, blockIndex)
	val, ok := f.lock.Load(lockKey)
	if ok {
		lock := val.(chan struct{})
		<-lock
	}
}
