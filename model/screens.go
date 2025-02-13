package model

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) mainView() string {
	switch m.CurrentScreen {
	case screenFeeds:
		return m.feedsView()
	case screenSettings:
		return m.settingsView()
	default:
		return "Unknown screen"
	}
}

func (m Model) feedsView() string {
	return fmt.Sprintf("%s\n\n%s\n", m.navigation(), "Feeds view")
}

func (m Model) settingsView() string {
	return fmt.Sprintf("%s\n\n%s\n", m.navigation(), "Settings view")
}

func (m Model) navigation() string {
	selectedStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFFFF")).
		Background(lipgloss.Color("#5B78D4")).
		Padding(0, 1)

	unselectedStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FAFAFA")).
		Padding(0, 1)

	s := ""
	for i, choice := range m.Choices {
		if i == m.Cursor {
			s += selectedStyle.Render(choice)
		} else {
			s += unselectedStyle.Render(choice)
		}

		if i < len(m.Choices)-1 {
			s += " • "
		}
	}

	borderStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, false, true, false).
		BorderForeground(lipgloss.Color("#5B78D4")).
		MarginTop(1).
		Width(m.width - 1)

	return borderStyle.Render(s)
}
