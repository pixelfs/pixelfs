package control

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/pixelfs/pixelfs/config"
	pb "github.com/pixelfs/pixelfs/gen/pixelfs/v1"
	"github.com/pixelfs/pixelfs/log"
	"github.com/pixelfs/pixelfs/util"
)

func (p *PixelFS) AuthLogout() error {
	if p.cfg.Token == "" {
		log.Cli().Fatal().Msg("not logged in")
	}

	if err := config.Remove("token"); err != nil {
		return err
	}

	if util.IsAvailableAddress(p.cfg.Daemon.Listen) {
		if _, err := p.Daemon.SystemService.Shutdown(
			context.Background(),
			connect.NewRequest(&pb.SystemShutdownRequest{
				Token: p.cfg.Token,
			}),
		); err != nil {
			return err
		}
	}

	fmt.Println("Logged out successfully")
	return nil
}
