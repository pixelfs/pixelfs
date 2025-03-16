package cli

import (
	"github.com/pixelfs/pixelfs/control"
	pb "github.com/pixelfs/pixelfs/gen/pixelfs/v1"
	"github.com/pixelfs/pixelfs/log"
	"github.com/spf13/cobra"
)

var syncAddId string
var syncAddName string
var syncAddEnabled bool
var syncAddDuplex bool
var syncAddInterval int64

var syncAddSrc string
var syncAddDest string

func init() {
	syncCmd.AddCommand(syncAddCmd)

	syncAddCmd.Flags().StringVarP(&syncAddId, "id", "", "", "sync id")
	syncAddCmd.Flags().StringVarP(&syncAddName, "name", "", "", "sync name")
	syncAddCmd.Flags().BoolVarP(&syncAddEnabled, "enabled", "", true, "sync enabled")
	syncAddCmd.Flags().BoolVarP(&syncAddDuplex, "duplex", "", false, "sync duplex")
	syncAddCmd.Flags().Int64VarP(&syncAddInterval, "interval", "", 3600, "sync interval")

	syncAddCmd.Flags().StringVarP(&syncAddSrc, "src", "", "", "sync source")
	syncAddCmd.Flags().StringVarP(&syncAddDest, "dest", "", "", "sync destination")
}

var syncAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add sync",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		app, err := control.NewPixelFS()
		if err != nil {
			log.Cli().Fatal().Err(err)
		}

		if syncAddSrc == "" || syncAddDest == "" {
			log.Cli().Fatal().Msg("src and dest are required")
		}

		srcContext, err := parseFileContext(syncAddSrc)
		if err != nil {
			log.Cli().Fatal().Err(err)
		}

		destContext, err := parseFileContext(syncAddDest)
		if err != nil {
			log.Cli().Fatal().Err(err)
		}

		syncConfig := &pb.SyncConfig{
			Interval: syncAddInterval,
			Duplex:   syncAddDuplex,
		}

		if err := app.SyncAdd(syncAddId, syncAddName, syncAddEnabled, syncConfig, srcContext, destContext); err != nil {
			log.Cli().Fatal().Err(err)
		}
	},
}
