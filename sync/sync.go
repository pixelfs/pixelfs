package sync

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"connectrpc.com/connect"
	"github.com/fsnotify/fsnotify"
	"github.com/pixelfs/pixelfs/config"
	pb "github.com/pixelfs/pixelfs/gen/pixelfs/v1"
	"github.com/pixelfs/pixelfs/log"
	"github.com/pixelfs/pixelfs/rpc/core"
	"github.com/pixelfs/pixelfs/util"
	gitignore "github.com/sabhiram/go-gitignore"
	"golang.org/x/sync/semaphore"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	tmpPrefix = ".pixelfstmp."
)

type Watcher struct {
	mu       sync.Mutex
	registry map[string]*fsnotify.Watcher
}

type FileSync struct {
	Core    *core.GrpcV1Client
	Lock    sync.Map
	NodeId  string
	watcher *Watcher
}

type ErrorReport struct {
	Op       string
	Name     string
	Platform string
	Dest     *pb.FileContext
}

func NewFileSync(cfg *config.Config) (*FileSync, error) {
	userInfo, err := core.NewGrpcV1Client(cfg).UserService.GetUserInfo(
		context.Background(),
		connect.NewRequest(&pb.GetUserInfoRequest{}),
	)
	if err != nil {
		return nil, err
	}

	nodeId, err := util.GetNodeId(userInfo.Msg.Id)
	if err != nil {
		return nil, err
	}

	return &FileSync{
		Core:   core.NewGrpcV1Client(cfg),
		NodeId: nodeId,
		watcher: &Watcher{
			mu:       sync.Mutex{},
			registry: make(map[string]*fsnotify.Watcher),
		},
	}, nil
}

func (fs *FileSync) Start(sync *pb.Sync) error {
	if !sync.Enabled {
		return errors.New("sync is disabled")
	}

	srcContext, destContext := fs.getFileContext(sync)
	if srcContext == nil || destContext == nil {
		return fmt.Errorf("node id not match")
	}

	syncLimit := int64(3)
	if sync.Config.Limit > 0 {
		syncLimit = sync.Config.Limit
	}

	srcDir, err := fs.getRealPathByContext(srcContext)
	if err != nil {
		return err
	}

	platform, err := fs.getPlatformByContext(destContext)
	if err != nil {
		return err
	}

	// do task
	scanTask := func() {
		_ = fs.updateStatus(sync.Id, pb.SyncStatus_SYNCING, "syncing files")

		if err := fs.scanFiles(srcDir, platform, destContext, syncLimit); err != nil {
			_ = fs.updateStatus(sync.Id, pb.SyncStatus_ERROR, err.Error())
			log.Error().Err(err).Msg("sync scan files task")
			return
		}

		_ = fs.updateStatus(sync.Id, pb.SyncStatus_SUCCESS, "")
	}

	task, err := util.NewTask(fs.getTaskId(sync.Id), func(task *util.Task) { scanTask() },
		time.Duration(sync.Config.Interval)*time.Second,
	)
	if err != nil {
		return err
	}

	// run scan task
	go task.Run(context.Background())
	go scanTask()

	// watch files
	return fs.watcherFiles(sync, srcDir, platform, destContext, syncLimit)
}

func (fs *FileSync) Stop(sync *pb.Sync) error {
	if err := util.StopTask(fs.getTaskId(sync.Id)); err != nil {
		return err
	}

	fs.watcher.mu.Lock()
	defer fs.watcher.mu.Unlock()

	existingWatcher, exists := fs.watcher.registry[sync.Id]
	if exists {
		existingWatcher.Close()
		delete(fs.watcher.registry, sync.Id)
	}

	return nil
}

