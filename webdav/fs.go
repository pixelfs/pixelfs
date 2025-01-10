package webdav

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"connectrpc.com/connect"
	"github.com/pixelfs/pixelfs/config"
	pb "github.com/pixelfs/pixelfs/gen/pixelfs/v1"
	"github.com/pixelfs/pixelfs/log"
	"github.com/pixelfs/pixelfs/rpc/core"
	"golang.org/x/net/webdav"
)

type fileSystem struct {
	webdav.FileSystem
	cfg  *config.Config
	core *core.GrpcV1Client
}

func newFileSystem(cfg *config.Config) *fileSystem {
	return &fileSystem{cfg: cfg, core: core.NewGrpcV1Client(cfg)}
}

func (fs *fileSystem) OpenFile(ctx context.Context, name string, flag int, perm os.FileMode) (webdav.File, error) {
	var location *pb.Location
	fileContext := fs.parseNameToContext(name)

	if flag&os.O_CREATE != 0 && flag&os.O_TRUNC != 0 {
		if fileContext.NodeId == "" || fileContext.Location == "" || fileContext.Path == "" {
			log.Error().Any("context", fileContext).Msg("invalid file context")
			return nil, fmt.Errorf("invalid file context")
		}

		_, err := fs.core.FileSystemService.Create(
			context.Background(),
			connect.NewRequest(&pb.FileCreateRequest{
				Context: fileContext,
			}),
		)
		if err != nil {
			return nil, err
		}
	}

	fileInfo, err := fs.stat(name)
	if err != nil {
		return nil, err
	}

	if fileContext != nil && fileContext.NodeId != "" && fileContext.Location != "" {
		response, err := fs.core.LocationService.GetLocationByContext(
			context.Background(),
			connect.NewRequest(&pb.GetLocationByContextRequest{
				Context: fileContext,
			}),
		)
		if err != nil {
			return nil, err
		}

		location = response.Msg.Location
	}

	request := ctx.Value(httpRequestKey).(*http.Request)
	return newFile(request, fs.cfg, fs.core, fileContext, fileInfo, location), nil
}

func (fs *fileSystem) Mkdir(ctx context.Context, name string, perm os.FileMode) error {
	_, err := fs.core.FileSystemService.Mkdir(
		context.Background(),
		connect.NewRequest(&pb.FileMkdirRequest{
			Context: fs.parseNameToContext(name),
		}),
	)
	if err != nil {
		return err
	}

	return nil
}

func (fs *fileSystem) RemoveAll(ctx context.Context, name string) error {
	_, err := fs.core.FileSystemService.Remove(
		context.Background(),
		connect.NewRequest(&pb.FileRemoveRequest{
			Context: fs.parseNameToContext(name),
		}),
	)
	if err != nil {
		return err
	}

	return nil
}

func (fs *fileSystem) Rename(ctx context.Context, oldName, newName string) error {
	srcContext := fs.parseNameToContext(oldName)
	destContext := fs.parseNameToContext(newName)

	if srcContext.NodeId != destContext.NodeId {
		return fmt.Errorf("rename across nodes is not supported")
	}

	_, err := fs.core.FileSystemService.Move(
		context.Background(),
		connect.NewRequest(&pb.FileMoveRequest{
			Src:  srcContext,
			Dest: destContext,
		}),
	)

	return err
}

func (fs *fileSystem) Stat(ctx context.Context, name string) (os.FileInfo, error) {
	fileInfo, err := fs.stat(name)
	if err != nil {
		return nil, err
	}

	return newFileInfo(fileInfo), nil
}

func (fs *fileSystem) stat(name string) (*pb.File, error) {
	response, err := fs.core.FileSystemService.Stat(
		context.Background(),
		connect.NewRequest(&pb.FileStatRequest{
			Context: fs.parseNameToContext(name),
			Hash:    true,
		}),
	)
	if err != nil {
		log.Error().Err(err).Str("name", name).Msg("stat failed")

		if strings.Contains(err.Error(), "no such file or directory") {
			return nil, os.ErrNotExist
		}

		return nil, err
	}

	return response.Msg.File, nil
}

func (fs *fileSystem) parseNameToContext(name string) *pb.FileContext {
	var nodeId, location, path string

	if fs.cfg.Webdav.Prefix != "" {
		name = strings.TrimPrefix(name, fs.cfg.Webdav.Prefix)
	}

	parts := strings.Split(strings.Trim(name, "/"), "/")
	if len(parts) == 0 {
		return nil
	}

	nodeId = parts[0]

	if len(parts) > 1 {
		location = parts[1]
	}

	if len(parts) > 2 {
		path = strings.Join(parts[2:], "/")
	}

	return &pb.FileContext{NodeId: nodeId, Location: location, Path: path}
}
