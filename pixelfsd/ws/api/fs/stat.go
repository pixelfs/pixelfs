package fs

import (
	"context"
	"os"
	"path/filepath"
	"runtime"

	"connectrpc.com/connect"
	"github.com/lesismal/arpc"
	pb "github.com/pixelfs/pixelfs/gen/pixelfs/v1"
	"github.com/pixelfs/pixelfs/log"
	"github.com/pixelfs/pixelfs/util"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Stat(ctx *arpc.Context) {
	var request pb.FileStatRequest
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
	file, err := os.Open(absolutePath)
	if handleError(ctx, err) {
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if handleError(ctx, err) {
		return
	}

	var hash string
	if request.Hash && !fileInfo.IsDir() {
		hash, err = util.GetFileHash(absolutePath)
		if handleError(ctx, err) {
			return
		}
	}

	var duration float64
	if request.Duration && !fileInfo.IsDir() {
		metadata, err := util.GetFFmpegMetadata(absolutePath)
		if handleError(ctx, err) {
			return
		}

		duration = metadata.Duration
	}

	response := &pb.FileStatResponse{
		File: &pb.File{
			Name:       fileInfo.Name(),
			Type:       util.GetFileType(fileInfo),
			Size:       fileInfo.Size(),
			Hash:       hash,
			Duration:   duration,
			Mode:       uint32(fileInfo.Mode()),
			Platform:   runtime.GOOS,
			ModifiedAt: timestamppb.New(fileInfo.ModTime()),
		},
	}

	if err := ctx.Write(response); err != nil {
		log.Error().Caller().Err(err).Msg("write response")
	}
}
