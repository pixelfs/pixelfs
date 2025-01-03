package control

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	pb "github.com/pixelfs/pixelfs/gen/pixelfs/v1"
	"github.com/pixelfs/pixelfs/util"
)

func (p *PixelFS) LocationLs() error {
	response, err := p.Core.LocationService.GetLocations(
		context.Background(),
		connect.NewRequest(&pb.GetLocationsRequest{}),
	)
	if err != nil {
		return err
	}

	if len(response.Msg.Locations) == 0 {
		fmt.Println("No locations found.")
		return nil
	}

	columns := []util.TableColumn{
		{Key: "id", Title: "LOCATION_ID"},
		{Key: "node", Title: "NODE_ID"},
		{Key: "name", Title: "NAME"},
		{Key: "path", Title: "PATH"},
		{Key: "block_size", Title: "BLOCK_SIZE"},
		{Key: "block_duration", Title: "BLOCK_DURATION"},
	}

	rows := make([]map[string]string, len(response.Msg.Locations))
	for i, location := range response.Msg.Locations {
		rows[i] = map[string]string{
			"id":             location.Id,
			"node":           location.NodeId,
			"name":           location.Name,
			"path":           location.Path,
			"block_size":     util.Bytes(uint64(location.BlockSize)),
			"block_duration": fmt.Sprintf("%ds", location.BlockDuration),
		}
	}

	util.PrintTable(columns, rows, true)
	return nil
}
