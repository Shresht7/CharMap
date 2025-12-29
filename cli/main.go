package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// The main entrypoint to the application
func main() {
	// Load the character-map data
	// TODO: Hardcoded data-path. Embed or make it configurable
	// 		 This makes the application less flexible if the data file needs to be
	//		 located elsewhere, or i the executable is run from a different directory than intended.
	symbols, err := LoadSymbols("../data/charmap.json")
	if err != nil {
		fmt.Println("Error loading symbols:", err)
		return
	}

	// Initialize the bubbletea application model
	initialModel := NewModel(symbols)

	// Run the application
	p := tea.NewProgram(initialModel)
	if _, err := p.Run(); err != nil {
		log.Fatalf("Failed to start program: %s\n", err)
		os.Exit(1)
	}
}
