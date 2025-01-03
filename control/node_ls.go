package control

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	pb "github.com/pixelfs/pixelfs/gen/pixelfs/v1"
	"github.com/pixelfs/pixelfs/util"
)

func (p *PixelFS) NodeLs() error {
	response, err := p.Core.NodeService.GetNodes(
		context.Background(),
		connect.NewRequest(&pb.GetNodesRequest{}),
	)
	if err != nil {
		return err
	}

	if len(response.Msg.Nodes) == 0 {
		fmt.Println("No nodes found.")
		return nil
	}

	columns := []util.TableColumn{
		{Key: "id", Title: "NODE_ID"},
		{Key: "name", Title: "NAME"},
		{Key: "status", Title: "STATUS"},
		{Key: "createdAt", Title: "CREATED_AT"},
		{Key: "updatedAt", Title: "UPDATED_AT"},
	}

	rows := make([]map[string]string, len(response.Msg.Nodes))
	for i, node := range response.Msg.Nodes {
		rows[i] = map[string]string{
			"id":        node.Id,
			"name":      node.Name,
			"status":    node.Status.String(),
			"createdAt": node.CreatedAt.AsTime().Local().Format("02/01/2006 15:04"),
			"updatedAt": node.UpdatedAt.AsTime().Local().Format("02/01/2006 15:04"),
		}
	}

	util.PrintTable(columns, rows, true)
	return nil
}
