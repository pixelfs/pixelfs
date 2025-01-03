package control

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	pb "github.com/pixelfs/pixelfs/gen/pixelfs/v1"
	"github.com/pixelfs/pixelfs/util"
)

func (p *PixelFS) StorageLinkLs() error {
	response, err := p.Core.StorageService.GetStorageLinks(
		context.Background(),
		connect.NewRequest(&pb.GetStorageLinksRequest{}),
	)
	if err != nil {
		return err
	}

	if len(response.Msg.StorageLinks) == 0 {
		fmt.Println("No storage links found.")
		return nil
	}

	columns := []util.TableColumn{
		{Key: "id", Title: "STORAGE_LINK_ID"},
		{Key: "storage_id", Title: "STORAGE_ID"},
		{Key: "node_id", Title: "NODE_ID"},
		{Key: "location_id", Title: "LOCATION_ID"},
		{Key: "limit_size", Title: "LIMIT_SIZE"},
		{Key: "used_size", Title: "USED_SIZE"},
	}

	rows := make([]map[string]string, len(response.Msg.StorageLinks))

	for i, storageLink := range response.Msg.StorageLinks {
		rows[i] = map[string]string{
			"id":          storageLink.Id,
			"storage_id":  storageLink.StorageId,
			"node_id":     storageLink.NodeId,
			"location_id": storageLink.LocationId,
			"limit_size":  util.Bytes(uint64(storageLink.LimitSize)),
			"used_size":   util.Bytes(uint64(storageLink.UsedSize)),
		}
	}

	util.PrintTable(columns, rows, true)
	return nil
}
