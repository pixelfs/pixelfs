package cli

import (
	"github.com/pixelfs/pixelfs/control"
	"github.com/pixelfs/pixelfs/log"
	"github.com/spf13/cobra"
)

func init() {
	storageCmd.AddCommand(storageLsCmd)
}

var storageLsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List storages",
	Run: func(cmd *cobra.Command, args []string) {
		app, err := control.NewPixelFS()
		if err != nil {
			log.Cli().Fatal().Err(err)
		}

		if err := app.StorageLs(); err != nil {
			log.Cli().Fatal().Err(err)
		}
	},
}
