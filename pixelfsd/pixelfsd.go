package pixelfsd

import (
	"net/http"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/pixelfs/pixelfs/config"
	"github.com/pixelfs/pixelfs/log"
	"github.com/pixelfs/pixelfs/pixelfsd/grpc"
	"github.com/pixelfs/pixelfs/pixelfsd/ws"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type PixelFSD struct {
	cfg *config.Config
}

func NewPixelFSD() (*PixelFSD, error) {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to get pixelfs configuration")
	}

	return &PixelFSD{cfg: cfg}, nil
}

// Serve launches the HTTP and gRPC service and the API.
func (p *PixelFSD) Serve() error {
	gin.SetMode(gin.ReleaseMode)

	engine := gin.Default()
	engine.Use(gin.Recovery())
	engine.Use(gzip.Gzip(gzip.DefaultCompression))

	// gRPC
	grpc.NewGrpcV1APIService(engine).Register()

	log.Info().Str("listen", p.cfg.Daemon.Listen).Msg("pixelfs daemon is running")

	if p.cfg.Token != "" {
		if err := ws.StartClient(p.cfg); err != nil {
			return err
		}
	} else {
		log.Warn().Msg("pixelfs rpc client is not initialized, token are required")
	}

	// Clean FFmpeg cache task
	CleanFFmpegCache(p.cfg)

	return http.ListenAndServe(
		p.cfg.Daemon.Listen,
		h2c.NewHandler(engine, &http2.Server{}),
	)
}
