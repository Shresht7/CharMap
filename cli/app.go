package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	list []string
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
	for _, r := range m.list {
		s.WriteString(fmt.Sprintf("%s\n", r))
	}
	return s.String()
}
