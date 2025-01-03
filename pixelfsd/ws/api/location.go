package api

import (
	"fmt"
	"os"

	"github.com/lesismal/arpc"
	pb "github.com/pixelfs/pixelfs/gen/pixelfs/v1"
	"github.com/pixelfs/pixelfs/log"
)

func LocationCheck(ctx *arpc.Context) {
	var request pb.LocationCheckRequest
	if err := ctx.Bind(&request); err != nil {
		ctx.Error(err)
		return
	}

	fileInfo, err := os.Stat(request.Path)
	if err != nil {
		ctx.Error(err)
		return
	}

	if !fileInfo.IsDir() {
		ctx.Error(fmt.Errorf("%s is not a directory", request.Path))
		return
	}

	if err := ctx.Write(&pb.LocationCheckResponse{}); err != nil {
		log.Error().Caller().Err(err).Msg("write response")
	}
}
