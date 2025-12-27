package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/sahilm/fuzzy"
)

type model struct {
	symbols map[string]Symbol
	query   string
	results []Symbol
}

func (m *model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		default:
			return m, nil
		}
	}
	return m, nil
}

func (m *model) View() string {
	var s strings.Builder

	list := GetAllSymbols(m.symbols)
	if len(m.query) > 0 {
		list = FuzzySearch(list, m.query)
	}

	for _, symbol := range list {
		s.WriteString(fmt.Sprintf("%s\n", symbol.Symbol))
	}

	return s.String()
}

func FuzzySearch(symbols []Symbol, query string) []Symbol {
	matches := fuzzy.FindFrom(query, SymbolList(symbols))
	results := []Symbol{}
	for _, m := range matches {
		results = append(results, symbols[m.Index])
	}
	return results
}
