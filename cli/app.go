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
	selectedSymbol Symbol
	listWidth      int
	previewWidth   int

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
	l.SetShowStatusBar(false)
	l.SetShowHelp(true)

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
	return tea.Batch(m.refreshList(""), m.updateSelectedSymbolCmd())
}

func (m *model) updateSelectedSymbolCmd() tea.Cmd {
	if item := m.list.SelectedItem(); item != nil {
		if symbolItem, ok := item.(symbolItem); ok {
			m.selectedSymbol = symbolItem.Symbol
		}
	}
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		inputHeight := 7
		dividerHeight := 1 // Height of the divider line
		// Calculate split widths
		totalWidth := msg.Width - listHPadding // Account for outer container padding
		m.listWidth = int(float64(totalWidth) * 0.70)
		m.previewWidth = totalWidth - m.listWidth

		// Set input width to take up full available width
		m.input.Width = totalWidth - m.container.GetHorizontalPadding() - m.container.GetHorizontalBorderSize() // Adjust for container's padding

		m.list.SetSize(m.listWidth, msg.Height-inputHeight-dividerHeight-listVPadding)
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+q", "esc":
			return m, tea.Quit
		case "enter":
			// TODO: Copy selected item to clipboard
			return m, m.updateSelectedSymbolCmd()
		case "up":
			m.list.CursorUp()
			return m, m.updateSelectedSymbolCmd()
		case "down":
			m.list.CursorDown()
			return m, m.updateSelectedSymbolCmd()
		}

		// Update input and refresh list when text changes
		var cmd tea.Cmd
		prev := m.input.Value()
		m.input, cmd = m.input.Update(msg)
		if m.input.Value() != prev {
			return m, tea.Batch(cmd, m.refreshList(m.input.Value()), m.updateSelectedSymbolCmd())
		}
		return m, cmd
	}
	return m, nil
}

func (m *model) View() string {
	// Input field always takes full width
	inputView := m.input.View()

	// Divider below the input, spanning its full width
	divider := m.separator.Render(strings.Repeat("-", max(3, lipgloss.Width(inputView))))

	// Main list content (left panel) - no longer includes the input header
	listPanel := lipgloss.NewStyle().Width(m.listWidth).Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			m.list.View(),
			lipgloss.NewStyle().Height(1).Render(""), // Add padding after list/help text
		),
	)

	// Preview panel (right panel)
	previewPanelContent := ""
	if m.selectedSymbol.Symbol != "" { // Only show if a symbol is selected
		// Define styles for preview panel
		previewTitleStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("205"))
		previewValueStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("246"))

		previewPanelContent = lipgloss.JoinVertical(lipgloss.Left,
			previewTitleStyle.Render("Symbol:"), previewValueStyle.Render(m.selectedSymbol.Symbol), lipgloss.NewStyle().Height(1).Render(""),
			previewTitleStyle.Render("Description:"), previewValueStyle.Render(m.selectedSymbol.Description), lipgloss.NewStyle().Height(1).Render(""),
			previewTitleStyle.Render("Category:"), previewValueStyle.Render(m.selectedSymbol.Category), lipgloss.NewStyle().Height(1).Render(""),
			previewTitleStyle.Render("Unicode:"), previewValueStyle.Render(m.selectedSymbol.Unicode), lipgloss.NewStyle().Height(1).Render(""),
			previewTitleStyle.Render("Decimal:"), previewValueStyle.Render(m.selectedSymbol.Decimal), lipgloss.NewStyle().Height(1).Render(""),
			previewTitleStyle.Render("LaTeX:"), previewValueStyle.Render(m.selectedSymbol.Latex), lipgloss.NewStyle().Height(1).Render(""),
			previewTitleStyle.Render("Keywords:"), previewValueStyle.Render(strings.Join(m.selectedSymbol.Keywords, "\n")),
		)
	}

	previewPanel := lipgloss.NewStyle().
		Width(m.previewWidth). // Use calculated width
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("240")).
		Padding(1, 2).
		Render(previewPanelContent)

	// Join the input, divider, and then the horizontal panels vertically
	return m.container.Render(
		lipgloss.JoinVertical(lipgloss.Left,
			inputView,
			divider,
			lipgloss.JoinHorizontal(lipgloss.Top,
				listPanel,
				previewPanel,
			),
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
