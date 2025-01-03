package control

import (
	"context"

	"connectrpc.com/connect"
	pb "github.com/pixelfs/pixelfs/gen/pixelfs/v1"
)

func (p *PixelFS) Mv(src *pb.FileContext, dest *pb.FileContext) error {
	if src.NodeId == dest.NodeId {
		_, err := p.Core.FileSystemService.Move(
			context.Background(),
			connect.NewRequest(&pb.FileMoveRequest{
				Src:  src,
				Dest: dest,
			}),
		)

		return err
	}

	if err := p.Cp(src, dest); err != nil {
		return err
	}
	if err := p.Rm(src, true); err != nil {
		return err
	}

	return nil
}
