package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/Shresht7/CharMap/cli/charmap"
)

type Preview struct {
	symbol charmap.Symbol
}

func NewPreview() *Preview {
	return &Preview{}
}

func (p *Preview) SetSymbol(symbol charmap.Symbol) {
	p.symbol = symbol
}

func (p *Preview) View() string {
	if p.symbol.Symbol == "" {
		return ""
	}

	// Define styles for preview panel
	previewTitleStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("205"))
	previewValueStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("246"))

	content := lipgloss.JoinVertical(lipgloss.Left,
		previewTitleStyle.Render("Symbol:"), previewValueStyle.Render(p.symbol.Symbol), lipgloss.NewStyle().Height(1).Render(""),
		previewTitleStyle.Render("Description:"), previewValueStyle.Render(p.symbol.Description), lipgloss.NewStyle().Height(1).Render(""),
		previewTitleStyle.Render("Category:"), previewValueStyle.Render(p.symbol.Category), lipgloss.NewStyle().Height(1).Render(""),
		previewTitleStyle.Render("Unicode:"), previewValueStyle.Render(p.symbol.Unicode), lipgloss.NewStyle().Height(1).Render(""),
		previewTitleStyle.Render("Decimal:"), previewValueStyle.Render(p.symbol.Decimal), lipgloss.NewStyle().Height(1).Render(""),
		previewTitleStyle.Render("LaTeX:"), previewValueStyle.Render(p.symbol.Latex), lipgloss.NewStyle().Height(1).Render(""),
		previewTitleStyle.Render("Keywords:"), previewValueStyle.Render(strings.Join(p.symbol.Keywords, "\n")),
	)

	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("240")).
		Padding(1, 2).
		Render(content)
}
