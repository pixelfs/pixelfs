package control

import (
	"context"

	"connectrpc.com/connect"
	pb "github.com/pixelfs/pixelfs/gen/pixelfs/v1"
)

func (p *PixelFS) Mkdir(ctx *pb.FileContext) error {
	_, err := p.Core.FileSystemService.Mkdir(
		context.Background(),
		connect.NewRequest(&pb.FileMkdirRequest{
			Context: ctx,
		}),
	)
	if err != nil {
		return err
	}

	return nil
}
