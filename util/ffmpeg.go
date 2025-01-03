package util

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/tidwall/gjson"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

type FFmpegMetadata struct {
	Duration   float64
	Width      int32
	Height     int32
	VideoCodec string
	AudioCodec string
}

type FFmpegEncoder struct {
	HWAccel string
	Encoder string
}

func GetFFmpegEncoder() (*FFmpegEncoder, error) {
	hwaccels, err := exec.Command("ffmpeg", "-hwaccels").Output()
	if err != nil {
		return nil, fmt.Errorf("failed to get ffmpeg hardware accelerations: %w", err)
	}

	encoders, err := exec.Command("ffmpeg", "-encoders").Output()
	if err != nil {
		return nil, fmt.Errorf("failed to get ffmpeg encoders: %w", err)
	}

	hwToEncoder := map[string]string{
		"videotoolbox": "h264_videotoolbox",
		"vaapi":        "h264_vaapi",
	}

	for hw, encoder := range hwToEncoder {

		if hw == "vaapi" {
			_, err = os.Stat("/dev/dri/renderD128")
			if err != nil {
				continue
			}
		}

		if bytes.Contains(hwaccels, []byte(hw)) {
			if bytes.Contains(encoders, []byte(encoder)) {
				return &FFmpegEncoder{HWAccel: hw, Encoder: encoder}, nil
			}

			return nil, nil
		}
	}

	return nil, nil
}

func GenerateSegmentFiles(input string, output string, blockDuration int64) error {
	metadata, err := GetFFmpegMetadata(input)
	if err != nil {
		return err
	}

	outputArgs := ffmpeg.KwArgs{
		"f":                         "hls",
		"hls_time":                  blockDuration,
		"hls_playlist_type":         "vod",
		"hls_segment_filename":      filepath.Join(output, "%05d.ts"),
		"codec:v:0":                 "copy",
		"codec:a:0":                 "copy",
		"start_number":              "0",
		"threads":                   "0",
		"map_metadata":              "-1",
		"map_chapters":              "-1",
		"individual_header_trailer": "0",
	}

	if !isBrowserSupportedAudio(metadata.AudioCodec) {
		outputArgs["codec:a:0"] = "libmp3lame"
	}

	if err := EnsureDir(output); err != nil {
		return err
	}
	if err := ffmpeg.Input(input).
		Output(filepath.Join(output, "playlist.m3u8"), outputArgs).
		OverWriteOutput().
		Run(); err != nil {
		return err
	}

	return nil
}

func GetFFmpegMetadata(input string) (*FFmpegMetadata, error) {
	metadata, err := ffmpeg.Probe(input)
	if err != nil {
		return nil, fmt.Errorf("failed to probe %s, please check if ffmpeg is installed, %s", input, err)
	}

	var videoCodec string
	var width, height int32
	gjson.Get(metadata, "streams").ForEach(func(key, value gjson.Result) bool {
		if value.Get("codec_type").String() == "video" {
			videoCodec = value.Get("codec_name").String()
			width = int32(value.Get("width").Int())
			height = int32(value.Get("height").Int())

			return false
		}

		return true
	})

	var audioCodec string
	gjson.Get(metadata, "streams").ForEach(func(key, value gjson.Result) bool {
		if value.Get("codec_type").String() == "audio" {
			audioCodec = value.Get("codec_name").String()
			return false
		}

		return true
	})

	return &FFmpegMetadata{
		Duration:   gjson.Get(metadata, "format.duration").Float(),
		Width:      width,
		Height:     height,
		VideoCodec: videoCodec,
		AudioCodec: audioCodec,
	}, nil
}

func isBrowserSupportedAudio(codec string) bool {
	for _, item := range []string{"mp3", "aac", "flac"} {
		if strings.EqualFold(item, codec) {
			return true
		}
	}
	return false
}
