package cli

import (
	"github.com/pixelfs/pixelfs/control"
	"github.com/pixelfs/pixelfs/log"
	"github.com/spf13/cobra"
)

func init() {
	nodeCmd.AddCommand(nodeRmCmd)
}

var nodeRmCmd = &cobra.Command{
	Use:   "rm <node>",
	Short: "Remove node",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		app, err := control.NewPixelFS()
		if err != nil {
			log.Cli().Fatal().Err(err)
		}

		if args[0] == "" {
			log.Cli().Fatal().Msg("node id cannot be empty")
		}

		if err := app.NodeRm(args[0]); err != nil {
			log.Cli().Fatal().Err(err)
		}
	},
}