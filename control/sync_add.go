package control

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	pb "github.com/pixelfs/pixelfs/gen/pixelfs/v1"
)

func (p *PixelFS) SyncAdd(id, name string, enabled bool, config *pb.SyncConfig, srcContext, destContext *pb.FileContext) error {
	response, err := p.Core.SyncService.CreateSync(
		context.Background(),
		connect.NewRequest(&pb.SyncCreateRequest{
			Sync: &pb.Sync{
				Id:          id,
				Name:        name,
				Enabled:     enabled,
				SrcContext:  srcContext,
				DestContext: destContext,
				Config:      config,
			},
		}),
	)
	if err != nil {
		return err
	}

	targetFunc := p.SyncStart
	if !response.Msg.Sync.Enabled {
		targetFunc = p.SyncStop
	}

	if err = targetFunc(response.Msg.Sync.Id); err != nil {
		return err
	}

	fmt.Println("sync added successfully")
	return nil
}
