package control

import (
	"context"
	"fmt"
	"path/filepath"
	"time"

	"connectrpc.com/connect"
	pb "github.com/pixelfs/pixelfs/gen/pixelfs/v1"
	"github.com/pixelfs/pixelfs/util"
)

func (p *PixelFS) Cp(src *pb.FileContext, dest *pb.FileContext) error {
	if src.NodeId == dest.NodeId {
		_, err := p.Core.FileSystemService.Copy(
			context.Background(),
			connect.NewRequest(&pb.FileCopyRequest{
				Src:  src,
				Dest: dest,
			}),
		)

		return err
	}

	stat, err := p.Core.FileSystemService.Stat(
		context.Background(),
		connect.NewRequest(&pb.FileStatRequest{
			Context: src,
			Hash:    true,
		}),
	)
	if err != nil {
		return err
	}

	if stat.Msg.File.Type == pb.FileType_DIR {
		_, err := p.Core.FileSystemService.Mkdir(
			context.Background(),
			connect.NewRequest(&pb.FileMkdirRequest{
				Context: dest,
			}),
		)
		if err != nil {
			return err
		}

		list, err := p.Core.FileSystemService.List(
			context.Background(),
			connect.NewRequest(&pb.FileListRequest{
				Context: src,
			}),
		)
		if err != nil {
			return err
		}

		for _, file := range list.Msg.Files {
			if err := p.Cp(
				&pb.FileContext{
					NodeId:   src.NodeId,
					Location: src.Location,
					Path:     filepath.Join(src.Path, file.Name),
				},
				&pb.FileContext{
					NodeId:   dest.NodeId,
					Location: dest.Location,
					Path:     filepath.Join(dest.Path, file.Name),
				},
			); err != nil {
				return err
			}
		}

		return nil
	}

	location, err := p.Core.LocationService.GetLocationByContext(
		context.Background(),
		connect.NewRequest(&pb.GetLocationByContextRequest{
			Context: src,
		}),
	)
	if err != nil {
		return err
	}

	blockCount := stat.Msg.File.Size / location.Msg.Location.BlockSize
	bar := util.NewProgressBar(int(blockCount+1), fmt.Sprintf("Copying %s", src.Path))
	if err = bar.RenderBlank(); err != nil {
		return err
	}

	for index := int64(0); index <= blockCount; index++ {
		var read *connect.Response[pb.FileReadResponse]
		for retries := 0; retries < 20; retries++ {
			read, err = p.Core.FileSystemService.Read(
				context.Background(),
				connect.NewRequest(&pb.FileReadRequest{
					Context:    src,
					BlockType:  pb.BlockType_SIZE,
					BlockIndex: index,
				}),
			)
			if err != nil {
				return err
			}

			if read.Msg.BlockStatus != pb.BlockStatus_PENDING {
				break
			}

			time.Sleep(5 * time.Second)
		}

		if read == nil || read.Msg.BlockStatus == pb.BlockStatus_PENDING {
			return fmt.Errorf("block %d is still pending after retries", index)
		}

		_, err := p.Core.FileSystemService.Write(
			context.Background(),
			connect.NewRequest(&pb.FileWriteRequest{
				Context:    dest,
				Hash:       stat.Msg.File.Hash,
				BlockType:  pb.BlockType_SIZE,
				BlockIndex: index,
				Offset:     index * location.Msg.Location.BlockSize,
				Url:        read.Msg.Url,
			}),
		)
		if err != nil {
			return err
		}

		bar.Add(1)
	}

	return nil
}
