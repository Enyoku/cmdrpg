package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(InitialModel())

	_, err := p.Run()
	if err != nil {
		fmt.Printf("Alas, there`s been an error: %v", err)
		os.Exit(1)
	}
}