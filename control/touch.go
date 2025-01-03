package control

import (
	"context"

	"connectrpc.com/connect"
	pb "github.com/pixelfs/pixelfs/gen/pixelfs/v1"
)

func (p *PixelFS) Touch(ctx *pb.FileContext) error {
	_, err := p.Core.FileSystemService.Create(
		context.Background(),
		connect.NewRequest(&pb.FileCreateRequest{
			Context: ctx,
		}),
	)
	if err != nil {
		return err
	}

	return nil
}
