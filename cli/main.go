package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	symbols, err := LoadSymbols("../data/charmap.json")
	if err != nil {
		fmt.Println("Error loading symbols:", err)
		return
	}

	initialModel := &model{
		symbols: symbols,
		query:   "",
		results: []Symbol{},
	}

	p := tea.NewProgram(initialModel)
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
	}
}
