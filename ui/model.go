package ui

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/michaeltukdev/Vestnik/feeds"
	"github.com/pkg/browser"
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
	TotalPages    int
	Mode          mode
	FeedItems     []feeds.Item
}

func InitialModel() *Model {
	return &Model{
		CurrentScreen: screenFeeds,
		Choices:       []string{"feeds", "settings"},
		Selected:      make(map[int]struct{}),
		CurrentPage:   0,
		ItemsPerPage:  10,
		TotalPages:    0,
		Mode:          0,
	}
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
	case tea.KeyMsg:
		return m.handleKeyPress(msg)
	}
	return m, nil
}

func (m *Model) View() string {
	switch m.CurrentScreen {
	case screenFeeds:
		return m.feedsView()
	case screenSettings:
		return m.settingsView()
	default:
		return "Unknown screen"
	}
}

func (m *Model) handleKeyPress(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
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
	if m.Mode != 1 || m.FeedItems == nil {
		return
	}
	visibleCount := m.visibleItemCount()
	switch key {
	case "up":
		if m.ItemCursor > 0 {
			m.ItemCursor--
		}
	case "down":
		if m.ItemCursor < visibleCount-1 {
			m.ItemCursor++
		}
	}
}

func (m *Model) handlePagination(key string) {
	totalItems := len(m.FeedItems)
	switch key {
	case "n":
		if (m.CurrentPage+1)*m.ItemsPerPage < totalItems {
			m.CurrentPage++
			m.adjustCursorPosition()
		}
	case "p":
		if m.CurrentPage > 0 {
			m.CurrentPage--
			m.adjustCursorPosition()
		}
	}
}

func (m *Model) handleEnterKey() {
	if m.Mode == 0 {
		m.handleSelection()
	} else if m.Mode == 1 {
		m.openRSSFeed()
	}
}

func (m *Model) handleSelection() {
	switch m.ChoicesCursor {
	case 0:
		m.CurrentScreen = screenFeeds
		m.ItemCursor = 0
		m.CurrentPage = 0
		m.FeedItems = nil
	case 1:
		m.CurrentScreen = screenSettings
	}
}

func (m *Model) openRSSFeed() {
	if m.FeedItems == nil {
		return
	}
	idx := m.CurrentPage*m.ItemsPerPage + m.ItemCursor
	if idx < 0 || idx >= len(m.FeedItems) {
		return
	}
	url := m.FeedItems[idx].Link
	if url == "" {
		return
	}
	if err := browser.OpenURL(url); err != nil {
		log.Println("Error opening URL:", err)
	}
}

func (m *Model) toggleMode() {
	if m.Mode == 0 {
		m.Mode = 1
	} else {
		m.Mode = 0
	}
}

func (m *Model) calculateTotalPages(totalItems int) {
	m.TotalPages = (totalItems + m.ItemsPerPage - 1) / m.ItemsPerPage
}

func (m *Model) feedsView() string {
	if m.FeedItems == nil {
		items, err := feeds.FetchAndCombineFeeds()
		if err != nil {
			log.Fatal("Error fetching and combining RSS feeds:", err)
		}
		m.FeedItems = items
	}
	totalItems := len(m.FeedItems)
	m.calculateTotalPages(totalItems)
	startIndex := m.CurrentPage * m.ItemsPerPage
	endIndex := startIndex + m.ItemsPerPage
	if startIndex > totalItems {
		startIndex = totalItems
	}
	if endIndex > totalItems {
		endIndex = totalItems
	}
	if m.ItemCursor > (endIndex - startIndex - 1) {
		m.ItemCursor = endIndex - startIndex - 1
	}
	s := ""
	categoryStyle := lipgloss.NewStyle().
		Background(lipgloss.Color("#FFA500")).
		Foreground(lipgloss.Color("#000000")).
		Padding(0, 1).
		Bold(true)
	for i, item := range m.FeedItems[startIndex:endIndex] {
		cursor := " "
		if i == m.ItemCursor {
			cursor = ">"
		}
		description := item.Description
		if len(description) > 80 {
			description = description[:80]
		}
		categoryLabel := categoryStyle.Render(item.Category)
		s += fmt.Sprintf("%s %s %s [%s]\n %s \n", cursor, categoryLabel,
			item.Title, item.Source, RSSDescription.Render("'"+description+"'"))
	}
	paginationControls := m.paginationControls(totalItems)
	return fmt.Sprintf("%s\n\n%s\n\n%s\n", m.navigation(), s, paginationControls)
}

func (m *Model) settingsView() string {
	return fmt.Sprintf("%s\n\n%s\n", m.navigation(), "Settings view")
}

func (m *Model) visibleItemCount() int {
	totalItems := len(m.FeedItems)
	startIndex := m.CurrentPage * m.ItemsPerPage
	endIndex := startIndex + m.ItemsPerPage
	if startIndex > totalItems {
		return 0
	}
	if endIndex > totalItems {
		endIndex = totalItems
	}
	return endIndex - startIndex
}

func (m *Model) adjustCursorPosition() {
	visibleCount := m.visibleItemCount()
	if visibleCount == 0 {
		m.ItemCursor = 0
	} else if m.ItemCursor > visibleCount-1 {
		m.ItemCursor = visibleCount - 1
	}
}
