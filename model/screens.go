package model

import (
	"fmt"
	"log"

	"github.com/charmbracelet/lipgloss"
	"github.com/michaeltukdev/Vestnik/rss"
	"github.com/michaeltukdev/Vestnik/styles"
)

// View for 'Feeds' tab

func (m Model) feedsView() string {
	s := ""

	var allItems []rss.Item
	for _, feed := range rss.GetFeeds() {
		data, err := rss.FetchRSSFeed(feed.URL)
		if err != nil {
			log.Fatal("Error fetching RSS feed:", err)
		}
		allItems = append(allItems, data.Channel.Items...)
	}

	totalItems := len(allItems)
	startIndex := m.CurrentPage * m.ItemsPerPage
	endIndex := startIndex + m.ItemsPerPage

	if startIndex >= totalItems {
		startIndex = totalItems
	}
	if endIndex > totalItems {
		endIndex = totalItems
	}

	for i, item := range allItems[startIndex:endIndex] {
		cursor := " "
		if i == m.ItemCursor {
			cursor = ">"
		}

		description := item.Description
		if len(description) > 80 {
			description = description[:80]
		}

		s += fmt.Sprintf("%s %s\n %s \n", cursor, item.Title, styles.RSSDescription.Render(`'`+description+`'`))
	}

	paginationControls := m.paginationControls(totalItems)

	return fmt.Sprintf("%s\n\n%s\n\n%s\n", m.navigation(), s, paginationControls)
}

// View for 'Settings' tab

func (m Model) settingsView() string {
	return fmt.Sprintf("%s\n\n%s\n", m.navigation(), "Settings view")
}

// Navigation Component

func (m Model) navigation() string {
	s := ""
	for i, choice := range m.Choices {
		if i == m.ChoicesCursor {
			s += styles.SelectedStyle.Render(choice)
		} else {
			s += styles.UnselectedStyle.Render(choice)
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

	modeLabel := styles.SecondaryText.Render(fmt.Sprintf("Current Mode: %s", currentMode))
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

// Pagination Component

func (m Model) paginationControls(totalItems int) string {
	totalPages := (totalItems + m.ItemsPerPage - 1) / m.ItemsPerPage

	prevButton := styles.RenderButton("Previous", m.CurrentPage > 0)
	nextButton := styles.RenderButton("Next", (m.CurrentPage+1)*m.ItemsPerPage < totalItems)

	pageInfo := fmt.Sprintf("Page %d of %d", m.CurrentPage+1, totalPages)

	return fmt.Sprintf("%s • %s • %s", prevButton, pageInfo, nextButton)
}
