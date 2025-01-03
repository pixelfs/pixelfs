package control

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	pb "github.com/pixelfs/pixelfs/gen/pixelfs/v1"
)

func (p *PixelFS) LocationRm(locationId string) error {
	_, err := p.Core.LocationService.RemoveLocation(
		context.Background(),
		connect.NewRequest(&pb.RemoveLocationRequest{
			LocationId: locationId,
		}),
	)
	if err != nil {
		return err
	}

	fmt.Println("location removed successfully")
	return nil
}
