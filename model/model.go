package model

import (
	tea "github.com/charmbracelet/bubbletea"
)

type screen int

const (
	screenFeeds screen = iota
	screenSettings
)

type Model struct {
	CurrentScreen screen
	Choices       []string
	Cursor        int
	width         int
	height        int
	Selected      map[int]struct{}
}

func InitialModel() Model {
	return Model{
		CurrentScreen: screenFeeds,
		Choices: []string{
			"feeds",
			"settings",
		},
		Selected: make(map[int]struct{}),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "left":
			if m.Cursor > 0 {
				m.Cursor--
			}
		case "right":
			if m.Cursor < len(m.Choices)-1 {
				m.Cursor++
			}
		case "enter": // Select the current menu item
			m.handleSelection()
		case "esc": // Go back to the main screen

		case "ctrl+c", "q": // Quit the program
			return m, tea.Quit
		}
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
	switch m.Cursor {
	case 0:
		m.CurrentScreen = screenFeeds
	case 1:
		m.CurrentScreen = screenSettings
	}
}
