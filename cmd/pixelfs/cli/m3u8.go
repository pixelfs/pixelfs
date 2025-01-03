package cli

import (
	"github.com/pixelfs/pixelfs/control"
	"github.com/pixelfs/pixelfs/log"
	"github.com/spf13/cobra"
)

var m3u8Width int
var m3u8Height int
var m3u8Bitrate int

func init() {
	rootCmd.AddCommand(m3u8Cmd)

	m3u8Cmd.Flags().IntVarP(&m3u8Width, "width", "", 0, "video width, eg. 1920")
	m3u8Cmd.Flags().IntVarP(&m3u8Height, "height", "", 0, "video height, eg. 1080")
	m3u8Cmd.Flags().IntVarP(&m3u8Bitrate, "bitrate", "", 0, "video bitrate, in kbps, eg. 5000")
}

var m3u8Cmd = &cobra.Command{
	Use:   "m3u8 <file>",
	Short: "Generate m3u8 playlist",
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

		if err := app.M3U8(ctx, m3u8Width, m3u8Height, m3u8Bitrate); err != nil {
			log.Cli().Fatal().Err(err)
		}
	},
}
