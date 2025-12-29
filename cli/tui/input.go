package tui

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

func NewInput() textinput.Model {
	input := textinput.New()
	input.Placeholder = "  Search..."
	input.Cursor.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#ffffff"))
	input.PlaceholderStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#6e6e6eff"))
	input.Width = 20
	input.Focus()
	return input
}
