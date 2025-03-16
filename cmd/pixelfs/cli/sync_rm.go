package cli

import (
	"github.com/pixelfs/pixelfs/control"
	"github.com/pixelfs/pixelfs/log"
	"github.com/spf13/cobra"
)

func init() {
	syncCmd.AddCommand(syncRmCmd)
}

var syncRmCmd = &cobra.Command{
	Use:   "rm <sync>",
	Short: "Remove sync",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		app, err := control.NewPixelFS()
		if err != nil {
			log.Cli().Fatal().Err(err)
		}

		if err := app.SyncRm(args[0]); err != nil {
			log.Cli().Fatal().Err(err)
		}
	},
}
