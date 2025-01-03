package control

import (
	"context"
	"fmt"
	"time"

	"connectrpc.com/connect"
	"github.com/charmbracelet/lipgloss"
	"github.com/pixelfs/pixelfs/config"
	pb "github.com/pixelfs/pixelfs/gen/pixelfs/v1"
	"github.com/pixelfs/pixelfs/log"
	"github.com/pixelfs/pixelfs/util"
	"github.com/spf13/viper"
)

func (p *PixelFS) AuthLogin() error {
	token, err := util.GenerateAuthToken()
	if err != nil {
		return err
	}

	if _, err = p.Core.AuthService.CreateCliSession(
		context.Background(),
		connect.NewRequest(&pb.CreateCliSessionRequest{
			Token: token,
		}),
	); err != nil {
		return err
	}

	cfg, err := config.GetConfig()
	if err != nil {
		log.Cli().Fatal().Err(err)
	}

	fmt.Println("To authenticate, Please visit:")
	fmt.Println("")
	fmt.Println(lipgloss.NewStyle().
		Bold(true).PaddingLeft(4).Render(cfg.Endpoint + "/auth/cli/" + token))

	fmt.Println("")
	fmt.Println("Waiting for session...")

	hanndleSession := func(task *util.Task) {
		response, err := p.Core.AuthService.VerifyCliSession(
			context.Background(),
			connect.NewRequest(&pb.VerifyCliSessionRequest{
				Token: token,
			}),
		)
		if err != nil {
			log.Cli().Error().Err(err)
			return
		}

		if response.Msg.AuthToken != "" {
			if err := config.Set("token", response.Msg.AuthToken); err != nil {
				log.Cli().Error().Err(err)
				return
			}

			task.Stop()
			fmt.Println("Session verified.")
		}
	}

	task, err := util.NewTask("auth:session", hanndleSession, 5*time.Second)
	if err != nil {
		return err
	}

	task.Run(context.Background())
	if util.IsAvailableAddress(p.cfg.Daemon.Listen) {
		if _, err := p.Daemon.SystemService.Startup(
			context.Background(),
			connect.NewRequest(&pb.SystemStartupRequest{
				Token: viper.GetString("token"),
			}),
		); err != nil {
			return err
		}
	}

	return nil
}
