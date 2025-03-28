package sync

import (
	"context"
	"os"
	"runtime"
	"time"

	"connectrpc.com/connect"
	"github.com/fsnotify/fsnotify"
	pb "github.com/pixelfs/pixelfs/gen/pixelfs/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (fs *FileSync) handleWatcherEvent(event fsnotify.Event, platform string, destContext *pb.FileContext, watcher *fsnotify.Watcher) error {
	switch {
	case event.Has(fsnotify.Create):
		return fs.handleCreate(event.Name, platform, destContext, watcher)
	case event.Has(fsnotify.Write):
		return fs.handleWrite(event.Name, platform, destContext)
	case event.Has(fsnotify.Remove), event.Has(fsnotify.Rename):
		return fs.handleRemove(event.Name, destContext)
	case event.Has(fsnotify.Chmod):
		return fs.handleChmod(event.Name, destContext)
	default:
		return nil
	}
}

func (fs *FileSync) handleCreate(name, platform string, destContext *pb.FileContext, watcher *fsnotify.Watcher) error {
	fileInfo, err := os.Stat(name)
	if err != nil {
		return err
	}

	if fileInfo.IsDir() {
		if err = watcher.Add(name); err != nil {
			return err
		}

		if runtime.GOOS == "windows" {
			_, err = fs.Core.FileSystemService.Mkdir(
				context.Background(),
				connect.NewRequest(&pb.FileMkdirRequest{
					Context: destContext,
					Mtime:   timestamppb.New(fileInfo.ModTime()),
				}),
			)

			return err
		}
	}

	if runtime.GOOS != "windows" {
		if err = fs.copyFile(name, destContext, platform); err != nil {
			return err
		}
	}

	return nil
}

func (fs *FileSync) handleWrite(name, platform string, destContext *pb.FileContext) error {
	// hack for windows, wait for file unlock
	if runtime.GOOS == "windows" {
		time.Sleep(100 * time.Millisecond)
	}

	return fs.copyFile(name, destContext, platform)
}

func (fs *FileSync) handleRemove(name string, destContext *pb.FileContext) error {
	fs.lockFile(name)
	defer fs.unlockFile(name)

	if _, err := fs.stat(destContext); err != nil {
		if os.IsNotExist(err) {
			return nil // 没有错误日志，直接返回
		}
		return err
	}

	return fs.remove(destContext)
}

func (fs *FileSync) handleChmod(name string, destContext *pb.FileContext) error {
	fs.lockFile(name)
	defer fs.unlockFile(name)

	fileInfo, err := os.Stat(name)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // 没有错误日志，直接返回
		}
		return err
	}

	destInfo, err := fs.stat(destContext)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // 没有错误日志，直接返回
		}
		return err
	}
	if destInfo.Platform == "windows" {
		return nil
	}
	if destInfo.Mode == 0 || destInfo.Mode == uint32(fileInfo.Mode()) {
		return nil
	}

	return fs.chmod(destContext, fileInfo.Mode())
}
