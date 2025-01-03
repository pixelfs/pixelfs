package control

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"time"

	"connectrpc.com/connect"
	"github.com/briandowns/spinner"
	pb "github.com/pixelfs/pixelfs/gen/pixelfs/v1"
)

func (p *PixelFS) M3U8(ctx *pb.FileContext, width, height, bitrate int) error {
	s := spinner.New(spinner.CharSets[31], 100*time.Millisecond)
	s.Start()

	response, err := p.Core.FileSystemService.M3U8(
		context.Background(),
		connect.NewRequest(&pb.FileM3U8Request{
			Context: ctx,
			BlockSettings: &pb.BlockSettings{
				Width:   int32(width),
				Height:  int32(height),
				Bitrate: int32(bitrate),
			},
		}),
	)
	if err != nil {
		s.Stop()
		return err
	}

	s.Stop()

	signature, err := extractSignature(response.Msg.Url)
	if err != nil {
		return err
	}

	fmt.Println("Play url, open in Browser:")
	fmt.Println("")
	fmt.Println(p.cfg.Endpoint + "/player/" + signature)

	fmt.Println("")

	fmt.Println("M3U8 playlist generated at:")
	fmt.Println("")
	fmt.Println(response.Msg.Url)
	return nil
}

func extractSignature(m3u8 string) (string, error) {
	parsedURL, err := url.Parse(m3u8)
	if err != nil {
		return "", fmt.Errorf("failed to parse url: %w", err)
	}

	parts := strings.Split(parsedURL.Path, "/")
	return parts[2], nil
}
