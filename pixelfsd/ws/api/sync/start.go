package sync

import (
	"context"

	"connectrpc.com/connect"
	"github.com/lesismal/arpc"
	pb "github.com/pixelfs/pixelfs/gen/pixelfs/v1"
	"github.com/pixelfs/pixelfs/log"
)

func Start(ctx *arpc.Context) {
	var request pb.SyncStartRequest
	if err := ctx.Bind(&request); handleError(ctx, err) {
		return
	}

	sync, err := rpc.SyncService.GetSyncById(
		context.Background(),
		connect.NewRequest(&pb.SyncGetByIdRequest{
			SyncId: request.SyncId,
		}),
	)
	if handleError(ctx, err) {
		return
	}

	if err := fileSync.Start(sync.Msg.Sync); handleError(ctx, err) {
		return
	}
	if err := ctx.Write(&pb.SyncStartResponse{}); err != nil {
		log.Error().Caller().Err(err).Msg("write response")
	}
}
