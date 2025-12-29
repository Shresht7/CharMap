package main

import (
	_ "embed"
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/Shresht7/CharMap/cli/charmap"
	"github.com/Shresht7/CharMap/cli/tui"
)

//go:embed charmap.json
var charmapJSON []byte

// The main entrypoint to the application
func main() {
	// Load the character-map data
	symbols, err := charmap.LoadSymbols(charmapJSON)
	if err != nil {
		fmt.Println("Error loading symbols:", err)
		return
	}

	// Initialize the bubbletea application model
	initialModel := tui.NewModel(symbols)

	// Run the application
	p := tea.NewProgram(initialModel)
	if _, err := p.Run(); err != nil {
		log.Fatalf("Failed to start program: %s\n", err)
	}
}
