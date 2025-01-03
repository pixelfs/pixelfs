package grpc

import (
	"context"

	"connectrpc.com/connect"
	pb "github.com/pixelfs/pixelfs/gen/pixelfs/v1"
	"github.com/pixelfs/pixelfs/gen/pixelfs/v1/v1connect"
)

type MetaV1APIService struct {
	v1connect.UnimplementedMetaServiceHandler
}

func NewMetaV1APIService() v1connect.MetaServiceHandler {
	return MetaV1APIService{}
}

func (api MetaV1APIService) GetVersion(
	ctx context.Context,
	request *connect.Request[pb.GetVersionRequest],
) (*connect.Response[pb.GetVersionResponse], error) {
	return connect.NewResponse(&pb.GetVersionResponse{
		Name:    "pixelfsd",
		Version: "v0.0.1",
	}), nil
}
