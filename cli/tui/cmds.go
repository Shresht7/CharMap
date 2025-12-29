package tui

import (
	"sort"
	"strings"

	bubblesList "github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/sahilm/fuzzy"

	charmap "github.com/Shresht7/CharMap/cli/charmap"
)

func (m *Model) updateSelectedSymbolCmd() tea.Cmd {
	if item := m.list.SelectedItem(); item != nil {
		if symbolItem, ok := item.(SymbolItem); ok {
			m.selectedSymbol = symbolItem.Symbol
		}
	}
	return nil
}

// refreshList runs fuzzy search and updates the list items
func (m *Model) refreshList(query string) tea.Cmd {
	// Convert map to slice
	all := charmap.GetAllSymbols(m.symbols)
	// Sort by symbol for consistent ordering
	sort.Slice(all, func(i, j int) bool { return all[i].Symbol < all[j].Symbol })

	items := []bubblesList.Item{}
	if strings.TrimSpace(query) == "" {
		for _, s := range all {
			items = append(items, SymbolItem{Symbol: s})
		}
	} else {
		matches := fuzzy.FindFrom(query, charmap.SymbolList(all))
		for _, m := range matches {
			items = append(items, SymbolItem{Symbol: all[m.Index]})
		}
	}

	m.list.SetItems(items)

	// Keep cursor in bounds
	if len(items) > 0 {
		m.list.Select(0)
	}

	return nil
}
