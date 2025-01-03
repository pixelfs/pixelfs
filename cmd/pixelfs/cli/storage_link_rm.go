package cli

import (
	"github.com/pixelfs/pixelfs/control"
	"github.com/pixelfs/pixelfs/log"
	"github.com/spf13/cobra"
)

func init() {
	storageLinkCmd.AddCommand(storageLinkRmCmd)
}

var storageLinkRmCmd = &cobra.Command{
	Use:   "rm <storage-link>",
	Short: "Remove a storage link",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		app, err := control.NewPixelFS()
		if err != nil {
			log.Cli().Fatal().Err(err)
		}

		if err := app.StorageLinkRm(args[0]); err != nil {
			log.Cli().Fatal().Err(err)
		}
	},
}
