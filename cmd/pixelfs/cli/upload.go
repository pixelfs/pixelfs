package cli

import (
	"github.com/pixelfs/pixelfs/control"
	"github.com/pixelfs/pixelfs/log"
	"github.com/spf13/cobra"
)

var uploadInput string

func init() {
	rootCmd.AddCommand(uploadCmd)

	uploadCmd.Flags().StringVarP(&uploadInput, "input", "i", "", "input file path")
}

var uploadCmd = &cobra.Command{
	Use:   "upload <file/dir>",
	Short: "Upload a file or directory",
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

		if err := app.Upload(ctx, uploadInput); err != nil {
			log.Cli().Fatal().Err(err)
		}
	},
}
