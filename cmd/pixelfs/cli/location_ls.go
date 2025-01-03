package cli

import (
	"github.com/pixelfs/pixelfs/control"
	"github.com/pixelfs/pixelfs/log"
	"github.com/spf13/cobra"
)

func init() {
	locationCmd.AddCommand(locationLsCmd)
}

var locationLsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List locations",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		app, err := control.NewPixelFS()
		if err != nil {
			log.Cli().Fatal().Err(err)
		}

		if err := app.LocationLs(); err != nil {
			log.Cli().Fatal().Err(err)
		}
	},
}
