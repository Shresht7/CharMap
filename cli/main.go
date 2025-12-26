package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	initialModel := &model{
		list: []string{"Item 1", "Item 2", "Item 3"},
	}
	p := tea.NewProgram(initialModel)
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
	}
}
