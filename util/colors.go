package util

import "github.com/charmbracelet/lipgloss"

var (
	Style     = lipgloss.NewStyle()
	Bold      = Style.Bold(true)
	Yellow    = Style.Foreground(lipgloss.Color("221"))
	Blue      = Style.Foreground(lipgloss.Color("20"))
	Red       = Style.Foreground(lipgloss.Color("9"))
	Cyan      = Style.Foreground(lipgloss.Color("81"))
	Green     = Style.Foreground(lipgloss.Color("106"))
	Purple    = Style.Foreground(lipgloss.Color("89"))
	BlueLight = Style.Foreground(lipgloss.Color("105"))
	Grey      = Style.Foreground(lipgloss.Color("241"))
)
