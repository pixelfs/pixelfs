package util

import (
	"fmt"

	"github.com/k0kubun/go-ansi"
	"github.com/schollz/progressbar/v3"
)

func NewProgressBar(total int, description string) *progressbar.ProgressBar {
	runeDescription := []rune(description)
	if len(runeDescription) > 40 {
		description = string(runeDescription[:37]) + "..."
	} else {
		description = PadRight(description, 40, " ")
	}

	return progressbar.NewOptions(
		total,
		progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowCount(),
		progressbar.OptionShowIts(),
		progressbar.OptionSetWidth(35),
		progressbar.OptionOnCompletion(func() {
			fmt.Println()
		}),
		progressbar.OptionSetDescription(description),
		progressbar.OptionSetTheme(
			progressbar.Theme{
				Saucer:        "[green]=[reset]",
				SaucerHead:    "[green]>[reset]",
				SaucerPadding: " ",
				BarStart:      "[",
				BarEnd:        "]",
			},
		),
	)
}
