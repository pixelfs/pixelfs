package pixelfsd

import (
	"context"
	"os"
	"path/filepath"
	"time"

	"github.com/pixelfs/pixelfs/config"
	"github.com/pixelfs/pixelfs/log"
	"github.com/pixelfs/pixelfs/util"
)

func CleanFFmpegCache(cfg *config.Config) {
	task, err := util.NewTask(
		"clean:cache:ffmpeg",
		func(task *util.Task) {
			dirs, err := os.ReadDir(cfg.FFmpeg.Cache.Path)
			if err != nil {
				return
			}

			for _, dir := range dirs {
				info, err := dir.Info()
				if err != nil {
					return
				}

				name := filepath.Join(cfg.FFmpeg.Cache.Path, dir.Name())
				if time.Since(info.ModTime()) > time.Duration(cfg.FFmpeg.Cache.Expire)*time.Second {
					log.Info().Str("path", name).Msg("removing file")
					if err := os.RemoveAll(name); err != nil {
						log.Error().Err(err).Msg("failed to remove file")
					}
				}
			}
		},
		1*time.Hour,
	)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create ffmpeg clean task")
	}

	go task.Run(context.Background())
	log.Debug().Msg("ffmpeg cache clean task started")
}
