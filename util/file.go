package util

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/maypok86/otter"
	pb "github.com/pixelfs/pixelfs/gen/pixelfs/v1"
)

const (
	PermissionFallback = 0o700
)

var (
	cacheHash otter.Cache[string, string]
)

func init() {
	cacheHash, _ = otter.MustBuilder[string, string](10_000).
		CollectStats().
		Cost(func(key string, value string) uint32 {
			return 1
		}).
		WithTTL(time.Hour).
		Build()
}

func GetFileHash(path string) (string, error) {
	info, err := os.Stat(path)
	if err != nil {
		return "", fmt.Errorf("failed to stat file: %w", err)
	}

	modTime := info.ModTime()
	cacheKey := fmt.Sprintf("%s:%d", path, modTime.UnixNano())

	if hash, ok := cacheHash.Get(cacheKey); ok {
		return hash, nil
	}

	file, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	hasher := md5.New()
	if info.Size() > 4_096_000 {
		if _, err := io.CopyN(hasher, file, 2_097_152); err != nil {
			return "", fmt.Errorf("failed to read first chunk: %w", err)
		}
		if _, err := file.Seek(-2_097_152, io.SeekEnd); err != nil {
			return "", fmt.Errorf("failed to seek to last chunk: %w", err)
		}
		if _, err := io.Copy(hasher, file); err != nil {
			return "", fmt.Errorf("failed to read last chunk: %w", err)
		}

		hasher.Write([]byte(fmt.Sprintf("%d", info.Size())))
	} else {
		if _, err := io.Copy(hasher, file); err != nil {
			return "", fmt.Errorf("failed to read file for hashing: %w", err)
		}
	}

	hash := fmt.Sprintf("%x", hasher.Sum(nil))
	cacheHash.Set(cacheKey, hash)

	return hash, nil
}

func EnsureDir(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, PermissionFallback)
		if err != nil {
			if errors.Is(err, os.ErrPermission) {
				return fmt.Errorf(
					"creating directory %s, failed with permission error, is it located somewhere can write?",
					dir,
				)
			}

			return fmt.Errorf("creating directory %s: %w", dir, err)
		}
	}

	return nil
}

func GetHomeDir() (string, error) {
	if os.Getenv("PIXELFS_HOME") != "" {
		return os.Getenv("PIXELFS_HOME"), nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	pixelfsHome := fmt.Sprintf("%s/.pixelfs", home)
	if err := EnsureDir(pixelfsHome); err != nil {
		return "", err
	}

	return pixelfsHome, nil
}

func GetFileType(fileInfo os.FileInfo) pb.FileType {
	switch {
	case fileInfo.IsDir():
		return pb.FileType_DIR
	case IsImage(fileInfo.Name()):
		return pb.FileType_IMAGE
	case IsDocument(fileInfo.Name()):
		return pb.FileType_DOCUMENT
	default:
		return pb.FileType_UNKNOWN
	}
}

func SplitPath(platform string, path string) (dir, file string) {
	if platform == "windows" {
		path = strings.ReplaceAll(path, "/", "\\")
		lastSep := strings.LastIndex(path, "\\")
		if lastSep == -1 {
			return "", path
		}
		return path[:lastSep+1], path[lastSep+1:]
	}

	return filepath.Split(path)
}

func JoinPath(platform string, elems ...string) string {
	if platform == "windows" {
		return strings.ReplaceAll(filepath.Join(elems...), "/", "\\")
	}

	return strings.ReplaceAll(filepath.Join(elems...), "\\", "/")
}
