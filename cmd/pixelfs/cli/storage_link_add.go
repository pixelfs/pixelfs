package cli

import (
	"github.com/pixelfs/pixelfs/control"
	"github.com/pixelfs/pixelfs/log"
	"github.com/pixelfs/pixelfs/util"
	"github.com/spf13/cobra"
)

var storageLinkAddStorageId string
var storageLinkAddNodeId string
var storageLinkAddLocationId string
var storageLinkAddLimitSize string

func init() {
	storageLinkCmd.AddCommand(storageLinkAddCmd)

	storageLinkAddCmd.Flags().StringVarP(&storageLinkAddStorageId, "storage-id", "", "", "storage id")
	storageLinkAddCmd.Flags().StringVarP(&storageLinkAddNodeId, "node-id", "", "", "node id")
	storageLinkAddCmd.Flags().StringVarP(&storageLinkAddLocationId, "location-id", "", "", "location id")
	storageLinkAddCmd.Flags().StringVarP(&storageLinkAddLimitSize, "limit-size", "", "128MB", "limit size of the storage link")
}

var storageLinkAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a storage link",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		app, err := control.NewPixelFS()
		if err != nil {
			log.Cli().Fatal().Err(err)
		}

		limitSize, err := util.ParseBytes(storageLinkAddLimitSize)
		if err != nil {
			log.Cli().Fatal().Err(err)
		}

		if storageLinkAddNodeId != "" && !util.IsNodeId(storageLinkAddNodeId) {
			log.Cli().Fatal().Msgf("invalid node id: %s", storageLinkAddNodeId)
		}

		if err := app.StorageLinkAdd(storageLinkAddStorageId, storageLinkAddNodeId, storageLinkAddLocationId, int64(limitSize)); err != nil {
			log.Cli().Fatal().Err(err)
		}
	},
}
