package cli

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/pixelfs/pixelfs/config"
	pb "github.com/pixelfs/pixelfs/gen/pixelfs/v1"
	"github.com/pixelfs/pixelfs/util"
)

func parseFileContext(input string) (*pb.FileContext, error) {
	cfg, err := config.GetConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to get config: %w", err)
	}

	var nodeId, path string
	if cfg.Pwd != "" {
		nodeId, path = extractNodeIdAndPath(cfg.Pwd)
	}

	if input != "" {
		argNodeId, argPath := extractNodeIdAndPath(input)
		if argNodeId != "" {
			nodeId = argNodeId
			path = argPath
		} else {
			path = joinPaths(path, input)
		}
	}

	absolutePath := filepath.Clean(path)
	if absolutePath == "." {
		absolutePath = ""
	} else if absolutePath == ".." {
		absolutePath = filepath.Dir(path)
	}

	if path != "" && nodeId == "" {
		return nil, fmt.Errorf("node-id is required when path is provided")
	}

	location, absolutePath := splitLocationAndPath(absolutePath)
	return &pb.FileContext{NodeId: nodeId, Location: location, Path: absolutePath}, nil
}

func splitLocationAndPath(path string) (string, string) {
	if path == "" {
		return "", ""
	}

	path = filepath.FromSlash(path)
	parts := strings.Split(path, string(filepath.Separator))
	if len(parts) < 2 {
		return "", ""
	}

	return parts[1], strings.Join(parts[2:], string(filepath.Separator))
}

func extractNodeIdAndPath(input string) (string, string) {
	if util.IsNodeId(input) {
		return input, ""
	}
	parts := strings.SplitN(input, ":", 2)
	if len(parts) == 2 && util.IsNodeId(parts[0]) {
		return parts[0], parts[1]
	}
	return "", ""
}

func joinPaths(base, additional string) string {
	if base == "" || strings.HasPrefix(additional, "/") {
		return additional
	}

	if strings.HasSuffix(base, "/") {
		base = base[:len(base)-1]
	}

	return base + "/" + additional
}
