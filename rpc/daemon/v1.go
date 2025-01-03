package daemon

import (
	"net/http"

	"connectrpc.com/connect"
	"github.com/pixelfs/pixelfs/gen/pixelfs/v1/v1connect"
)

type GrpcV1Client struct {
	Url           string
	SystemService v1connect.SystemServiceClient
}

func NewGrpcV1Client(url string) *GrpcV1Client {
	return &GrpcV1Client{
		Url: url,
		SystemService: v1connect.NewSystemServiceClient(
			http.DefaultClient,
			url,
			connect.WithGRPC(),
		),
	}
}
