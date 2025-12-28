// ui.go
package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type symbolItem struct{ Symbol }

func (i symbolItem) Title() string       { return i.Symbol.Symbol }
func (i symbolItem) Description() string { return i.Symbol.Description }

func (i symbolItem) FilterValue() string {
	// Allow searching across all fields
	return strings.Join([]string{
		i.Symbol.Symbol,
		strings.Join(i.Symbol.Keywords, " "),
		i.Symbol.Unicode,
		i.Symbol.Decimal,
		i.Symbol.Latex,
		i.Symbol.Category,
		i.Symbol.Description,
	}, " ")
}

type itemDelegate struct {
	tagStyle      lipgloss.Style
	titleStyle    lipgloss.Style
	metaStyle         lipgloss.Style
	selectedStyle     lipgloss.Style
	selectedTextStyle lipgloss.Style
	unselectedTextStyle lipgloss.Style
}

func newDelegate() itemDelegate {
	return itemDelegate{
		titleStyle:    lipgloss.NewStyle().Bold(true),
		tagStyle:      lipgloss.NewStyle().Foreground(lipgloss.Color("205")).PaddingRight(1),
		metaStyle:     lipgloss.NewStyle().Foreground(lipgloss.Color("246")),
		selectedStyle:     lipgloss.NewStyle().Bold(true),
		selectedTextStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Bold(true),
		unselectedTextStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("252")),
	}
}

func (d itemDelegate) Height() int  { return 2 }
func (d itemDelegate) Spacing() int { return 1 }

func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd {
	return nil
}

func (d itemDelegate) Render(w io.Writer, m list.Model, index int, it list.Item) {
	i := it.(symbolItem)

	// First line: symbol glyph + description
	line1 := fmt.Sprintf("%s  %s",
		d.titleStyle.Render(i.Symbol.Symbol),
		i.Description())

	// Second line: chips (unicode, decimal, category, latex)
	tags := []string{}
	if i.Category != "" {
		tags = append(tags, fmt.Sprintf("[Category: %s]", i.Category))
	}
	if i.Unicode != "" {
		tags = append(tags, fmt.Sprintf("[Unicode: %s]", i.Unicode))
	}
	if i.Decimal != "" {
		tags = append(tags, fmt.Sprintf("[Decimal: %s]", i.Decimal))
	}
	if i.Latex != "" {
		tags = append(tags, fmt.Sprintf("[LaTeX: \\%s]", i.Latex))
	}

	line2 := d.metaStyle.Render(strings.Join(tags, " "))

	var itemTextStyle lipgloss.Style
	if index == m.Index() {
		itemTextStyle = d.selectedTextStyle
	} else {
		itemTextStyle = d.unselectedTextStyle
	}

	// Create the content for the item
	itemContent := lipgloss.JoinVertical(lipgloss.Left,
		itemTextStyle.Render(line1),
		itemTextStyle.Render(line2),
	)

	// Prepend the chevron if selected, or spaces for alignment if not
	if index == m.Index() {
		fmt.Fprint(w, d.selectedStyle.Render("❯ ")+itemContent)
	} else {
		fmt.Fprint(w, "  "+itemContent)
	}

}
