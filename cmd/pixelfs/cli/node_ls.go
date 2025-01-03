package cli

import (
	"github.com/pixelfs/pixelfs/control"
	"github.com/pixelfs/pixelfs/log"
	"github.com/spf13/cobra"
)

func init() {
	nodeCmd.AddCommand(nodeLsCmd)
}

var nodeLsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List nodes",
	Run: func(cmd *cobra.Command, args []string) {
		app, err := control.NewPixelFS()
		if err != nil {
			log.Cli().Fatal().Err(err)
		}

		if err := app.NodeLs(); err != nil {
			log.Cli().Fatal().Err(err)
		}
	},
}
