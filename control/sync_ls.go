package control

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	pb "github.com/pixelfs/pixelfs/gen/pixelfs/v1"
	"github.com/pixelfs/pixelfs/util"
	"google.golang.org/protobuf/encoding/protojson"
)

func (p *PixelFS) SyncLs() error {
	response, err := p.Core.SyncService.GetSyncList(
		context.Background(),
		connect.NewRequest(&pb.SyncGetListRequest{}),
	)
	if err != nil {
		return err
	}

	if len(response.Msg.Syncs) == 0 {
		fmt.Println("No sync found.")
		return nil
	}

	columns := []util.TableColumn{
		{Key: "id", Title: "SYNC_ID"},
		{Key: "name", Title: "NAME"},
		{Key: "enabled", Title: "ENABLED"},
		{Key: "duplex", Title: "DUPLEX"},
		{Key: "interval", Title: "INTERVAL"},
		{Key: "status", Title: "STATUS"},
		{Key: "src", Title: "SRC"},
		{Key: "dest", Title: "DEST"},
		{Key: "log", Title: "LOG"},
	}

	rows := make([]map[string]string, len(response.Msg.Syncs))

	for i, sync := range response.Msg.Syncs {
		srcBytes, err := protojson.Marshal(sync.SrcContext)
		if err != nil {
			return err
		}

		destBytes, err := protojson.Marshal(sync.DestContext)
		if err != nil {
			return err
		}

		rows[i] = map[string]string{
			"id":       sync.Id,
			"name":     sync.Name,
			"enabled":  fmt.Sprintf("%t", sync.Enabled),
			"duplex":   fmt.Sprintf("%t", sync.Config.Duplex),
			"interval": fmt.Sprintf("%d", sync.Config.Interval),
			"status":   sync.Status.String(),
			"src":      string(srcBytes),
			"dest":     string(destBytes),
			"log":      sync.Config.Log,
		}
	}

	util.PrintTable(columns, rows, true)
	return nil
}
