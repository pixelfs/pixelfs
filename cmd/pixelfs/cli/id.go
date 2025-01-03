package cli

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/pixelfs/pixelfs/config"
	pb "github.com/pixelfs/pixelfs/gen/pixelfs/v1"
	"github.com/pixelfs/pixelfs/log"
	"github.com/pixelfs/pixelfs/rpc/core"
	"github.com/pixelfs/pixelfs/util"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(idCmd)
}

var idCmd = &cobra.Command{
	Use:   "id",
	Short: "Print the node id",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.GetConfig()
		if err != nil {
			log.Cli().Fatal().Msg(err)
		}

		if cfg.Token == "" {
			log.Cli().Fatal().Msg("not logged in")
		}

		userInfo, err := core.NewGrpcV1Client(cfg).UserService.GetUserInfo(
			context.Background(),
			connect.NewRequest(&pb.GetUserInfoRequest{}),
		)
		if err != nil {
			log.Cli().Fatal().Msg(err)
		}

		nodeId, err := util.GetNodeId(userInfo.Msg.Id)
		if err != nil {
			log.Cli().Fatal().Msg(err)
		}

		fmt.Println(nodeId)
	},
}
