package core

import (
	"context"

	"connectrpc.com/connect"
	"github.com/pixelfs/pixelfs/config"
	"github.com/pixelfs/pixelfs/gen/pixelfs/v1/v1connect"
	"github.com/pixelfs/pixelfs/util"
)

const (
	authHeader = "Authorization"
)

type GrpcV1Client struct {
	AuthService       v1connect.AuthServiceClient
	UserService       v1connect.UserServiceClient
	NodeService       v1connect.NodeServiceClient
	BlockService      v1connect.BlockServiceClient
	StorageService    v1connect.StorageServiceClient
	LocationService   v1connect.LocationServiceClient
	FileSystemService v1connect.FileSystemServiceClient
}

func NewGrpcV1Client(cfg *config.Config) *GrpcV1Client {
	client := util.Resty.GetClient()
	coreEndpoint := cfg.Endpoint + "/api"

	return &GrpcV1Client{
		AuthService: v1connect.NewAuthServiceClient(
			client,
			coreEndpoint,
			connect.WithInterceptors(newAuthInterceptor(cfg)),
		),
		UserService: v1connect.NewUserServiceClient(
			client,
			coreEndpoint,
			connect.WithInterceptors(newAuthInterceptor(cfg)),
		),
		NodeService: v1connect.NewNodeServiceClient(
			client,
			coreEndpoint,
			connect.WithInterceptors(newAuthInterceptor(cfg)),
		),
		BlockService: v1connect.NewBlockServiceClient(
			client,
			coreEndpoint,
			connect.WithInterceptors(newAuthInterceptor(cfg)),
		),
		StorageService: v1connect.NewStorageServiceClient(
			client,
			coreEndpoint,
			connect.WithInterceptors(newAuthInterceptor(cfg)),
		),
		LocationService: v1connect.NewLocationServiceClient(
			client,
			coreEndpoint,
			connect.WithInterceptors(newAuthInterceptor(cfg)),
		),
		FileSystemService: v1connect.NewFileSystemServiceClient(
			client,
			coreEndpoint,
			connect.WithInterceptors(newAuthInterceptor(cfg)),
		),
	}
}

func newAuthInterceptor(cfg *config.Config) connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			if req.Spec().IsClient && req.Header().Get(authHeader) == "" {
				req.Header().Set(authHeader, "Bearer "+cfg.Token)
			}

			return next(ctx, req)
		}
	}

	return interceptor
}
