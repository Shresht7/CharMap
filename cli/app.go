package main

import (
	"sort"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/sahilm/fuzzy"
)

// The main application model. Responsible for the entire lifecycle of the TUI application
const (
	listHPadding = 4
	listVPadding = 2
)

type model struct {
	// data
	symbols map[string]Symbol

	// components
	input textinput.Model
	list  list.Model

	// styles
	container lipgloss.Style
	separator lipgloss.Style
}

// Instantiates a new bubbletea application model
func NewModel(symbols map[string]Symbol) *model {
	// Text Input Component
	input := textinput.New()
	input.Placeholder = "  Search..."
	input.Cursor.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#ffffff"))
	input.PlaceholderStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#6e6e6eff"))
	input.Width = 20
	input.Focus()

	// List Component
	delegate := newDelegate()
	l := list.New([]list.Item{}, delegate, 0, 0)
	l.Title = ""
	l.SetFilteringEnabled(false)

	return &model{
		symbols:   symbols,
		input:     input,
		list:      l,
		container: lipgloss.NewStyle().Padding(1, 2),
		separator: lipgloss.NewStyle().Foreground(lipgloss.Color("9")),
	}
}

func (m *model) Init() tea.Cmd {
	// Initial: show all items
	return m.refreshList("")
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		// Layout: input height ~7, list takes the rest
		inputHeight := 7
		m.list.SetSize(msg.Width-listHPadding, msg.Height-inputHeight-listVPadding)
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+q", "esc":
			return m, tea.Quit
		case "enter":
			// TODO: Copy selected item to clipboard
		case "up":
			m.list.CursorUp()
			return m, nil
		case "down":
			m.list.CursorDown()
			return m, nil
		}

		// Update input and refresh list when text changes
		var cmd tea.Cmd
		prev := m.input.Value()
		m.input, cmd = m.input.Update(msg)
		if m.input.Value() != prev {
			return m, tea.Batch(cmd, m.refreshList(m.input.Value()))
		}
		return m, cmd
	}
	return m, nil
}

func (m *model) View() string {
	header := m.input.View()
	divider := m.separator.Render(strings.Repeat("-", max(3, len(stripANSI(header)))))
	return m.container.Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			header,
			divider,
			m.list.View(),
		),
	)
}

// refreshList runs fuzzy search and updates the list items
func (m *model) refreshList(query string) tea.Cmd {
	// Convert map to slice
	all := GetAllSymbols(m.symbols)
	// Sort by symbol for consistent ordering
	sort.Slice(all, func(i, j int) bool { return all[i].Symbol < all[j].Symbol })

	items := []list.Item{}
	if strings.TrimSpace(query) == "" {
		for _, s := range all {
			items = append(items, symbolItem{Symbol: s})
		}
	} else {
		matches := fuzzy.FindFrom(query, SymbolList(all))
		for _, m := range matches {
			items = append(items, symbolItem{Symbol: all[m.Index]})
		}
	}

	m.list.SetItems(items)

	// Keep cursor in bounds
	if len(items) > 0 {
		m.list.Select(0)
	}

	return nil
}

// HELPERS
// -------

// Helper function to strip a string of all ANSI escape codes to hopefully get the raw string
// Note: this is a basic implementation and might not handle all ANSI sequences. Consider an external library.
func stripANSI(s string) string {
	var b strings.Builder
	inANSI := false
	for _, r := range s {
		if r == '\x1b' { // ESC character
			inANSI = true
		} else if inANSI && (r == 'm' || (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z')) {
			// End of a common ANSI escape sequence (e.g., "\x1b[...m")
			inANSI = false
		} else if !inANSI {
			b.WriteRune(r)
		}
	}
	return b.String()
}
