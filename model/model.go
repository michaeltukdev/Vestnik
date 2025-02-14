package model

import (
	tea "github.com/charmbracelet/bubbletea"
)

type screen int
type mode int

const (
	screenFeeds screen = iota
	screenSettings
)

type Model struct {
	CurrentScreen screen
	Choices       []string
	ChoicesCursor int
	Width         int
	Selected      map[int]struct{}
	ItemCursor    int
	CurrentPage   int
	ItemsPerPage  int
	Mode          mode
}

// ^ Required variables

func InitialModel() Model {
	return Model{
		CurrentScreen: screenFeeds,
		Choices: []string{
			"feeds",
			"settings",
		},
		Selected:     make(map[int]struct{}),
		CurrentPage:  0,
		ItemsPerPage: 10,
		Mode:         0,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width

	case tea.KeyMsg:
		var cmd tea.Cmd
		m, cmd = m.handleKeyPress(msg)
		return m, cmd
	}

	return m, nil
}

func (m Model) View() string {
	switch m.CurrentScreen {
	case screenFeeds:
		return m.feedsView()
	case screenSettings:
		return m.settingsView()
	default:
		return "Unknown screen"
	}
}

func (m *Model) handleSelection() {
	switch m.ChoicesCursor {
	case 0:
		m.CurrentScreen = screenFeeds
	case 1:
		m.CurrentScreen = screenSettings
	}
}
