package fs

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"connectrpc.com/connect"
	"github.com/lesismal/arpc"
	"github.com/pixelfs/pixelfs/config"
	pb "github.com/pixelfs/pixelfs/gen/pixelfs/v1"
	"github.com/pixelfs/pixelfs/log"
	"github.com/pixelfs/pixelfs/util"
)

func M3U8(ctx *arpc.Context) {
	var request pb.FileM3U8Request
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

	metadata, err := util.GetFFmpegMetadata(absolutePath)
	if handleError(ctx, err) {
		return
	}

	if metadata.Duration == 0 {
		ctx.Error(fmt.Errorf("invalid duration"))
		return
	}

	if request.BlockSettings.Width > metadata.Width || request.BlockSettings.Height > metadata.Height {
		ctx.Error(fmt.Errorf("invalid resolution: %dx%d detected in stream", metadata.Width, metadata.Height))
		return
	}

	cfg, err := config.GetConfig()
	if handleError(ctx, err) {
		return
	}

	blockDuration := location.Msg.Location.BlockDuration
	cacheDir := filepath.Join(cfg.FFmpeg.Cache.Path, fmt.Sprintf("%s_%d", hash, blockDuration))
	_, err = os.Stat(cacheDir)
	if err != nil && os.IsNotExist(err) {
		err := util.GenerateSegmentFiles(absolutePath, cacheDir, blockDuration)
		if handleError(ctx, err) {
			return
		}
	}

	// set block duration
	err = setBlockDuration(&request, hash, cacheDir, blockDuration)
	if handleError(ctx, err) {
		return
	}

	if err := ctx.Write(&pb.FileM3U8Response{Duration: metadata.Duration}); err != nil {
		log.Error().Caller().Err(err).Msg("write response")
	}
}

func setBlockDuration(request *pb.FileM3U8Request, hash, cacheDir string, blockDuration int64) error {
	blocks, err := rpc.BlockService.GetBlockDuration(
		context.Background(),
		connect.NewRequest(&pb.GetBlockDurationRequest{
			NodeId:        request.Context.NodeId,
			Hash:          hash,
			BlockDuration: blockDuration,
		}),
	)
	if err != nil {
		return err
	}

	if len(blocks.Msg.Data) > 0 {
		return nil
	}

	data, err := parsePlaylistData(filepath.Join(cacheDir, "playlist.m3u8"))
	if err != nil {
		return err
	}

	_, err = rpc.BlockService.SetBlockDuration(
		context.Background(),
		connect.NewRequest(&pb.SetBlockDurationRequest{
			NodeId:   request.Context.NodeId,
			Location: request.Context.Location,
			Path:     request.Context.Path,
			Data:     data,
		}),
	)
	if err != nil {
		return err
	}

	return nil
}

func parsePlaylistData(filePath string) (map[int64]float64, error) {
	tsDurations := make(map[int64]float64)

	extinfRegex := regexp.MustCompile(`^#EXTINF:([0-9.]+),`)
	tsFileRegex := regexp.MustCompile(`^([0-9]+)\.ts$`)

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var currentDuration float64

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if matches := extinfRegex.FindStringSubmatch(line); matches != nil {
			duration, err := strconv.ParseFloat(matches[1], 64)
			if err != nil {
				return nil, err
			}
			currentDuration = duration
		} else if matches := tsFileRegex.FindStringSubmatch(line); matches != nil {
			index, err := strconv.Atoi(matches[1])
			if err != nil {
				return nil, err
			}

			tsDurations[int64(index)] = currentDuration
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return tsDurations, nil
}
