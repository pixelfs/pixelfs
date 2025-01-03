package control

import (
	"strings"

	"github.com/pixelfs/pixelfs/config"
	"github.com/pixelfs/pixelfs/log"
	"github.com/pixelfs/pixelfs/rpc/core"
	"github.com/pixelfs/pixelfs/rpc/daemon"
)

type PixelFS struct {
	cfg    *config.Config
	Core   *core.GrpcV1Client
	Daemon *daemon.GrpcV1Client
}

func NewPixelFS() (*PixelFS, error) {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Cli().Fatal().Err(err)
	}

	daemonEndpoint := cfg.Daemon.Listen
	if strings.HasPrefix(daemonEndpoint, "0.0.0.0") {
		daemonEndpoint = "127.0.0.1" + daemonEndpoint[7:]
	}

	return &PixelFS{
		cfg:    cfg,
		Core:   core.NewGrpcV1Client(cfg),
		Daemon: daemon.NewGrpcV1Client("http://" + daemonEndpoint),
	}, nil
}
