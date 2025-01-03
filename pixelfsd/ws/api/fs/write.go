package fs

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"connectrpc.com/connect"
	"github.com/lesismal/arpc"
	pb "github.com/pixelfs/pixelfs/gen/pixelfs/v1"
	"github.com/pixelfs/pixelfs/log"
	"github.com/pixelfs/pixelfs/util"
)

func Write(ctx *arpc.Context) {
	var request pb.FileWriteRequest
	if err := ctx.Bind(&request); handleError(ctx, err) {
		return
	}

	if request.BlockType != pb.BlockType_SIZE {
		ctx.Error(fmt.Errorf("unsupported block type: %s", request.BlockType))
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
	file, err := os.OpenFile(absolutePath, os.O_CREATE|os.O_WRONLY, 0644)
	if handleError(ctx, err) {
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if handleError(ctx, err) || fileInfo.IsDir() {
		ctx.Error(fmt.Errorf("%s is a directory", request.Context.Path))
		return
	}

	resp, err := util.Resty.R().Get(request.Url)
	if handleError(ctx, err) {
		return
	}

	_, err = file.WriteAt(resp.Body(), request.Offset)
	if handleError(ctx, err) {
		return
	}

	if err := ctx.Write(&pb.FileWriteResponse{}); err != nil {
		log.Error().Caller().Err(err).Msg("write response")
	}
}
