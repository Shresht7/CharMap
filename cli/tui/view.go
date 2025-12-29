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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
