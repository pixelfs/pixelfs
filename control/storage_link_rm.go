package control

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	pb "github.com/pixelfs/pixelfs/gen/pixelfs/v1"
)

func (p *PixelFS) StorageLinkRm(storageLinkId string) error {
	_, err := p.Core.StorageService.RemoveStorageLink(
		context.Background(),
		connect.NewRequest(&pb.RemoveStorageLinkRequest{
			StorageLinkId: storageLinkId,
		}),
	)
	if err != nil {
		return err
	}

	fmt.Println("storage link removed successfully")
	return nil
}
