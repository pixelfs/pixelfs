package control

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	pb "github.com/pixelfs/pixelfs/gen/pixelfs/v1"
)

func (p *PixelFS) SyncRm(syncId string) error {
	response, err := p.Core.SyncService.GetSyncById(
		context.Background(),
		connect.NewRequest(&pb.SyncGetByIdRequest{
			SyncId: syncId,
		}),
	)
	if err != nil {
		return err
	}

	_ = p.syncStop(response.Msg.Sync.Id, response.Msg.Sync.SrcNodeId)
	_ = p.syncStop(response.Msg.Sync.Id, response.Msg.Sync.DestNodeId)

	_, err = p.Core.SyncService.RemoveSync(
		context.Background(),
		connect.NewRequest(&pb.SyncRemoveRequest{
			SyncId: syncId,
		}),
	)
	if err != nil {
		return err
	}

	fmt.Println("sync removed successfully")
	return nil
}
