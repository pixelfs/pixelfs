package cli

import (
	"github.com/pixelfs/pixelfs/config"
	"github.com/pixelfs/pixelfs/control"
	"github.com/pixelfs/pixelfs/log"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(cdCmd)
}

var cdCmd = &cobra.Command{
	Use:   "cd [node:/location/path/to]",
	Short: "Change directory",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 || args[0] == "::" {
			if err := config.Set("pwd", "/"); err != nil {
				log.Cli().Fatal().Err(err)
			}
			return
		}

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

		if err := app.Cd(ctx); err != nil {
			log.Cli().Fatal().Err(err)
		}

		pwd := ctx.NodeId + ":/"
		if ctx.Location != "" {
			pwd += ctx.Location
		}

		if ctx.Path != "" {
			pwd += "/" + ctx.Path
		}

		if err = config.Set("pwd", pwd); err != nil {
			log.Cli().Fatal().Err(err)
		}
	},
}
