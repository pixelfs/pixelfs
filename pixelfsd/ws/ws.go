package ws

import (
	"context"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"connectrpc.com/connect"
	"github.com/lesismal/arpc"
	"github.com/lesismal/arpc/extension/protocol/websocket"
	"github.com/pixelfs/pixelfs/config"
	pb "github.com/pixelfs/pixelfs/gen/pixelfs/v1"
	"github.com/pixelfs/pixelfs/log"
	"github.com/pixelfs/pixelfs/pixelfsd/ws/api"
	"github.com/pixelfs/pixelfs/pixelfsd/ws/api/fs"
	"github.com/pixelfs/pixelfs/rpc/core"
	"github.com/pixelfs/pixelfs/util"
)

const CallTimeout = 2 * time.Minute

var (
	Client *arpc.Client
)

func InitRouters(cfg *config.Config, router arpc.Handler) error {
	router.Handle("/location/check", api.LocationCheck)
	router.Handle("/storage/validate", api.StorageValidate)
	router.Handle("/storage/remove-block", api.StorageRemoveBlock)

	// File System
	err := fs.InitRouters(cfg, router)
	if err != nil {
		return err
	}

	return nil
}

func StartClient(cfg *config.Config) error {
	userInfo, err := core.NewGrpcV1Client(cfg).UserService.GetUserInfo(
		context.Background(),
		connect.NewRequest(&pb.GetUserInfoRequest{}),
	)
	if err != nil {
		return err
	}

	token := cfg.Token
	nodeId, err := util.GetNodeId(userInfo.Msg.Id)
	if err != nil {
		return err
	}

	arpc.DefaultHandler.HandleConnected(func(c *arpc.Client) {
		hostname, err := os.Hostname()
		if err != nil {
			log.Fatal().Err(err).Msg("failed to get hostname")
		}

		var response pb.NodeRegisterResponse
		request := pb.NodeRegisterRequest{
			NodeId: nodeId,
			Token:  token,
			Name:   hostname,
		}

		if err = c.Call("/node/register", &request, &response, CallTimeout); err != nil {
			log.Fatal().Err(err).Msg("failed to call /node/register")
		}
	})

	wsEndpoint, err := httpToWebSocket(cfg.Endpoint)
	if err != nil {
		return err
	}

	client, err := arpc.NewClient(func() (net.Conn, error) {
		return websocket.Dial(wsEndpoint+"/ws?id="+nodeId+"&t="+token, nil)
	})
	if err != nil {
		return err
	}

	task, err := util.NewTask("ws:ping", func(task *util.Task) {
		if err = client.Call("/ping", nil, nil, 15*time.Second); err != nil {
			log.Error().Err(err).Msg("pixelfs rpc ping failed and restarting client")
			_ = client.Restart()
			return
		}
	}, 30*time.Second)

	if err != nil {
		return fmt.Errorf("failed to create ping task: %w", err)
	}

	go task.Run(context.Background())

	log.Info().
		Str("nodeId", nodeId).
		Str("userId", userInfo.Msg.Id).
		Msg("pixelfs rpc client is initialized and ready")

	Client = client
	return nil
}

func StopClient() {
	if Client == nil {
		return
	}

	Client.Stop()
	Client = nil
	return
}

func httpToWebSocket(url string) (string, error) {
	if strings.HasPrefix(url, "https://") {
		return "wss://" + strings.TrimPrefix(url, "https://"), nil
	} else if strings.HasPrefix(url, "http://") {
		return "ws://" + strings.TrimPrefix(url, "http://"), nil
	}

	return "", fmt.Errorf("invalid url: %s", url)
}
