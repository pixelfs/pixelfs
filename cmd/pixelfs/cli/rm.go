package cli

import (
	"github.com/pixelfs/pixelfs/control"
	"github.com/pixelfs/pixelfs/log"
	"github.com/spf13/cobra"
)

var rmRecursive bool

func init() {
	rootCmd.AddCommand(rmCmd)

	rmCmd.Flags().BoolVarP(&rmRecursive, "recursive", "r", false, "remove directories and their contents recursively")
}

var rmCmd = &cobra.Command{
	Use:   "rm <file/dir>",
	Short: "Remove file or directory",
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

		if err := app.Rm(ctx, rmRecursive); err != nil {
			log.Cli().Fatal().Err(err)
		}
	},
}
