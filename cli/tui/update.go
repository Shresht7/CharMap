package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		inputHeight := 7
		dividerHeight := 1 // Height of the divider line
		// Calculate split widths
		totalWidth := msg.Width - HPadding // Account for outer container padding
		m.listWidth = int(float64(totalWidth) * 0.70)
		m.previewWidth = totalWidth - m.listWidth

		// Set input width to take up full available width
		m.input.Width = totalWidth - m.container.GetHorizontalPadding() - m.container.GetHorizontalBorderSize() // Adjust for container's padding

		m.list.SetSize(m.listWidth, msg.Height-inputHeight-dividerHeight-VPadding)
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
