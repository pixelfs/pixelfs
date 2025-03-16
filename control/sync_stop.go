package control

import (
	"context"

	"connectrpc.com/connect"
	pb "github.com/pixelfs/pixelfs/gen/pixelfs/v1"
)

func (p *PixelFS) SyncStop(syncId string) error {
	response, err := p.Core.SyncService.GetSyncById(
		context.Background(),
		connect.NewRequest(&pb.SyncGetByIdRequest{
			SyncId: syncId,
		}),
	)
	if err != nil {
		return err
	}

	if !response.Msg.Sync.Enabled {
		return nil
	}

	if err = p.syncStop(response.Msg.Sync.Id, response.Msg.Sync.SrcNodeId); err != nil {
		return err
	}
	if err = p.syncStop(response.Msg.Sync.Id, response.Msg.Sync.DestNodeId); err != nil {
		return err
	}

	return nil
}

func (p *PixelFS) syncStop(syncId, nodeId string) error {
	_, err := p.Core.SyncService.Stop(
		context.Background(),
		connect.NewRequest(&pb.SyncStopRequest{
			SyncId: syncId,
			NodeId: nodeId,
		}),
	)

	return err
}
