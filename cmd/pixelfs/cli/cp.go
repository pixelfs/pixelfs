package cli

import (
	"github.com/pixelfs/pixelfs/control"
	"github.com/pixelfs/pixelfs/log"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(cpCmd)
}

var cpCmd = &cobra.Command{
	Use:   "cp <source> <destination>",
	Short: "Copy files",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		app, err := control.NewPixelFS()
		if err != nil {
			log.Cli().Fatal().Err(err)
		}

		src, err := parseFileContext(args[0])
		if err != nil {
			log.Cli().Fatal().Err(err)
		}

		dest, err := parseFileContext(args[1])
		if err != nil {
			log.Cli().Fatal().Err(err)
		}

		if err := app.Cp(src, dest); err != nil {
			log.Cli().Fatal().Err(err)
		}
	},
}
