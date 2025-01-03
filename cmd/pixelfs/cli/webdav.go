package cli

import (
	"fmt"

	"github.com/pixelfs/pixelfs/log"
	"github.com/pixelfs/pixelfs/webdav"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(webdavCmd)
}

var webdavCmd = &cobra.Command{
	Use:   "webdav",
	Short: "Run the webdav server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`
 ____  _ ___  _ _____ _     _____ ____ 
/  __\/ \\  \///  __// \   /    // ___\
|  \/|| | \  / |  \  | |   |  __\|    \
|  __/| | /  \ |  /_ | |_/\| |   \___ |
\_/   \_//__/\\\____\\____/\_/   \____/ Version: ` + Version +
			"\n")

		app, err := webdav.NewPixelFS()
		if err != nil {
			log.Fatal().Err(err).Msg("failed to create webdav server")
		}

		log.Fatal().Err(app.Serve()).Msg("failed to serve webdav server")
	},
}
