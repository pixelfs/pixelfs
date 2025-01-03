package cli

import (
	"github.com/pixelfs/pixelfs/control"
	"github.com/pixelfs/pixelfs/log"
	"github.com/spf13/cobra"
)

func init() {
	authCmd.AddCommand(authWhoamiCmd)
}

var authWhoamiCmd = &cobra.Command{
	Use:   "whoami",
	Short: "Prints the currently logged in user",
	Run: func(cmd *cobra.Command, args []string) {
		app, err := control.NewPixelFS()
		if err != nil {
			log.Cli().Fatal().Err(err)
		}

		if err := app.AuthWhoami(); err != nil {
			log.Cli().Fatal().Err(err)
		}
	},
}
