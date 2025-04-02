package control

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"connectrpc.com/connect"
	pb "github.com/pixelfs/pixelfs/gen/pixelfs/v1"
	"github.com/pixelfs/pixelfs/util"
)

func (p *PixelFS) Upload(ctx *pb.FileContext, input string) (err error) {
	location, err := p.Core.LocationService.GetLocationByContext(
		context.Background(),
		connect.NewRequest(&pb.GetLocationByContextRequest{
			Context: ctx,
		}),
	)
	if err != nil {
		return err
	}

	_, fileName := filepath.Split(ctx.Path)
	if input == "" {
		dir, err := os.Getwd()
		if err != nil {
			return err
		}

		input = filepath.Join(dir, fileName)
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

	if fileInfo.IsDir() {
		if _, err := p.Core.FileSystemService.Mkdir(
			context.Background(),
			connect.NewRequest(&pb.FileMkdirRequest{
				Context: ctx,
			}),
		); err != nil {
			return err
		}

		files, err := file.Readdir(0)
		if err != nil {
			return err
		}

		for _, file := range files {
			if err := p.Upload(
				&pb.FileContext{
					NodeId:   ctx.NodeId,
					Location: ctx.Location,
					Path:     filepath.Join(ctx.Path, file.Name()),
				},
				filepath.Join(input, file.Name()),
			); err != nil {
				return err
			}
		}

		return nil
	}

	fileDir, fileName := filepath.Split(ctx.Path)
	ctx.Path = filepath.ToSlash(filepath.Join(fileDir, tmpPrefix+fileName))

	defer func() {
		if err == nil {
			_, err = p.Core.FileSystemService.Move(
				context.Background(),
				connect.NewRequest(&pb.FileMoveRequest{
					Src: ctx,
					Dest: &pb.FileContext{
						NodeId:   ctx.NodeId,
						Location: ctx.Location,
						Path:     filepath.ToSlash(filepath.Join(fileDir, fileName)),
					},
				}),
			)
		} else {
			_ = p.Rm(ctx, false)
		}
	}()

	hash, err := util.GetFileHash(input)
	if err != nil {
		return err
	}

	blockCount := fileInfo.Size() / location.Msg.Location.BlockSize
	bar := util.NewProgressBar(int(blockCount+1), fmt.Sprintf("Uploading %s", ctx.Path))

	if err = bar.RenderBlank(); err != nil {
		return err
	}

	for index := int64(0); index <= blockCount; index++ {
		offset := index * location.Msg.Location.BlockSize
		if offset >= fileInfo.Size() {
			offset = fileInfo.Size() - 1
		}

		blockSize := location.Msg.Location.BlockSize
		if offset+blockSize > fileInfo.Size() {
			blockSize = fileInfo.Size() - offset
		}

		storage, err := p.Core.StorageService.Upload(
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

		_, err = p.Core.FileSystemService.Write(
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

		bar.Add(1)
	}

	return nil
}
