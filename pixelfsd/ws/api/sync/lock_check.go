package sync

import (
	"context"
	"fmt"
	"path/filepath"

	"connectrpc.com/connect"
	"github.com/lesismal/arpc"
	pb "github.com/pixelfs/pixelfs/gen/pixelfs/v1"
	"github.com/pixelfs/pixelfs/log"
)

func LockCheck(ctx *arpc.Context) {
	var request pb.SyncLockCheckRequest
	if err := ctx.Bind(&request); handleError(ctx, err) {
		return
	}

	location, err := rpc.LocationService.GetLocationByContext(
		context.Background(),
		connect.NewRequest(&pb.GetLocationByContextRequest{
			Context: request.Context,
		}),
	)
	if handleError(ctx, err) {
		return
	}

	absolutePath := filepath.Join(location.Msg.Location.Path, request.Context.Path)
	if _, exists := fileSync.Lock.Load(absolutePath); exists {
		ctx.Error(fmt.Errorf("file %s is being syncing", absolutePath))
		return
	}

	if err := ctx.Write(&pb.SyncLockCheckResponse{}); err != nil {
		log.Error().Caller().Err(err).Msg("write response")
	}
}
