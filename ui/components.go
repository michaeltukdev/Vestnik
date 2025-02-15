package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func (m *Model) navigation() string {
	s := ""
	for i, choice := range m.Choices {
		if i == m.ChoicesCursor {
			s += SelectedStyle.Render(choice)
		} else {
			s += UnselectedStyle.Render(choice)
		}
		if i < len(m.Choices)-1 {
			s += " • "
		}
	}
	currentMode := ""
	switch m.Mode {
	case 0:
		currentMode = "Navigating"
	case 1:
		currentMode = "Feed Selection"
	default:
		currentMode = "Unknown"
	}
	modeLabel := SecondaryText.Render(
		fmt.Sprintf("Current Mode: %s %d", currentMode, m.TotalPages))
	spacePadding := m.Width - lipgloss.Width(s) - lipgloss.Width(modeLabel) - 2
	if spacePadding > 0 {
		s += lipgloss.NewStyle().Width(spacePadding).Render(" ")
	}
	s += modeLabel
	borderStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, false, true, false).
		BorderForeground(lipgloss.Color("#5B78D4")).
		MarginTop(1).
		Width(m.Width - 1)
	return borderStyle.Render(s)
}

func (m *Model) paginationControls(totalItems int) string {
	prevButton := RenderButton("Previous", m.CurrentPage > 0)
	nextButton := RenderButton("Next", (m.CurrentPage+1)*m.ItemsPerPage < totalItems)
	pageInfo := fmt.Sprintf("Page %d of %d", m.CurrentPage+1, m.TotalPages)
	return fmt.Sprintf("%s • %s • %s", prevButton, pageInfo, nextButton)
}
