package control

import (
	"context"

	"connectrpc.com/connect"
	pb "github.com/pixelfs/pixelfs/gen/pixelfs/v1"
)

func (p *PixelFS) SyncStart(syncId string) error {
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

	if err = p.syncStart(response.Msg.Sync.Id, response.Msg.Sync.SrcNodeId); err != nil {
		return err
	}

	targetFunc := p.syncStop
	if response.Msg.Sync.Config.Duplex {
		targetFunc = p.syncStart
	}

	return targetFunc(response.Msg.Sync.Id, response.Msg.Sync.DestNodeId)
}

func (p *PixelFS) syncStart(syncId, nodeId string) error {
	_, err := p.Core.SyncService.Start(
		context.Background(),
		connect.NewRequest(&pb.SyncStartRequest{
			SyncId: syncId,
			NodeId: nodeId,
		}),
	)

	return err
}