func (fs *FileSync) watcherFiles(sync *pb.Sync, srcDir, platform string, dest *pb.FileContext, limit int64) error {
	fs.watcher.mu.Lock()
	defer fs.watcher.mu.Unlock()

	existingWatcher, exists := fs.watcher.registry[sync.Id]
	if exists {
		existingWatcher.Close()
		delete(fs.watcher.registry, sync.Id)
	}

	sem := semaphore.NewWeighted(limit)
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}

	// add watcher to registry
	fs.watcher.registry[sync.Id] = watcher

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if fs.isIgnore(event.Name) {
					continue
				}

				relPath, err := filepath.Rel(srcDir, event.Name)
				if err != nil {
					continue
				}

				destContext := &pb.FileContext{
					NodeId:   dest.NodeId,
					Location: dest.Location,
					Path:     util.JoinPath(platform, dest.Path, relPath),
				}

				go func() {
					_ = sem.Acquire(context.Background(), 1)
					defer sem.Release(1)

					if err := fs.handleWatcherEvent(event, platform, destContext, watcher); err != nil {
						_ = fs.writeErrorReport(&ErrorReport{
							Op:       event.Op.String(),
							Name:     event.Name,
							Platform: platform,
							Dest:     destContext,
						})

						log.Error().
							Str("event", event.Op.String()).
							Str("name", event.Name).Err(err).Msg("sync handle watcher event")
					}
				}()
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Error().Err(err)
			}
		}
	}()

	return filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return watcher.Add(path)
		}

		return nil
	})
}

func (fs *FileSync) scanFiles(srcDir, platform string, dest *pb.FileContext, limit int64) error {
	var wg sync.WaitGroup

	sem := semaphore.NewWeighted(limit)
	errChan := make(chan error, limit)

	// cancel context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// handle error report
	if err := fs.handleErrorReport(func(report *ErrorReport) error {
		switch report.Op {
		case "REMOVE", "RENAME":
			return fs.handleRemove(report.Name, report.Dest)
		case "CHMOD":
			return fs.handleChmod(report.Name, report.Dest)
		default:
			return nil
		}
	}); err != nil {
		return err
	}

	err := filepath.Walk(srcDir, func(srcPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if srcDir == srcPath {
			return nil
		}
		if fs.isIgnore(srcPath) {
			return nil
		}
		if ctx.Err() != nil {
			return fmt.Errorf("task canceled")
		}

		relPath, err := filepath.Rel(srcDir, srcPath)
		if err != nil {
			return err
		}

		wg.Add(1)
		go func(src string, dest *pb.FileContext) {
			defer wg.Done()

			_ = sem.Acquire(context.Background(), 1)
			defer sem.Release(1)

			if err = fs.copyFile(src, dest, platform); err != nil {
				select {
				case errChan <- err:
					cancel()
				default:
				}
			}
		}(srcPath, &pb.FileContext{
			NodeId:   dest.NodeId,
			Location: dest.Location,
			Path:     util.JoinPath(platform, dest.Path, relPath),
		})

		return nil
	})

	wg.Wait()
	close(errChan)

	for err = range errChan {
		if err != nil {
			return err
		}
	}

	return err
}

func (fs *FileSync) copyFile(src string, dest *pb.FileContext, platform string) error {
	fs.lockFile(src)
	defer fs.unlockFile(src)

	fileInfo, err := os.Stat(src)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // 没有错误日志，直接返回
		}
		return err
	}

	if fileInfo.IsDir() {
		if _, err = fs.stat(dest); err != nil && errors.Is(err, os.ErrNotExist) {
			_, err = fs.Core.FileSystemService.Mkdir(
				context.Background(),
				connect.NewRequest(&pb.FileMkdirRequest{
					Context: dest,
					Mtime:   timestamppb.New(fileInfo.ModTime()),
				}),
			)
		}

		return err
	}

	hash, err := util.GetFileHash(src)
	if err != nil {
		return err
	}

	destInfo, err := fs.stat(dest)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return fs.upload(src, hash, dest, platform)
		}

		return err
	}

	if hash == destInfo.Hash {
		return nil
	}
	return fs.upload(src, hash, dest, platform)
}

func (fs *FileSync) updateStatus(syncId string, status pb.SyncStatus, log string) error {
	_, err := fs.Core.SyncService.UpdateStatus(
		context.Background(),
		connect.NewRequest(&pb.SyncUpdateStatusRequest{
			SyncId: syncId,
			Status: status,
			Log:    log,
		}),
	)

	return err
}

