package fs

import (
	"context"
	"os"
	"path/filepath"

	"connectrpc.com/connect"
	"github.com/lesismal/arpc"
	pb "github.com/pixelfs/pixelfs/gen/pixelfs/v1"
	"github.com/pixelfs/pixelfs/log"
)

func Remove(ctx *arpc.Context) {
	var request pb.FileRemoveRequest
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

	if request.Recursive {
		if err := os.RemoveAll(absolutePath); handleError(ctx, err) {
			return
		}
	} else {
		if err := os.Remove(absolutePath); handleError(ctx, err) {
			return
		}
	}

	if err := ctx.Write(&pb.FileRemoveResponse{}); err != nil {
		log.Error().Caller().Err(err).Msg("write response")
	}
}
