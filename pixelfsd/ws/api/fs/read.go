package fs

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"connectrpc.com/connect"
	"github.com/lesismal/arpc"
	"github.com/pixelfs/pixelfs/config"
	pb "github.com/pixelfs/pixelfs/gen/pixelfs/v1"
	"github.com/pixelfs/pixelfs/log"
	"github.com/pixelfs/pixelfs/util"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func Read(ctx *arpc.Context) {
	var request pb.FileReadRequest
	if err := ctx.Bind(&request); handleError(ctx, err) {
		return
	}

	location, err := rpc.LocationService.GetLocationByContext(
		context.Background(),
		connect.NewRequest(&pb.GetLocationByContextRequest{
			Context: request.Context,
		}),
	)
	if handleError(ctx, err) {
		return
	}

	absolutePath := filepath.Join(location.Msg.Location.Path, request.Context.Path)
	file, err := os.Open(absolutePath)
	if handleError(ctx, err) {
		return
	}
	defer file.Close()

	hash, err := util.GetFileHash(absolutePath)
	if handleError(ctx, err) {
		return
	}

	fileInfo, err := file.Stat()
	if handleError(ctx, err) || fileInfo.IsDir() {
		ctx.Error(fmt.Errorf("%s is a directory", request.Context.Path))
		return
	}

	buffer, err := readBlock(absolutePath, file, fileInfo, hash, &request, location.Msg.Location)
	if handleError(ctx, err) {
		return
	}

	blockSize := int64(len(buffer))
	storage, err := rpc.StorageService.Upload(
		context.Background(),
		connect.NewRequest(&pb.StorageUploadRequest{
			Context:       request.Context,
			Hash:          hash,
			BlockType:     request.BlockType,
			BlockIndex:    request.BlockIndex,
			BlockSize:     blockSize,
			BlockSettings: request.BlockSettings,
		}),
	)
	if handleError(ctx, err) {
		return
	}

	if storage.Msg.Type != pb.StorageType_S3 {
		ctx.Error(fmt.Errorf("unsupported storage type: %s", storage.Msg.Type))
		return
	}

	if _, err = util.Resty.R().SetBody(buffer).Put(storage.Msg.Url); handleError(ctx, err) {
		return
	}

	if err := ctx.Write(&pb.FileReadResponse{
		BlockId:   storage.Msg.BlockId,
		BlockSize: blockSize,
	}); err != nil {
		log.Error().Caller().Err(err).Msg("write response")
	}
}

func readBlock(path string, file *os.File, fileInfo os.FileInfo, hash string, request *pb.FileReadRequest, location *pb.Location) ([]byte, error) {
	switch request.BlockType {
	case pb.BlockType_SIZE:
		return readBlockSize(file, fileInfo.Size(), request.BlockIndex, location.BlockSize)
	case pb.BlockType_DURATION:
		return readBlockDuration(path, hash, request.BlockIndex, location.BlockDuration, request.BlockSettings)
	default:
		return nil, fmt.Errorf("unsupported block type: %s", request.BlockType)
	}
}

func readBlockSize(file *os.File, fileSize, blockIndex, blockSize int64) ([]byte, error) {
	offset := blockIndex * blockSize
	if offset >= fileSize {
		offset = fileSize - 1
	}

	if offset+blockSize > fileSize {
		blockSize = fileSize - offset
	}

	buffer := make([]byte, blockSize)
	if _, err := file.ReadAt(buffer, offset); err != nil && err != io.EOF {
		return nil, err
	}

	return buffer, nil
}

