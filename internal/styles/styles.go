package styles

import (
	_ "embed"

	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

var (
	tagbarBorder = lipgloss.Border{
		Right: "│",
	}
	tagbarHighlight = lipgloss.AdaptiveColor{Light: "#555", Dark: "#555"}
	ActiveTagbar    = lipgloss.NewStyle().Foreground(lipgloss.Color("#00BFFF"))
	Tagbar          = lipgloss.NewStyle().BorderForeground(tagbarHighlight).Border(tagbarBorder, true).Margin(1, 0, 0, 1)

	Page   = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF8247")).Align(lipgloss.Right).MarginRight(1)
	Status = lipgloss.NewStyle().Faint(true).Foreground(lipgloss.AdaptiveColor{
		Light: "#909090",
		Dark:  "#626262",
	}).Align(lipgloss.Center)
	Search = lipgloss.NewStyle().Faint(true).Align(lipgloss.Left).MarginLeft(2)
)

//go:embed theme.json
var DefaultTheme []byte

// SelectTheme sets theme
func SelectTheme(theme string) glamour.TermRendererOption {
	switch theme {
	case "light":
		return glamour.WithStandardStyle("light")
	case "dark":
		return glamour.WithStandardStyle("dark")
	default:
		return getDefaultTheme()
	}
}

func getDefaultTheme() glamour.TermRendererOption {
	return glamour.WithStylesFromJSONBytes(DefaultTheme)
}
