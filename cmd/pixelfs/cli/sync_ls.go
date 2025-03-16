package cli

import (
	"github.com/pixelfs/pixelfs/control"
	"github.com/pixelfs/pixelfs/log"
	"github.com/spf13/cobra"
)

func init() {
	syncCmd.AddCommand(syncLsCmd)
}

var syncLsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List syncs",
	Run: func(cmd *cobra.Command, args []string) {
		app, err := control.NewPixelFS()
		if err != nil {
			log.Cli().Fatal().Err(err)
		}

		if err := app.SyncLs(); err != nil {
			log.Cli().Fatal().Err(err)
		}
	},
}
