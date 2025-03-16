package fs

import (
	"github.com/lesismal/arpc"
	"github.com/pixelfs/pixelfs/config"
	"github.com/pixelfs/pixelfs/rpc/core"
)

var (
	rpc *core.GrpcV1Client
)

func InitHandler(cfg *config.Config, router arpc.Handler) error {
	rpc = core.NewGrpcV1Client(cfg)

	// add routes
	router.Handle("/fs/list", List)
	router.Handle("/fs/stat", Stat)
	router.Handle("/fs/create", Create)
	router.Handle("/fs/remove", Remove)
	router.Handle("/fs/copy", Copy)
	router.Handle("/fs/move", Move)
	router.Handle("/fs/mkdir", Mkdir)
	router.Handle("/fs/read", Read)
	router.Handle("/fs/write", Write)
	router.Handle("/fs/m3u8", M3U8)
	router.Handle("/fs/chmod", Chmod)
	router.Handle("/fs/chtimes", Chtimes)

	return nil
}

func handleError(ctx *arpc.Context, err error) bool {
	if err != nil {
		ctx.Error(err)
		return true
	}
	return false
}