func (fs *FileSync) getRealPathByContext(ctx *pb.FileContext) (string, error) {
	location, err := fs.Core.LocationService.GetLocationByContext(
		context.Background(),
		connect.NewRequest(&pb.GetLocationByContextRequest{
			Context: ctx,
		}),
	)
	if err != nil {
		return "", err
	}

	return filepath.Join(location.Msg.Location.Path, ctx.Path), nil
}

func (fs *FileSync) getFileContext(sync *pb.Sync) (*pb.FileContext, *pb.FileContext) {
	var srcContext *pb.FileContext
	var destContext *pb.FileContext

	if sync.SrcContext.NodeId == fs.NodeId {
		srcContext = sync.SrcContext
		destContext = sync.DestContext
	} else if sync.DestContext.NodeId == fs.NodeId {
		srcContext = sync.DestContext
		destContext = sync.SrcContext
	}

	return srcContext, destContext
}

func (fs *FileSync) getPlatformByContext(ctx *pb.FileContext) (string, error) {
	fileInfo, err := fs.Core.FileSystemService.Stat(
		context.Background(),
		connect.NewRequest(&pb.FileStatRequest{
			Context: ctx,
		}),
	)
	if err != nil {
		return "", err
	}

	return fileInfo.Msg.File.Platform, nil
}

func (fs *FileSync) getTaskId(syncId string) string {
	return "sync:" + syncId
}

func (fs *FileSync) isIgnore(path string) bool {
	ignorePatterns := []string{tmpPrefix + "*", ".DS_Store"}

	if strings.Count(filepath.ToSlash(path), "/") > 20 {
		return true
	}

	// check .pixelfsignore
	dir := path
	for {
		dir = filepath.Dir(dir)
		ignoreFile := filepath.Join(dir, ".pixelfsignore")
		if content, err := os.ReadFile(ignoreFile); err == nil {
			ignorePatterns = append(ignorePatterns, strings.Split(strings.TrimSpace(string(content)), "\n")...)
		}

		if dir == "/" || dir == filepath.VolumeName(dir)+`\` {
			break
		}
	}

	ignoreMatcher := gitignore.CompileIgnoreLines(ignorePatterns...)
	return ignoreMatcher.MatchesPath(path)
}

func (fs *FileSync) handleErrorReport(handleFunc func(*ErrorReport) error) error {
	home, err := util.GetHomeDir()
	if err != nil {
		return err
	}

	logPath := filepath.Join(home, "sync_error_report")
	tmpPath := logPath + ".tmp"

	file, err := os.Open(logPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // 没有错误日志，直接返回
		}
		return err
	}
	defer file.Close()

	tmpFile, err := os.OpenFile(tmpPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer tmpFile.Close()

	scanner := bufio.NewScanner(file)
	var hasProcessed bool

	for scanner.Scan() {
		var entry ErrorReport
		line := scanner.Text()

		if err := json.Unmarshal([]byte(line), &entry); err != nil {
			_, _ = tmpFile.WriteString(line + "\n")
			continue
		}

		if err := handleFunc(&entry); err != nil {
			_, _ = tmpFile.WriteString(line + "\n")
		} else {
			hasProcessed = true
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	if hasProcessed {
		if err := os.Rename(tmpPath, logPath); err != nil {
			return err
		}
	} else {
		_ = os.Remove(tmpPath)
	}

	return nil
}

func (fs *FileSync) writeErrorReport(entry *ErrorReport) error {
	jsonData, _ := json.Marshal(entry)
	jsonData = append(jsonData, '\n')

	home, err := util.GetHomeDir()
	if err != nil {
		return err
	}

	file, err := os.OpenFile(filepath.Join(home, "sync_error_report"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	return err
}

func (fs *FileSync) lockFile(name string) {
	value, _ := fs.Lock.LoadOrStore(name, make(chan struct{}, 1))
	lock := value.(chan struct{})

	lock <- struct{}{}
}

func (fs *FileSync) unlockFile(name string) {
	value, ok := fs.Lock.Load(name)
	if ok {
		lock := value.(chan struct{})

		<-lock
		fs.Lock.Delete(name)
	}
}
