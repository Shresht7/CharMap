package tui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/Shresht7/CharMap/cli/charmap"
	"github.com/Shresht7/CharMap/cli/tui/components"
)

// The main application model. Responsible for the entire lifecycle of the TUI application
type Model struct {
	// data
	symbols map[string]charmap.Symbol

	// components
	input   textinput.Model
	list    list.Model
	preview *components.Preview

	// layout
	width    int
	height   int
	hPadding int
	vPadding int

	// styles
	container lipgloss.Style
	separator lipgloss.Style
}

// Instantiates a new bubbletea application model
func NewModel(symbols map[string]charmap.Symbol) *Model {
	return &Model{
		symbols:   symbols,
		input:     components.NewInput(),
		list:      components.NewList(),
		preview:   components.NewPreview(),
		hPadding:  4,
		vPadding:  2,
		container: lipgloss.NewStyle(),
		separator: lipgloss.NewStyle().Foreground(lipgloss.Color("9")),
	}
}


func (m *Model) Init() tea.Cmd {
	// Initial: show all items
	return tea.Batch(m.refreshList(""), m.updateSelectedSymbolCmd())
}
