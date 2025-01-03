package pixelfsd

import (
	"net/http"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/lesismal/arpc"
	arpccodec "github.com/lesismal/arpc/codec"
	arpcgzip "github.com/lesismal/arpc/extension/middleware/coder/gzip"
	"github.com/lesismal/arpc/extension/middleware/router"
	arpclog "github.com/lesismal/arpc/log"
	"github.com/pixelfs/pixelfs/config"
	"github.com/pixelfs/pixelfs/log"
	"github.com/pixelfs/pixelfs/pixelfsd/grpc"
	"github.com/pixelfs/pixelfs/pixelfsd/ws"
	"github.com/pixelfs/pixelfs/pixelfsd/ws/codec"
	"github.com/pixelfs/pixelfs/pixelfsd/ws/middleware"
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

	// aRPC
	handler := arpc.DefaultHandler
	handler.UseCoder(arpcgzip.New(1024))
	handler.Use(router.Recover())
	handler.Use(middleware.Logger())

	// Logger
	handler.SetLogTag("pixelfs rpc")
	arpclog.SetLogger(&log.ArpcLogger{})
	arpccodec.SetCodec(&codec.GRPCCodec{})

	if err := ws.InitRouters(p.cfg, handler); err != nil {
		return err
	}

	log.Info().Str("listen", p.cfg.Daemon.Listen).Msg("pixelfs daemon is running")

	if p.cfg.Token != "" {
		if err := ws.StartClient(p.cfg); err != nil {
			return err
		}
	} else {
		log.Warn().Msg("pixelfs rpc client is not initialized, token are required")
	}

	// Clean FFmpeg cache task
	cleanFFmpegCache(p.cfg)

	return http.ListenAndServe(
		p.cfg.Daemon.Listen,
		h2c.NewHandler(engine, &http2.Server{}),
	)
}
