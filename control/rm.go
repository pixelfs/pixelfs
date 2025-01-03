package control

import (
	"context"

	"connectrpc.com/connect"
	pb "github.com/pixelfs/pixelfs/gen/pixelfs/v1"
)

func (p *PixelFS) Rm(ctx *pb.FileContext, recursive bool) error {
	_, err := p.Core.FileSystemService.Remove(
		context.Background(),
		connect.NewRequest(&pb.FileRemoveRequest{
			Context:   ctx,
			Recursive: recursive,
		}),
	)
	if err != nil {
		return err
	}

	return nil
}
