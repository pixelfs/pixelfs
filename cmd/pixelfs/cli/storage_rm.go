package cli

import (
	"github.com/pixelfs/pixelfs/control"
	"github.com/pixelfs/pixelfs/log"
	"github.com/spf13/cobra"
)

func init() {
	storageCmd.AddCommand(storageRmCmd)
}

var storageRmCmd = &cobra.Command{
	Use:   "rm <storage>",
	Short: "Remove a storage",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		app, err := control.NewPixelFS()
		if err != nil {
			log.Cli().Fatal().Err(err)
		}

		if err := app.StorageRm(args[0]); err != nil {
			log.Cli().Fatal().Err(err)
		}
	},
}
