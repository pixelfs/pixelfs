package control

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"time"

	"connectrpc.com/connect"
	pb "github.com/pixelfs/pixelfs/gen/pixelfs/v1"
	"github.com/pixelfs/pixelfs/util"
)

func (p *PixelFS) Ls(ctx *pb.FileContext) error {
	response, err := p.Core.FileSystemService.List(
		context.Background(),
		connect.NewRequest(&pb.FileListRequest{
			Context: ctx,
		}),
	)
	if err != nil {
		return err
	}

	if len(response.Msg.Files) == 0 {
		return nil
	}

	// Sort files by name, case-insensitive
	sort.Slice(response.Msg.Files, func(i, j int) bool {
		return strings.ToLower(response.Msg.Files[i].Name) < strings.ToLower(response.Msg.Files[j].Name)
	})

	var userMaxLen int
	var nameMaxLen int
	for _, file := range response.Msg.Files {
		if len(file.User) > userMaxLen {
			userMaxLen = len(file.User)
		}

		if len(file.Name) > nameMaxLen {
			nameMaxLen = len(file.Name)
		}
	}

	for _, file := range response.Msg.Files {
		fileType := util.Cyan.Bold(true).Render("d")

		isDir := file.Type == pb.FileType_DIR || file.Type == pb.FileType_NODE || file.Type == pb.FileType_LOCATION
		if !isDir {
			fileType = util.Style.Bold(true).Render(".")
		}

		permissions := fmt.Sprintf(
			"%s%s%s",
			formatPermissions(file.Perm, 2),
			formatPermissions(file.Perm, 1),
			formatPermissions(file.Perm, 0),
		)

		size := formatSize(file.Size, isDir)
		user := formatUser(file.User, userMaxLen)

		modTime := formatTime(file.ModifiedAt.AsTime().Local())
		fileName := formatFileName(file, nameMaxLen)

		if file.Type == pb.FileType_NODE {
			anyPb, err := file.Extensions["node"].UnmarshalNew()
			if err != nil {
				return err
			}

			node := anyPb.(*pb.Node)
			var nodeStatus string
			if node.Status == pb.NodeStatus_ONLINE {
				nodeStatus = util.Green.Bold(true).Render("ONLINE ")
			} else {
				nodeStatus = util.Red.Bold(true).Render("OFFLINE")
			}

			fileName += " " + nodeStatus + " " + util.Yellow.Bold(true).Render(strings.ToUpper(node.Name))
		}

		if file.Type == pb.FileType_LOCATION {
			anyPb, err := file.Extensions["location"].UnmarshalNew()
			if err != nil {
				return err
			}

			location := anyPb.(*pb.Location)
			fileName += " " + util.Green.Bold(true).Render(location.Path)
		}

		fmt.Println(fileType + permissions + " " + size + " " + user + " " + modTime + " " + fileName)
	}

	return nil
}

func formatUser(user string, length int) string {
	if len(user) > length {
		user = user[:length]
	}

	return util.Yellow.Bold(true).Render(util.PadLeft(user, length, " "))
}

func formatFileName(file *pb.File, length int) string {
	if len(file.Name) > length {
		file.Name = file.Name[:length]
	}

	file.Name = util.PadRight(file.Name, length, " ")

	switch {
	case file.Type == pb.FileType_DIR || file.Type == pb.FileType_NODE || file.Type == pb.FileType_LOCATION:
		return util.Cyan.Bold(true).Render(file.Name)
	case file.Type == pb.FileType_IMAGE:
		return util.Purple.Render(file.Name)
	case file.Type == pb.FileType_DOCUMENT:
		return util.BlueLight.Render(file.Name)
	default:
		return util.Style.Render(file.Name)
	}
}

func formatPermissions(mode uint32, i uint) string {
	bits := mode >> (i * 3)

	permissions := []struct {
		mask  uint32
		char  string
		color func(string) string
	}{
		{4, "r", func(c string) string { return util.Yellow.Bold(i == 2).Render(c) }},
		{2, "w", func(c string) string { return util.Red.Bold(i == 2).Render(c) }},
		{1, "x", func(c string) string { return util.Green.Bold(i == 2).Render(c) }},
	}

	var str string
	for _, perm := range permissions {
		if bits&perm.mask != 0 {
			str += perm.color(perm.char)
		} else {
			str += util.Grey.Render("-")
		}
	}

	return str
}

func formatSize(size int64, isDir bool) string {
	if size == 0 || isDir {
		return util.Grey.Render(util.PadLeft("-", 4, " "))
	}

	const unitStep = 1000
	units := []string{"", "k", "M", "G", "T", "P", "E"}
	sizeFloat := float64(size)
	unitIndex := 0

	for sizeFloat >= unitStep && unitIndex < len(units)-1 {
		sizeFloat /= unitStep
		unitIndex++
	}

	format := "%.1f"
	if sizeFloat >= 10 {
		format = "%.0f"
	}

	formattedSize := fmt.Sprintf(format, sizeFloat)
	unit := units[unitIndex]

	pad := 3
	if unit == "" {
		pad = 4
	}

	return util.Green.Bold(true).Render(util.PadLeft(formattedSize, pad, " ")) + util.Green.Render(unit)
}

func formatTime(t time.Time) string {
	return util.Blue.Render(util.PadLeft(t.Format("2 Jan 15:04"), 12, " "))
}
