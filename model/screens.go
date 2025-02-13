package model

import (
	"fmt"
	"log"

	"github.com/charmbracelet/lipgloss"
	"github.com/michaeltukdev/Vestnik/rss"
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
	s := ""

	// Fetch all feeds
	var allItems []rss.Item
	for _, feed := range rss.GetFeeds() {
		data, err := rss.FetchRSSFeed(feed.URL)
		if err != nil {
			log.Fatal("Error fetching RSS feed:", err)
		}
		allItems = append(allItems, data.Channel.Items...)
	}

	// Pagination logic
	totalItems := len(allItems)
	startIndex := m.CurrentPage * m.ItemsPerPage
	endIndex := startIndex + m.ItemsPerPage

	if startIndex >= totalItems {
		startIndex = totalItems
	}
	if endIndex > totalItems {
		endIndex = totalItems
	}

	var descriptionStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#888888")).
		PaddingTop(0).
		PaddingBottom(1).
		PaddingLeft(4)

	// Display only the items for the current page
	for i, item := range allItems[startIndex:endIndex] {
		cursor := " "          // Default cursor (no selection)
		if i == m.ItemCursor { // Highlight the selected item
			cursor = ">"
		}

		description := item.Description
		if len(description) > 80 {
			description = description[:80]
		}

		s += fmt.Sprintf("%s %s\n %s \n", cursor, item.Title, descriptionStyle.Render(`'`+description+`'`))
	}

	// Add pagination controls
	paginationControls := m.paginationControls(totalItems)

	return fmt.Sprintf("%s\n\n%s\n\n%s\n", m.navigation(), s, paginationControls)
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
		if i == m.ChoicesCursor {
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

func (m Model) paginationControls(totalItems int) string {
	totalPages := (totalItems + m.ItemsPerPage - 1) / m.ItemsPerPage // Calculate total pages

	// Create "Previous" and "Next" buttons
	prevButton := "Previous"
	nextButton := "Next"

	if m.CurrentPage > 0 {
		prevButton = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#5B78D4")).
			Render("[Previous]")
	} else {
		prevButton = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#888888")).
			Render("Previous")
	}

	if (m.CurrentPage+1)*m.ItemsPerPage < totalItems {
		nextButton = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#5B78D4")).
			Render("[Next]")
	} else {
		nextButton = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#888888")).
			Render("Next")
	}

	// Display current page and total pages
	pageInfo := fmt.Sprintf("Page %d of %d", m.CurrentPage+1, totalPages)

	return fmt.Sprintf("%s • %s • %s", prevButton, pageInfo, nextButton)
}
