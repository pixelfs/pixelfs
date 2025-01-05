package webdav

import (
	"context"
	"os"
	"path/filepath"
	"time"

	"github.com/pixelfs/pixelfs/config"
	"github.com/pixelfs/pixelfs/log"
	"github.com/pixelfs/pixelfs/util"
)

func cleanWebdavCache(cfg *config.Config) {
	task, err := util.NewTask(
		"clean:cache:webdav",
		func(task *util.Task) {
			_ = filepath.Walk(cfg.Webdav.Cache.Path, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}

				if path == cfg.Webdav.Cache.Path {
					return nil
				}

				if time.Since(info.ModTime()) > time.Duration(cfg.Webdav.Cache.Expire)*time.Second {
					log.Info().Str("path", path).Msg("removing file")
					if err := os.RemoveAll(path); err != nil {
						log.Error().Err(err).Msg("failed to remove file")
					}
				}

				return nil
			})
		},
		1*time.Hour,
	)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create webdav clean task")
	}

	go task.Run(context.Background())
	log.Debug().Msg("webdav cache clean task started")
}