func readBlockDuration(input, hash string, blockIndex, blockDuration int64, blockSettings *pb.BlockSettings) ([]byte, error) {
	cfg, err := config.GetConfig()
	if err != nil {
		return nil, err
	}

	metadata, err := util.GetFFmpegMetadata(input)
	if err != nil {
		return nil, err
	}

	if metadata.Duration == 0 {
		return nil, fmt.Errorf("failed to get duration")
	}

	if blockIndex*blockDuration >= int64(metadata.Duration) {
		return nil, fmt.Errorf("block index out of range")
	}

	cacheDir := filepath.Join(cfg.FFmpeg.Cache.Path, fmt.Sprintf("%s_%d", hash, blockDuration))
	cacheFileName := filepath.Join(cacheDir, util.PadLeft(fmt.Sprintf("%d", blockIndex), 5, "0")+".ts")
	_, err = os.Stat(cacheFileName)
	if err != nil && os.IsNotExist(err) {
		err := util.GenerateSegmentFiles(input, cacheDir, blockDuration)
		if err != nil {
			return nil, err
		}
	}

	ouputFile := generateOutputFileName(hash, blockIndex, blockDuration, blockSettings)

	encoder, err := util.GetFFmpegEncoder()
	if err != nil {
		log.Error().Err(err).Msg("get ffmpeg encoder")
	}

	// build ffmpeg args
	outputArgs := buildOutputFFmpegArgs(encoder, metadata, blockSettings, cfg.FFmpeg.Options)
	inputArgs := buildInputFFmpegArgs(encoder, outputArgs)

	if err = ffmpeg.Input(cacheFileName, inputArgs).Output(ouputFile, outputArgs).OverWriteOutput().Run(); err != nil {
		return nil, err
	}

	buffer, err := os.ReadFile(ouputFile)
	if err != nil {
		return nil, err
	}
	if err := os.Remove(ouputFile); err != nil {
		log.Error().Err(err).Msg("remove output")
	}

	return buffer, nil
}

func buildInputFFmpegArgs(encoder *util.FFmpegEncoder, outputArgs ffmpeg.KwArgs) ffmpeg.KwArgs {
	args := ffmpeg.KwArgs{}
	if encoder != nil && encoder.HWAccel == "vaapi" && outputArgs["codec:v:0"] == encoder.Encoder {
		args["hwaccel"] = encoder.HWAccel
		args["hwaccel_device"] = "/dev/dri/renderD128"
		args["hwaccel_output_format"] = encoder.HWAccel
	}

	return args
}

func buildOutputFFmpegArgs(encoder *util.FFmpegEncoder, metadata *util.FFmpegMetadata, blockSettings *pb.BlockSettings, options map[string]any) ffmpeg.KwArgs {
	args := ffmpeg.KwArgs{
		"codec:v:0": "libx264",
		"codec:a:0": "copy",
		"f":         "mpegts",
		"copyts":    "",
	}

	if encoder != nil && encoder.Encoder != "" {
		args["codec:v:0"] = encoder.Encoder
	}

	if blockSettings != nil {
		if blockSettings.Width != 0 || blockSettings.Height != 0 {
			scale := "scale="
			if encoder != nil && encoder.HWAccel == "vaapi" {
				scale = "format=vaapi,hwupload,scale_vaapi="
			}

			if blockSettings.Width != 0 {
				scale += fmt.Sprintf("w=%d:", blockSettings.Width)
			} else {
				scale += "w=-1:"
			}
			if blockSettings.Height != 0 {
				scale += fmt.Sprintf("h=%d", blockSettings.Height)
			} else {
				scale += "h=-1"
			}
			args["vf"] = scale
		}
		if blockSettings.Bitrate != 0 {
			args["b:v:0"] = fmt.Sprintf("%dk", blockSettings.Bitrate) // Video bitrate
		}
	}

	if _, hasFilters := args["vf"]; !hasFilters {
		if _, hasBitrate := args["b:v:0"]; !hasBitrate {
			if strings.EqualFold(metadata.VideoCodec, "h264") {
				args["codec:v:0"] = "copy"
			}
		}
	}

	if options != nil {
		for key, value := range options {
			args[key] = value
		}
	}

	return args
}

func generateOutputFileName(hash string, blockIndex, blockDuration int64, blockSettings *pb.BlockSettings) string {
	if blockSettings == nil {
		return fmt.Sprintf("%d", blockIndex)
	}

	parts := []string{}
	if blockSettings.Width != 0 {
		parts = append(parts, fmt.Sprintf("%d", blockSettings.Width))
	}
	if blockSettings.Height != 0 {
		parts = append(parts, fmt.Sprintf("%d", blockSettings.Height))
	}
	if blockSettings.Bitrate != 0 {
		parts = append(parts, fmt.Sprintf("%d", blockSettings.Bitrate))
	}

	parts = append(parts, fmt.Sprintf("%d", blockIndex))
	tmpDir := filepath.Join(os.TempDir(), "pixelfs", fmt.Sprintf("%s_%d", hash, blockDuration))
	_ = util.EnsureDir(tmpDir)

	return filepath.Join(tmpDir, fmt.Sprintf("%s.ts", strings.Join(parts, "_")))
}
