package control

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	pb "github.com/pixelfs/pixelfs/gen/pixelfs/v1"
	"github.com/pixelfs/pixelfs/util"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func (p *PixelFS) StorageLs() error {
	response, err := p.Core.StorageService.GetStorages(
		context.Background(),
		connect.NewRequest(&pb.GetStoragesRequest{}),
	)
	if err != nil {
		return err
	}

	if len(response.Msg.Storages) == 0 {
		fmt.Println("No storage found.")
		return nil
	}

	columns := []util.TableColumn{
		{Key: "id", Title: "STORAGE_ID"},
		{Key: "name", Title: "NAME"},
		{Key: "type", Title: "TYPE"},
		{Key: "network", Title: "NETWORK"},
		{Key: "config", Title: "CONFIG"},
	}

	rows := make([]map[string]string, len(response.Msg.Storages))

	for i, storage := range response.Msg.Storages {
		var config proto.Message
		switch c := storage.Config.(type) {
		case *pb.Storage_S3:
			config = c.S3
		default:
			return fmt.Errorf("unsupported storage config type: %T", storage.Config)
		}

		jsonBytes, err := protojson.Marshal(config)
		if err != nil {
			return err
		}

		rows[i] = map[string]string{
			"id":      storage.Id,
			"name":    storage.Name,
			"type":    storage.Type.String(),
			"network": storage.Network.String(),
			"config":  string(jsonBytes),
		}
	}

	util.PrintTable(columns, rows, true)
	return nil
}
