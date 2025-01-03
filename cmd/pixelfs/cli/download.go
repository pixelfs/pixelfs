package cli

import (
	"github.com/pixelfs/pixelfs/control"
	"github.com/pixelfs/pixelfs/log"
	"github.com/spf13/cobra"
)

var downloadOutput string
var downloadThread int

func init() {
	rootCmd.AddCommand(downloadCmd)

	downloadCmd.Flags().StringVarP(&downloadOutput, "output", "o", "", "output file path")
	downloadCmd.Flags().IntVarP(&downloadThread, "thread", "t", 1, "number of threads")
}

var downloadCmd = &cobra.Command{
	Use:   "download <file/dir>",
	Short: "Download a file or directory",
	Args:  cobra.ExactArgs(1),
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

		if err := app.Download(ctx, downloadOutput, downloadThread); err != nil {
			log.Cli().Fatal().Err(err)
		}
	},
}
