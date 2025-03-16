package sync

import (
	"context"

	"connectrpc.com/connect"
	"github.com/lesismal/arpc"
	pb "github.com/pixelfs/pixelfs/gen/pixelfs/v1"
	"github.com/pixelfs/pixelfs/log"
)

func Stop(ctx *arpc.Context) {
	var request pb.SyncStopRequest
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

	if err := fileSync.Stop(sync.Msg.Sync); handleError(ctx, err) {
		return
	}
	if err := ctx.Write(&pb.SyncStopResponse{}); err != nil {
		log.Error().Caller().Err(err).Msg("write response")
	}
}
