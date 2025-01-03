package cli

import (
	"github.com/pixelfs/pixelfs/control"
	"github.com/pixelfs/pixelfs/log"
	"github.com/spf13/cobra"
)

func init() {
	storageLinkCmd.AddCommand(storageLinkCleanCmd)
}

var storageLinkCleanCmd = &cobra.Command{
	Use:   "clean <storage-link>",
	Short: "Clean a storage link",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		app, err := control.NewPixelFS()
		if err != nil {
			log.Cli().Fatal().Err(err)
		}

		if err := app.StorageLinkClean(args[0]); err != nil {
			log.Cli().Fatal().Err(err)
		}
	},
}
