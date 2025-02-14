package model

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) handleKeyPress(msg tea.KeyMsg) (Model, tea.Cmd) {
	switch msg.String() {
	case "right", "left":
		m.handleNavigationKeys(msg.String())
	case "up", "down":
		m.handleItemNavigation(msg.String())
	case "n", "p":
		m.handlePagination(msg.String())
	case "enter":
		m.handleEnterKey()
	case "tab":
		m.toggleMode()
	case "ctrl+c", "q":
		return m, tea.Quit
	}
	return m, nil
}

func (m *Model) handleNavigationKeys(key string) {
	switch key {
	case "right":
		if m.ChoicesCursor < len(m.Choices)-1 && m.Mode == 0 {
			m.ChoicesCursor++
		}
	case "left":
		if m.ChoicesCursor > 0 && m.Mode == 0 {
			m.ChoicesCursor--
		}
	}
}

func (m *Model) handleItemNavigation(key string) {
	switch key {
	case "up":
		if m.ItemCursor > 0 && m.Mode == 1 {
			m.ItemCursor--
		}
	case "down":
		if m.Mode == 1 {
			m.ItemCursor++
		}
	}
}

func (m *Model) handlePagination(key string) {
	switch key {
	case "n":
		m.CurrentPage++
	case "p":
		if m.CurrentPage > 0 {
			m.CurrentPage--
		}
	}
}

func (m *Model) handleEnterKey() {
	switch m.Mode {
	case 0: // Menu navigation mode
		m.handleSelection()
	case 1: // Feed selection mode
		// Open the selected RSS feed
		// m.openRSSFeed()
	}
}

func (m *Model) toggleMode() {
	if m.Mode == 0 {
		m.Mode = 1
	} else {
		m.Mode = 0
	}
}
