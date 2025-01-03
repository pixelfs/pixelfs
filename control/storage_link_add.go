package control

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	pb "github.com/pixelfs/pixelfs/gen/pixelfs/v1"
)

func (p *PixelFS) StorageLinkAdd(storageId, nodeId, locationId string, limitSize int64) error {
	_, err := p.Core.StorageService.AddStorageLink(
		context.Background(),
		connect.NewRequest(&pb.AddStorageLinkRequest{
			StorageLink: &pb.StorageLink{
				StorageId:  storageId,
				NodeId:     nodeId,
				LocationId: locationId,
				LimitSize:  limitSize,
			},
		}),
	)
	if err != nil {
		return err
	}

	fmt.Println("storage link added successfully")
	return nil
}
