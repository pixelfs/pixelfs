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

func Chtimes(ctx *arpc.Context) {
	var request pb.FileChtimesRequest
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
	if err := os.Chtimes(absolutePath, request.Atime.AsTime(), request.Mtime.AsTime()); handleError(ctx, err) {
		return
	}
	if err := ctx.Write(&pb.FileChtimesResponse{}); err != nil {
		log.Error().Caller().Err(err).Msg("write response")
	}
}
