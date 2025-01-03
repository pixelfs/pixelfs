package cli

import (
	"github.com/pixelfs/pixelfs/control"
	"github.com/pixelfs/pixelfs/log"
	"github.com/spf13/cobra"
)

func init() {
	storageLinkCmd.AddCommand(storageLinkLsCmd)
}

var storageLinkLsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List storage links",
	Run: func(cmd *cobra.Command, args []string) {
		app, err := control.NewPixelFS()
		if err != nil {
			log.Cli().Fatal().Err(err)
		}

		if err := app.StorageLinkLs(); err != nil {
			log.Cli().Fatal().Err(err)
		}
	},
}
