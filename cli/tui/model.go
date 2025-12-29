package tui

import (
	bubblesList "github.com/charmbracelet/bubbles/list"
	bubblesTextInput "github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	charmap "github.com/Shresht7/CharMap/cli/charmap"
)

// The main application model. Responsible for the entire lifecycle of the TUI application
type Model struct {
	// data
	symbols map[string]charmap.Symbol

	// components
	input bubblesTextInput.Model
	list  bubblesList.Model

	// State
	selectedSymbol charmap.Symbol

	// layout
	listWidth    int
	previewWidth int

	// styles
	container lipgloss.Style
	separator lipgloss.Style
}

// Instantiates a new bubbletea application model
func NewModel(symbols map[string]charmap.Symbol) *Model {
	// Text Input Component
	input := bubblesTextInput.New()
	input.Placeholder = "  Search..."
	input.Cursor.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#ffffff"))
	input.PlaceholderStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#6e6e6eff"))
	input.Width = 20
	input.Focus()

	// List Component
	l := NewList()

	return &Model{
		symbols:   symbols,
		input:     input,
		list:      l,
		container: lipgloss.NewStyle().Padding(1, 2),
		separator: lipgloss.NewStyle().Foreground(lipgloss.Color("9")),
	}
}

func (m *Model) Init() tea.Cmd {
	// Initial: show all items
	return tea.Batch(m.refreshList(""), m.updateSelectedSymbolCmd())
}
