package cli

import (
	"github.com/pixelfs/pixelfs/control"
	"github.com/pixelfs/pixelfs/log"
	"github.com/pixelfs/pixelfs/util"
	"github.com/spf13/cobra"
)

var locationAddNodeId string
var locationAddName string
var locationAddType string
var locationAddPath string
var locationAddBlockSize string
var locationAddBlockDuration int64

func init() {
	locationCmd.AddCommand(locationAddCmd)

	locationAddCmd.Flags().StringVarP(&locationAddNodeId, "node-id", "", "", "node id")
	locationAddCmd.Flags().StringVarP(&locationAddName, "name", "", "", "location name")
	locationAddCmd.Flags().StringVarP(&locationAddType, "type", "", "local", "location type")
	locationAddCmd.Flags().StringVarP(&locationAddPath, "path", "", "", "location path")
	locationAddCmd.Flags().StringVarP(&locationAddBlockSize, "block-size", "", "4MB", "block size")
	locationAddCmd.Flags().Int64VarP(&locationAddBlockDuration, "block-duration", "", 20, "block duration")
}

var locationAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a location",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if !util.IsNodeId(locationAddNodeId) {
			log.Cli().Fatal().Msgf("invalid node id: %s", locationAddNodeId)
		}

		if locationAddName == "" {
			log.Cli().Fatal().Msg("location name is required")
		}

		if locationAddType == "" {
			log.Cli().Fatal().Msg("location type is required")
		}

		if locationAddPath == "" {
			log.Cli().Fatal().Msg("location path is required")
		}

		app, err := control.NewPixelFS()
		if err != nil {
			log.Cli().Fatal().Err(err)
		}

		blockSize, err := util.ParseBytes(locationAddBlockSize)
		if err != nil {
			log.Cli().Fatal().Err(err)
		}

		if err := app.LocationAdd(locationAddNodeId, locationAddName, locationAddType, locationAddPath, int64(blockSize), locationAddBlockDuration); err != nil {
			log.Cli().Fatal().Err(err)
		}
	},
}
