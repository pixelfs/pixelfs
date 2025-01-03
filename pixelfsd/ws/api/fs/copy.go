package fs

import (
	"context"
	"path/filepath"

	"connectrpc.com/connect"
	"github.com/lesismal/arpc"
	cp "github.com/otiai10/copy"
	pb "github.com/pixelfs/pixelfs/gen/pixelfs/v1"
	"github.com/pixelfs/pixelfs/log"
)

func Copy(ctx *arpc.Context) {
	var request pb.FileCopyRequest
	if err := ctx.Bind(&request); handleError(ctx, err) {
		return
	}

	if request.Src.NodeId != request.Dest.NodeId {
		ctx.Error("src and dest must be in the same node")
		return
	}

	src, err := rpc.LocationService.GetLocationByContext(
		context.Background(),
		connect.NewRequest(&pb.GetLocationByContextRequest{
			Context: request.Src,
		}),
	)
	if handleError(ctx, err) {
		return
	}

	dest, err := rpc.LocationService.GetLocationByContext(
		context.Background(),
		connect.NewRequest(&pb.GetLocationByContextRequest{
			Context: request.Dest,
		}),
	)
	if handleError(ctx, err) {
		return
	}

	srcPath := filepath.Join(src.Msg.Location.Path, request.Src.Path)
	destPath := filepath.Join(dest.Msg.Location.Path, request.Dest.Path)

	if err := cp.Copy(srcPath, destPath); handleError(ctx, err) {
		return
	}
	if err := ctx.Write(&pb.FileCopyResponse{}); err != nil {
		log.Error().Caller().Err(err).Msg("write response")
	}
}
