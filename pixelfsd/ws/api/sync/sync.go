package sync

import (
	"context"

	"connectrpc.com/connect"
	"github.com/lesismal/arpc"
	"github.com/pixelfs/pixelfs/config"
	pb "github.com/pixelfs/pixelfs/gen/pixelfs/v1"
	"github.com/pixelfs/pixelfs/log"
	"github.com/pixelfs/pixelfs/rpc/core"
	"github.com/pixelfs/pixelfs/sync"
)

var (
	rpc            *core.GrpcV1Client
	fileSync       *sync.FileSync
	isInitFileSync bool
)

func InitHandler(cfg *config.Config, router arpc.Handler) error {
	rpc = core.NewGrpcV1Client(cfg)

	// add routes
	router.Handle("/sync/lock-check", LockCheck)
	router.Handle("/sync/start", Start)
	router.Handle("/sync/stop", Stop)

	return nil
}

func InitFileSync(cfg *config.Config) error {
	if isInitFileSync {
		return nil
	}

	sc, err := sync.NewFileSync(cfg)
	if err != nil {
		return err
	}

	syncs, err := rpc.SyncService.GetSyncList(context.Background(),
		connect.NewRequest(&pb.SyncGetListRequest{
			NodeId: sc.NodeId,
		}),
	)
	if err != nil {
		return err
	}

	for _, v := range syncs.Msg.Syncs {
		if !v.Enabled {
			continue
		}

		if sc.NodeId == v.DestNodeId && !v.Config.Duplex {
			if err = startSync(v.Id, v.SrcNodeId); err != nil {
				log.Error().Err(err).Msgf("failed to start reverse sync %s", v.Id)
			}
		} else {
			if err = sc.Start(v); err != nil {
				log.Error().Err(err).Msgf("failed to start sync %s", v.Id)
			}
		}

		if v.Config.Duplex {
			switch sc.NodeId {
			case v.SrcNodeId:
				err = startSync(v.Id, v.DestNodeId)
			case v.DestNodeId:
				err = startSync(v.Id, v.SrcNodeId)
			}
			if err != nil {
				log.Error().Err(err).Msgf("failed to start duplex sync %s", v.Id)
			}
		}
	}

	fileSync = sc
	isInitFileSync = true
	return nil
}

func startSync(syncId, nodeId string) error {
	_, err := rpc.SyncService.Start(
		context.Background(),
		connect.NewRequest(&pb.SyncStartRequest{
			SyncId: syncId,
			NodeId: nodeId,
		}),
	)
	return err
}

func handleError(ctx *arpc.Context, err error) bool {
	if err != nil {
		ctx.Error(err)
		return true
	}
	return false
}
