package components

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	charmap "github.com/Shresht7/CharMap/cli/charmap"
)

func NewList() list.Model {
	delegate := newDelegate()
	l := list.New([]list.Item{}, delegate, 0, 0)
	l.Title = ""
	l.SetFilteringEnabled(false)
	l.SetShowStatusBar(false)
	l.SetShowHelp(true)
	return l
}


type SymbolItem struct{ charmap.Symbol }

func (i SymbolItem) Title() string       { return i.Symbol.Symbol }
func (i SymbolItem) Description() string { return i.Symbol.Description }

func (i SymbolItem) FilterValue() string {
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
	tagStyle            lipgloss.Style
	titleStyle          lipgloss.Style
	metaStyle           lipgloss.Style
	selectedStyle       lipgloss.Style
	selectedTextStyle   lipgloss.Style
	unselectedTextStyle lipgloss.Style
}

func newDelegate() itemDelegate {
	return itemDelegate{
		titleStyle:          lipgloss.NewStyle().Bold(true),
		tagStyle:            lipgloss.NewStyle().Foreground(lipgloss.Color("205")).PaddingRight(1),
		metaStyle:           lipgloss.NewStyle().Foreground(lipgloss.Color("246")),
		selectedStyle:       lipgloss.NewStyle().Bold(true),
		selectedTextStyle:   lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Bold(true),
		unselectedTextStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("240")),
	}
}

func (d itemDelegate) Height() int  { return 1 }
func (d itemDelegate) Spacing() int { return 1 }

func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd {
	return nil
}

func (d itemDelegate) Render(w io.Writer, m list.Model, index int, it list.Item) {
	i := it.(SymbolItem)

	// First line: symbol glyph + description
	line1 := fmt.Sprintf("%s  %s",
		d.titleStyle.Render(i.Symbol.Symbol),
		i.Symbol.Description)

	var itemTextStyle lipgloss.Style
	if index == m.Index() {
		itemTextStyle = d.selectedTextStyle
	} else {
		itemTextStyle = d.unselectedTextStyle
	}

	// Create the content for the item
	itemContent := itemTextStyle.Render(line1)

	// Prepend the chevron if selected, or spaces for alignment if not
	if index == m.Index() {
		fmt.Fprint(w, d.selectedStyle.Render("❯ ")+itemContent)
	} else {
		// Apply the unselectedTextStyle to the entire line including padding
		fmt.Fprint(w, d.unselectedTextStyle.Render("  "+itemContent))
	}

}
