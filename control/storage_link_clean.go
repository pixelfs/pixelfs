package control

import (
	"context"
	"fmt"
	"time"

	"connectrpc.com/connect"
	"github.com/briandowns/spinner"
	pb "github.com/pixelfs/pixelfs/gen/pixelfs/v1"
)

func (p *PixelFS) StorageLinkClean(storageLinkId string) error {
	s := spinner.New(spinner.CharSets[31], 100*time.Millisecond)
	s.Start()

	_, err := p.Core.StorageService.CleanStorageLink(
		context.Background(),
		connect.NewRequest(&pb.CleanStorageLinkRequest{
			StorageLinkId: storageLinkId,
		}),
	)
	if err != nil {
		s.Stop()
		return err
	}

	s.Stop()
	fmt.Println("storage link cleaned successfully")
	return nil
}
