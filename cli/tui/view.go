package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m *Model) View() string {
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
	previewPanel := lipgloss.NewStyle().
		Width(m.previewWidth).
		Render(m.preview.View())

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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
