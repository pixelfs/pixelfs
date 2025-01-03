package cli

import (
	"fmt"

	"github.com/pixelfs/pixelfs/log"
	"github.com/pixelfs/pixelfs/pixelfsd"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(daemonCmd)
}

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Run the daemon server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`
 ____  _ ___  _ _____ _     _____ ____ 
/  __\/ \\  \///  __// \   /    // ___\
|  \/|| | \  / |  \  | |   |  __\|    \
|  __/| | /  \ |  /_ | |_/\| |   \___ |
\_/   \_//__/\\\____\\____/\_/   \____/ Version: ` + Version +
			"\n")

		app, err := pixelfsd.NewPixelFSD()
		if err != nil {
			log.Fatal().Err(err).Msg("failed to create pixelfsd")
		}

		log.Fatal().Err(app.Serve()).Msg("failed to serve pixelfsd")
	},
}
