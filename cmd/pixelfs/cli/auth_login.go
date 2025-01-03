package cli

import (
	"github.com/pixelfs/pixelfs/control"
	"github.com/pixelfs/pixelfs/log"
	"github.com/spf13/cobra"
)

func init() {
	authCmd.AddCommand(authLoginCmd)
}

var authLoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in a user",
	Run: func(cmd *cobra.Command, args []string) {
		app, err := control.NewPixelFS()
		if err != nil {
			log.Cli().Fatal().Err(err)
		}

		if err := app.AuthLogin(); err != nil {
			log.Cli().Fatal().Err(err)
		}
	},
}
