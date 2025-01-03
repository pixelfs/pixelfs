package cli

import (
	"github.com/pixelfs/pixelfs/control"
	pb "github.com/pixelfs/pixelfs/gen/pixelfs/v1"
	"github.com/pixelfs/pixelfs/log"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/proto"
)

var storageAddName string
var storageAddType string
var storageAddNetwork string
var storageAddEndpoint string
var storageAddAccessKey string
var storageAddSecretKey string
var storageAddRegion string
var storageAddBucket string
var storageAddPrefix string
var storageAddPathStyle bool

func init() {
	storageCmd.AddCommand(storageAddCmd)

	storageAddCmd.Flags().StringVarP(&storageAddName, "name", "", "", "storage name")
	storageAddCmd.Flags().StringVarP(&storageAddType, "type", "", "s3", "storage type")
	storageAddCmd.Flags().StringVarP(&storageAddNetwork, "network", "", "public", "storage network")
	storageAddCmd.Flags().StringVarP(&storageAddEndpoint, "endpoint", "", "", "storage endpoint")
	storageAddCmd.Flags().StringVarP(&storageAddAccessKey, "access-key", "", "", "access key id")
	storageAddCmd.Flags().StringVarP(&storageAddSecretKey, "secret-key", "", "", "secret access key")
	storageAddCmd.Flags().StringVarP(&storageAddRegion, "region", "", "", "s3 storage region")
	storageAddCmd.Flags().StringVarP(&storageAddBucket, "bucket", "", "", "s3 storage bucket")
	storageAddCmd.Flags().StringVarP(&storageAddPrefix, "prefix", "", "", "storage prefix")
	storageAddCmd.Flags().BoolVarP(&storageAddPathStyle, "path-style", "", false, "s3 use path style")
}

var storageAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a storage",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		var config proto.Message
		switch storageAddType {
		case "s3":
			config = &pb.StorageS3Config{
				Endpoint:  storageAddEndpoint,
				AccessKey: storageAddAccessKey,
				SecretKey: storageAddSecretKey,
				Region:    storageAddRegion,
				Bucket:    storageAddBucket,
				Prefix:    storageAddPrefix,
				PathStyle: storageAddPathStyle,
			}
		default:
			log.Cli().Fatal().Msg("unsupported storage type")
		}

		var network pb.StorageNetwork
		switch storageAddNetwork {
		case "public":
			network = pb.StorageNetwork_PUBLIC
		case "private":
			network = pb.StorageNetwork_PRIVATE
		default:
			log.Cli().Fatal().Msg("unsupported storage network")
		}

		app, err := control.NewPixelFS()
		if err != nil {
			log.Cli().Fatal().Err(err)
		}

		if err := app.StorageAdd(storageAddName, config, network); err != nil {
			log.Cli().Fatal().Err(err)
		}
	},
}
