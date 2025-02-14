package styles

import "github.com/charmbracelet/lipgloss"

var (
	SelectedStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF")).Background(lipgloss.Color("#5B78D4")).Padding(0, 1)
	UnselectedStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#FAFAFA")).Padding(0, 1)
	EnabledButtonStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#5B78D4"))
	DisabledButtonStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#888888"))
	RSSDescription      = lipgloss.NewStyle().Foreground(lipgloss.Color("#888888")).PaddingTop(0).PaddingBottom(1).PaddingLeft(4)
	SecondaryText       = lipgloss.NewStyle().Foreground(lipgloss.Color("#888888")).Padding(0, 1)
)

func RenderButton(label string, enabled bool) string {
	if enabled {
		return EnabledButtonStyle.Render("[" + label + "]")
	}
	return DisabledButtonStyle.Render(label)
}
