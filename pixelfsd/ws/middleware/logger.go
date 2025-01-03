package middleware

import (
	"fmt"
	"strconv"
	"time"

	"github.com/lesismal/arpc"
	"github.com/pixelfs/pixelfs/log"
)

var (
	cmdName = map[byte]string{
		arpc.CmdRequest: "request",
		arpc.CmdNotify:  "notify",
	}
)

// Logger returns the logger middleware.
func Logger() arpc.HandlerFunc {
	return func(ctx *arpc.Context) {
		t := time.Now()

		ctx.Next()

		cmd := ctx.Message.Cmd()
		method := ctx.Message.Method()
		bodyLen := ctx.Message.BodyLen()
		cost := time.Since(t).Milliseconds()

		switch cmd {
		case arpc.CmdRequest, arpc.CmdNotify:
			err := ctx.ResponseError()
			if err == nil {
				log.Info().
					Str("cmd", cmdName[cmd]).
					Str("method", method).
					Int("length", bodyLen).
					Str("cost", fmt.Sprintf("%vms", strconv.FormatInt(cost, 10))).
					Msg("rpc call method")
			} else {
				log.Error().Err(fmt.Errorf("%v", err)).
					Str("cmd", cmdName[cmd]).
					Str("method", method).
					Int("length", bodyLen).
					Str("cost", fmt.Sprintf("%vms", strconv.FormatInt(cost, 10))).
					Msg("rpc call method")
			}
		default:
			log.Error().Str("msg", fmt.Sprintf("unknown cmd: %v", cmd)).Msg("rpc call method")
			ctx.Done()
		}
	}
}
