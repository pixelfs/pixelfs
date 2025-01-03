package fs

import (
	"context"
	"os"
	"path/filepath"

	"connectrpc.com/connect"
	"github.com/lesismal/arpc"
	pb "github.com/pixelfs/pixelfs/gen/pixelfs/v1"
	"github.com/pixelfs/pixelfs/log"
	"github.com/pixelfs/pixelfs/util"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func List(ctx *arpc.Context) {
	var request pb.FileListRequest
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
	fileInfo, err := os.Stat(absolutePath)
	if err != nil {
		ctx.Error(err)
		return
	}

	if !fileInfo.IsDir() {
		var response pb.FileListResponse
		response.Files = append(response.Files, &pb.File{
			Name:       fileInfo.Name(),
			Type:       util.GetFileType(fileInfo),
			Size:       fileInfo.Size(),
			ModifiedAt: timestamppb.New(fileInfo.ModTime()),
		})

		if err := ctx.Write(&response); err != nil {
			log.Error().Caller().Err(err).Msg("write response")
		}
		return
	}

	files, err := os.ReadDir(absolutePath)
	if handleError(ctx, err) {
		return
	}

	var response pb.FileListResponse
	for _, file := range files {
		fileInfo, err := file.Info()
		if err != nil {
			continue
		}

		response.Files = append(response.Files, &pb.File{
			Name:       fileInfo.Name(),
			Type:       util.GetFileType(fileInfo),
			Size:       fileInfo.Size(),
			ModifiedAt: timestamppb.New(fileInfo.ModTime()),
		})
	}

	if err := ctx.Write(&response); err != nil {
		log.Error().Caller().Err(err).Msg("write response")
	}
}
