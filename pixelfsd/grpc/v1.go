package grpc

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pixelfs/pixelfs/gen/pixelfs/v1/v1connect"
)

type GrpcV1APIService struct {
	router *gin.Engine
}

func NewGrpcV1APIService(router *gin.Engine) *GrpcV1APIService {
	return &GrpcV1APIService{router: router}
}

func (g *GrpcV1APIService) Register() {
	g.registerMetaService()
	g.registerSystemService()
}

func (g *GrpcV1APIService) registerMetaService() {
	path, handler := v1connect.NewMetaServiceHandler(NewMetaV1APIService())
	g.registerService(path, handler)
}

func (g *GrpcV1APIService) registerSystemService() {
	path, handler := v1connect.NewSystemServiceHandler(NewSystemV1APIService())
	g.registerService(path, handler)
}

func (g *GrpcV1APIService) registerService(path string, handler http.Handler) {
	g.router.POST(path+"*rpc", func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	})
}
