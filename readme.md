# Vestnik

Vestnik is a terminal-based RSS feed reader built with [Bubble Tea](https://github.com/charmbracelet/bubbletea) and [Lip Gloss](https://github.com/charmbracelet/lipgloss). View and navigate RSS feed items, paginate through results, and open articles in your default web browser—all from the command line.

This was primarily built for myself but others can use, and improve it if they wish. One of my goals for 2025 was to read more hence why this project was made, it is part of my daily reading schedule.

## Features

- **Menu Navigation:** Switch between different screens (e.g. Feeds, Settings) using the arrow keys.
- **Feed Selection:** Browse RSS feed items, view descriptions and categories.
- **Pagination:** Navigate through feed items with pagination controls.
- **Open in Browser:** Press `Enter` on a selected RSS feed item to open its link in your default browser.
- **Keyboard Shortcuts:** Navigate using arrow keys, paginate using `n`/`p`, toggle selection mode with `tab`, and exit with `q` or `ctrl+c`.

### Keyboard Controls

In a future update, these will all be customisable but still the defaults.

- **Menu Mode (Navigating)**
  - **Left/Right Arrows:** Switch between menu options (`feeds`, `settings`).
  - **Enter:** Select the highlighted menu option.
  - **Tab:** Toggle between menu mode and feed selection mode.
  - **q / ctrl+c:** Quit the application.

- **Feed Selection Mode**
  - **Up/Down Arrows:** Navigate through RSS feed items.
  - **n / p:** Go to the next or previous page.
  - **Enter:** Open the selected RSS item in your browser.

## Todo List

All of the items in the list can be added by anywhere, feel free to contribute;
 - Better filtering and searching for the RSS feed (filtering would be through the category system already implemented)
 - A way to modify keybinds within the settings tab (probably written to the json file)
 - A way to manage feeds through the settings tab

## Project Structure

- **`model.go`:** Contains the main application logic, state management, and view rendering.
- **`ui/`:**
  - **`ui.go`:** Contains styling definitions using Lip Gloss and helper functions for button rendering.
- **`feeds/`:** Contains logic to fetch and combine RSS feeds.
- **`README.md`:** Project documentation and usage instructions.

## License

This project is licensed under the MIT License.

## Acknowledgements

- [Bubble Tea](https://github.com/charmbracelet/bubbletea)
- [Lip Gloss](https://github.com/charmbracelet/lipgloss)
- [pkg/browser](https://github.com/pkg/browser)
- Thanks to the open-source community for inspiration and contributions.
