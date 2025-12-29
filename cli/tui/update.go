package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.container.Padding(m.vPadding, m.hPadding)
		m.input.Width = m.width - m.hPadding*2
		m.list.SetSize(m.width/2-m.hPadding, m.height-m.vPadding*2-7)
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
