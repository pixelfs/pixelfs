package control

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	pb "github.com/pixelfs/pixelfs/gen/pixelfs/v1"
)

func (p *PixelFS) LocationAdd(nodeId, name, typ, path string, blockSize, blockDuration int64) error {
	var locationType pb.LocationType
	switch typ {
	case "local":
		locationType = pb.LocationType_LOCAL
	default:
		return fmt.Errorf("invalid location type: %s", typ)
	}

	_, err := p.Core.LocationService.AddLocation(
		context.Background(),
		connect.NewRequest(&pb.AddLocationRequest{
			Location: &pb.Location{
				NodeId:        nodeId,
				Name:          name,
				Type:          locationType,
				Path:          path,
				BlockSize:     blockSize,
				BlockDuration: blockDuration,
			},
		}),
	)
	if err != nil {
		return err
	}

	fmt.Println("location added successfully")
	return nil
}
