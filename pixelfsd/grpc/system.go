package grpc

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/pixelfs/pixelfs/config"
	pb "github.com/pixelfs/pixelfs/gen/pixelfs/v1"
	"github.com/pixelfs/pixelfs/gen/pixelfs/v1/v1connect"
	"github.com/pixelfs/pixelfs/pixelfsd/ws"
	"github.com/spf13/viper"
)

type SystemV1APIService struct {
	v1connect.UnimplementedSystemServiceHandler
}

func NewSystemV1APIService() v1connect.SystemServiceHandler {
	return SystemV1APIService{}
}

func (api SystemV1APIService) Startup(
	ctx context.Context,
	request *connect.Request[pb.SystemStartupRequest],
) (*connect.Response[pb.SystemStartupResponse], error) {
	if err := viper.ReadInConfig(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	cfg, err := config.GetConfig()
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if cfg.Token != request.Msg.Token {
		return nil, connect.NewError(connect.CodeUnauthenticated, fmt.Errorf("invalid token"))
	}

	if ws.Client != nil {
		return connect.NewResponse(&pb.SystemStartupResponse{}), nil
	}

	if cfg.Token == "" {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("token is empty"))
	}

	if err := ws.StartClient(cfg); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&pb.SystemStartupResponse{}), nil
}

func (api SystemV1APIService) Shutdown(
	ctx context.Context,
	request *connect.Request[pb.SystemShutdownRequest],
) (*connect.Response[pb.SystemShutdownResponse], error) {
	cfg, err := config.GetConfig()
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if cfg.Token != request.Msg.Token {
		return nil, connect.NewError(connect.CodeUnauthenticated, fmt.Errorf("invalid token"))
	}

	ws.StopClient() // stop client
	return connect.NewResponse(&pb.SystemShutdownResponse{}), nil
}
