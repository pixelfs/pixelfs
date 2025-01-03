package cli

import (
	"github.com/pixelfs/pixelfs/control"
	"github.com/pixelfs/pixelfs/log"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(touchCmd)
}

var touchCmd = &cobra.Command{
	Use:   "touch <file>",
	Short: "Create a new file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		app, err := control.NewPixelFS()
		if err != nil {
			log.Cli().Fatal().Err(err)
		}

		ctx, err := parseFileContext(args[0])
		if err != nil {
			log.Cli().Fatal().Err(err)
		}

		if err := app.Touch(ctx); err != nil {
			log.Cli().Fatal().Err(err)
		}
	},
}
