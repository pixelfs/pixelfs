package cli

import (
	"github.com/pixelfs/pixelfs/control"
	"github.com/pixelfs/pixelfs/log"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(lsCmd)
}

var lsCmd = &cobra.Command{
	Use:   "ls [node:/location/path/to]",
	Short: "List files in location",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		app, err := control.NewPixelFS()
		if err != nil {
			log.Cli().Fatal().Err(err)
		}

		input := ""
		if len(args) == 1 {
			input = args[0]
		}

		ctx, err := parseFileContext(input)
		if err != nil {
			log.Cli().Fatal().Err(err)
		}

		if err := app.Ls(ctx); err != nil {
			log.Cli().Fatal().Err(err)
		}
	},
}
