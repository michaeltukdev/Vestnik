package styles

import "github.com/charmbracelet/lipgloss"

var (
	KeywordStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("204")).Background(lipgloss.Color("235"))
	HelpStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
)
