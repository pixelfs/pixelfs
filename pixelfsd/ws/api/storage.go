package api

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/lesismal/arpc"
	pb "github.com/pixelfs/pixelfs/gen/pixelfs/v1"
	"github.com/pixelfs/pixelfs/log"
)

func StorageValidate(ctx *arpc.Context) {
	var request pb.StorageValidateRequest
	if err := ctx.Bind(&request); err != nil {
		ctx.Error(err)
		return
	}

	err := validateStorageConfig(request.Storage)
	if err != nil {
		ctx.Error(err)
		return
	}

	if err := ctx.Write(&pb.StorageValidateResponse{}); err != nil {
		log.Error().Caller().Err(err).Msg("write response")
	}
}

func StorageRemoveBlock(ctx *arpc.Context) {
	var request pb.StorageRemoveBlockRequest
	if err := ctx.Bind(&request); err != nil {
		ctx.Error(err)
		return
	}

	err := removeStorageBlock(request.Storage, request.Path)
	if err != nil {
		ctx.Error(err)
		return
	}

	if err := ctx.Write(&pb.StorageRemoveBlockResponse{}); err != nil {
		log.Error().Caller().Err(err).Msg("write response")
	}
}

func validateStorageConfig(storage *pb.Storage) error {
	switch c := storage.Config.(type) {
	case *pb.Storage_S3:
		client, err := getS3Client(c.S3)
		if err != nil {
			return err
		}

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		_, err = client.HeadBucket(ctx, &s3.HeadBucketInput{Bucket: aws.String(c.S3.Bucket)})
		if err != nil {
			return fmt.Errorf("failed to validate S3 bucket: %w", err)
		}
	default:
		return fmt.Errorf("unsupported storage config type: %T", c)
	}
	return nil
}

func removeStorageBlock(storage *pb.Storage, path string) error {
	switch c := storage.Config.(type) {
	case *pb.Storage_S3:
		client, err := getS3Client(c.S3)
		if err != nil {
			return err
		}

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		_, err = client.DeleteObject(ctx, &s3.DeleteObjectInput{Bucket: aws.String(c.S3.Bucket), Key: aws.String(path)})
		if err != nil {
			return fmt.Errorf("failed to delete S3 object: %w", err)
		}
	default:
		return fmt.Errorf("unsupported storage config type: %T", c)
	}
	return nil
}

func getS3Client(storage *pb.StorageS3Config) (*s3.Client, error) {
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(storage.AccessKey, storage.SecretKey, "")),
		config.WithRegion(storage.Region),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to load S3 config: %w", err)
	}

	return s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(storage.Endpoint)
		o.UsePathStyle = storage.PathStyle
	}), nil
}
