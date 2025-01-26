package control

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"connectrpc.com/connect"
	pb "github.com/pixelfs/pixelfs/gen/pixelfs/v1"
	"github.com/pixelfs/pixelfs/util"
	"path/filepath"
)

type blockResult struct {
	index int64
	data  []byte
	err   error
}

func (p *PixelFS) Download(ctx *pb.FileContext, output string, thread int) error {
	stat, err := p.Core.FileSystemService.Stat(
		context.Background(),
		connect.NewRequest(&pb.FileStatRequest{
			Context: ctx,
		}),
	)
	if err != nil {
		return err
	}

	_, fileName := filepath.Split(ctx.Path)
	if output == "" {
		dir, err := os.Getwd()
		if err != nil {
			return err
		}

		output = filepath.Join(dir, fileName)
	}

	if stat.Msg.File.Type == pb.FileType_DIR {
		if err := os.MkdirAll(output, 0755); err != nil {
			return err
		}

		list, err := p.Core.FileSystemService.List(
			context.Background(),
			connect.NewRequest(&pb.FileListRequest{
				Context: ctx,
			}),
		)
		if err != nil {
			return err
		}

		for _, file := range list.Msg.Files {
			if err := p.Download(
				&pb.FileContext{
					NodeId:   ctx.NodeId,
					Location: ctx.Location,
					Path:     filepath.Join(ctx.Path, file.Name),
				},
				filepath.Join(output, file.Name),
				thread,
			); err != nil {
				return err
			}
		}

		return nil
	}

	location, err := p.Core.LocationService.GetLocationByContext(
		context.Background(),
		connect.NewRequest(&pb.GetLocationByContextRequest{
			Context: ctx,
		}),
	)
	if err != nil {
		return err
	}

	blockCount := stat.Msg.File.Size / location.Msg.Location.BlockSize
	bar := util.NewProgressBar(int(blockCount+1), fmt.Sprintf("Downloading %s", ctx.Path))

	file, err := os.OpenFile(output, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open output file: %w", err)
	}
	defer file.Close()

	if err = bar.RenderBlank(); err != nil {
		return err
	}

	var wg sync.WaitGroup
	ch := make(chan int64, thread)
	resultCh := make(chan blockResult)

	go func() {
		for index := int64(0); index <= blockCount; index++ {
			ch <- index
		}
		close(ch)
	}()

	for i := 0; i < thread; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for index := range ch {
				var err error
				var read *connect.Response[pb.FileReadResponse]

				for retries := 0; retries < 20; retries++ {
					read, err = p.Core.FileSystemService.Read(
						context.Background(),
						connect.NewRequest(&pb.FileReadRequest{
							Context:    ctx,
							BlockType:  pb.BlockType_SIZE,
							BlockIndex: index,
						}),
					)

					if err != nil {
						resultCh <- blockResult{index: index, err: fmt.Errorf("failed to read block %d: %w", index, err)}
						return
					}

					if read.Msg.BlockStatus != pb.BlockStatus_PENDING {
						break
					}

					time.Sleep(5 * time.Second)
				}

				if read == nil || read.Msg.BlockStatus == pb.BlockStatus_PENDING {
					resultCh <- blockResult{index: index, err: fmt.Errorf("block %d is still pending after retries", index)}
					return
				}

				resp, err := util.Resty.R().Get(read.Msg.Url)
				if err != nil {
					resultCh <- blockResult{index: index, err: fmt.Errorf("failed to download block %d: %w", index, err)}
					return
				}

				resultCh <- blockResult{index: index, data: resp.Body()}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	var writeIndex int64
	results := make(map[int64][]byte)

	for res := range resultCh {
		if res.err != nil {
			return res.err
		}

		results[res.index] = res.data
		for {
			data, ok := results[writeIndex]
			if !ok {
				break
			}

			if _, err := file.Write(data); err != nil {
				return fmt.Errorf("failed to write to destination file: %w", err)
			}

			bar.Add(1)
			delete(results, writeIndex)
			writeIndex++
		}
	}

	return nil
}
