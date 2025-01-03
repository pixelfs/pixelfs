package cli

import (
	"github.com/pixelfs/pixelfs/control"
	"github.com/pixelfs/pixelfs/log"
	"github.com/spf13/cobra"
)

func init() {
	locationCmd.AddCommand(locationRmCmd)
}

var locationRmCmd = &cobra.Command{
	Use:   "rm <location-id>",
	Short: "Remove a location",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		app, err := control.NewPixelFS()
		if err != nil {
			log.Cli().Fatal().Err(err)
		}

		if args[0] == "" {
			log.Cli().Fatal().Msg("location id cannot be empty")
		}

		if err := app.LocationRm(args[0]); err != nil {
			log.Cli().Fatal().Err(err)
		}
	},
}
