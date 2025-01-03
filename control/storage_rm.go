package control

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	pb "github.com/pixelfs/pixelfs/gen/pixelfs/v1"
)

func (p *PixelFS) StorageRm(id string) error {
	_, err := p.Core.StorageService.RemoveStorage(
		context.Background(),
		connect.NewRequest(&pb.RemoveStorageRequest{
			StorageId: id,
		}),
	)
	if err != nil {
		return err
	}

	fmt.Println("storage removed successfully")
	return nil
}
