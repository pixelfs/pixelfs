package control

import (
	"context"

	"connectrpc.com/connect"
	pb "github.com/pixelfs/pixelfs/gen/pixelfs/v1"
)

func (p *PixelFS) Cd(ctx *pb.FileContext) error {
	_, err := p.Core.FileSystemService.Stat(
		context.Background(),
		connect.NewRequest(&pb.FileStatRequest{
			Context: ctx,
		}),
	)
	if err != nil {
		return err
	}

	return nil
}
