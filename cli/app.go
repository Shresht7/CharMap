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
	list := ToStringList(GetAllSymbols(m.symbols))

	if len(m.results) > 0 {
		list = ToStringList(m.results)
	}

	for _, item := range list {
		s.WriteString(fmt.Sprintf("%s\n", item))
	}

	return s.String()
}

func ToStringList(symbols []Symbol) []string {
	list := make([]string, len(symbols))
	for i, s := range symbols {
		list[i] = s.String()
	}
	return list
}

func FuzzySearch(symbols []Symbol, query string) []Symbol {
	matches := fuzzy.Find(query, ToStringList(symbols))
	results := []Symbol{}
	for _, m := range matches {
		results = append(results, symbols[m.Index])
	}
	return results
}
