package cli

import (
	"github.com/pixelfs/pixelfs/control"
	"github.com/pixelfs/pixelfs/log"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(mvCmd)
}

var mvCmd = &cobra.Command{
	Use:   "mv <source> <destination>",
	Short: "Move file or directory",
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

		if err := app.Mv(src, dest); err != nil {
			log.Cli().Fatal().Err(err)
		}
	},
}
