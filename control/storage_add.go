package control

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	pb "github.com/pixelfs/pixelfs/gen/pixelfs/v1"
	"google.golang.org/protobuf/proto"
)

func (p *PixelFS) StorageAdd(name string, config proto.Message, network pb.StorageNetwork) error {
	var storage *pb.Storage
	switch c := config.(type) {
	case *pb.StorageS3Config:
		storage = &pb.Storage{
			Name:    name,
			Type:    pb.StorageType_S3,
			Network: network,
			Config:  &pb.Storage_S3{S3: c},
		}
	default:
		return fmt.Errorf("unsupported storage config type: %T", c)
	}

	_, err := p.Core.StorageService.AddStorage(
		context.Background(),
		connect.NewRequest(&pb.AddStorageRequest{
			Storage: storage,
		}),
	)
	if err != nil {
		return err
	}

	fmt.Println("storage added successfully")
	return nil
}
