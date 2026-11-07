package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// Model holds the application's state
type model struct {
	count int
}

// Init runs when the program starts
func (m model) Init() tea.Cmd {
	return nil
}

// Update handles incoming events
// These are the buttons I need
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "+":
			m.count++
		case "-":
			m.count--
		case "q", "ctrl+c":
			return m, tea.Quit
		case "r":
			m.count = 0
		}
	}

	return m, nil
}

// View describes what appears in the terminal
func (m model) View() string {
	return fmt.Sprintf(
		"My first Go TUI\n\n"+
			"Counter: %d\n\n"+
			"+ increase | - decrease | r reset | q quit\n",
		m.count,
	)
}

func main() {
	p := tea.NewProgram(model{})

	if _, err := p.Run(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
