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
	ChoicesCursor int
	width         int
	height        int
	Selected      map[int]struct{}
	ItemCursor    int
	CurrentPage   int
	ItemsPerPage  int
}

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
		case "right":
			if m.ChoicesCursor < len(m.Choices)-1 {
				m.ChoicesCursor++
			}
		case "left":
			if m.ChoicesCursor > 0 {
				m.ChoicesCursor--
			}
		case "up":
			if m.ItemCursor > 0 {
				m.ItemCursor--
			}
		case "down":
			m.ItemCursor++
		case "n":
			m.CurrentPage++
		case "p":
			if m.CurrentPage > 0 {
				m.CurrentPage--
			}
		case "ctrl+c", "q":
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
	switch m.ChoicesCursor {
	case 0:
		m.CurrentScreen = screenFeeds
	case 1:
		m.CurrentScreen = screenSettings
	}
}
