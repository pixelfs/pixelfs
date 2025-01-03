package control

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	pb "github.com/pixelfs/pixelfs/gen/pixelfs/v1"
)

func (p *PixelFS) NodeRm(nodeId string) error {
	_, err := p.Core.NodeService.Remove(
		context.Background(),
		connect.NewRequest(&pb.NodeRemoveRequest{
			NodeId: nodeId,
		}),
	)
	if err != nil {
		return err
	}

	fmt.Println("node removed successfully")
	return nil
}
