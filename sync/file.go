package sync

import (
	"context"
	"io"
	"os"
	"runtime"
	"strings"
	"time"

	"connectrpc.com/connect"
	pb "github.com/pixelfs/pixelfs/gen/pixelfs/v1"
	"github.com/pixelfs/pixelfs/util"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (fs *FileSync) upload(input string, hash string, ctx *pb.FileContext, platform string) (err error) {
	if _, err = fs.Core.SyncService.LockCheck(
		context.Background(),
		connect.NewRequest(&pb.SyncLockCheckRequest{
			Context: ctx,
		}),
	); err != nil {
		return err
	}

	file, err := os.Open(input)
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	if fileInfo.Size() == 0 {
		_, err = fs.Core.FileSystemService.Create(
			context.Background(),
			connect.NewRequest(&pb.FileCreateRequest{
				Context: ctx,
			}),
		)

		return err
	}

	fileDir, fileName := util.SplitPath(platform, ctx.Path)
	ctx.Path = util.JoinPath(platform, fileDir, tmpPrefix+fileName)

	defer func() {
		if err == nil {
			if runtime.GOOS != "windows" && platform != "windows" {
				if err = fs.chmod(ctx, fileInfo.Mode()); err != nil {
					return
				}
			}

			if err = fs.chtimes(ctx, fileInfo.ModTime()); err != nil {
				return
			}
			_, err = fs.Core.FileSystemService.Move(
				context.Background(),
				connect.NewRequest(&pb.FileMoveRequest{
					Src: ctx,
					Dest: &pb.FileContext{
						NodeId:   ctx.NodeId,
						Location: ctx.Location,
						Path:     util.JoinPath(platform, fileDir, fileName),
					},
				}),
			)
		} else {
			_ = fs.remove(ctx)
		}
	}()

	location, err := fs.Core.LocationService.GetLocationByContext(
		context.Background(),
		connect.NewRequest(&pb.GetLocationByContextRequest{
			Context: ctx,
		}),
	)
	if err != nil {
		return err
	}

	blockCount := fileInfo.Size() / location.Msg.Location.BlockSize

	for index := int64(0); index <= blockCount; index++ {
		offset := index * location.Msg.Location.BlockSize
		if offset >= fileInfo.Size() {
			offset = fileInfo.Size() - 1
		}

		blockSize := location.Msg.Location.BlockSize
		if offset+blockSize > fileInfo.Size() {
			blockSize = fileInfo.Size() - offset
		}

		storage, err := fs.Core.StorageService.Upload(
			context.Background(),
			connect.NewRequest(&pb.StorageUploadRequest{
				Context:    ctx,
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

		_, err = fs.Core.FileSystemService.Write(
			context.Background(),
			connect.NewRequest(&pb.FileWriteRequest{
				Context:    ctx,
				Hash:       hash,
				BlockType:  pb.BlockType_SIZE,
				BlockIndex: index,
			}),
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func (fs *FileSync) stat(ctx *pb.FileContext) (*pb.File, error) {
	response, err := fs.Core.FileSystemService.Stat(
		context.Background(),
		connect.NewRequest(&pb.FileStatRequest{
			Context: ctx,
			Hash:    true,
		}),
	)

	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") ||
			strings.Contains(err.Error(), "The system cannot find the path specified") ||
			strings.Contains(err.Error(), "The system cannot find the file specified") {
			return nil, os.ErrNotExist
		}

		return nil, err
	}

	return response.Msg.File, nil
}

func (fs *FileSync) remove(ctx *pb.FileContext) error {
	_, err := fs.Core.FileSystemService.Remove(
		context.Background(),
		connect.NewRequest(&pb.FileRemoveRequest{
			Context:   ctx,
			Recursive: true,
		}),
	)

	return err
}

func (fs *FileSync) chmod(ctx *pb.FileContext, mode os.FileMode) error {
	_, err := fs.Core.FileSystemService.Chmod(
		context.Background(),
		connect.NewRequest(&pb.FileChmodRequest{
			Context: ctx,
			Mode:    uint32(mode),
		}),
	)

	return err
}

func (fs *FileSync) chtimes(ctx *pb.FileContext, mtime time.Time) error {
	_, err := fs.Core.FileSystemService.Chtimes(
		context.Background(),
		connect.NewRequest(&pb.FileChtimesRequest{
			Context: ctx,
			Atime:   timestamppb.Now(),
			Mtime:   timestamppb.New(mtime),
		}),
	)

	return err
}
