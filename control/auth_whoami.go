package control

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	pb "github.com/pixelfs/pixelfs/gen/pixelfs/v1"
	"github.com/pixelfs/pixelfs/log"
)

func (p *PixelFS) AuthWhoami() error {
	if p.cfg.Token == "" {
		log.Cli().Fatal().Msg("not logged in")
	}

	userInfo, err := p.Core.UserService.GetUserInfo(
		context.Background(),
		connect.NewRequest(&pb.GetUserInfoRequest{}),
	)
	if err != nil {
		return err
	}

	fmt.Println("ID:", userInfo.Msg.Id)
	fmt.Println("Name:", userInfo.Msg.Name)
	fmt.Println("Email:", userInfo.Msg.Email)

	return nil
}
