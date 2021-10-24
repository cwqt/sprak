package main

import (
	"fmt"
	"os"
	Bus "sprak/bus"
	UI "sprak/ui"

	tea "github.com/charmbracelet/bubbletea"
)

var Router = UI.CreateRouter(Routes)

func main() {
	fmt.Println("Starting språk..")
	done := make(chan struct{})

	p := tea.NewProgram(UI.Create(Router), tea.WithAltScreen())
	Bus.AttachToProgram(p)

	go func() {
		if err := p.Start(); err != nil {
			fmt.Println("Error running program:", err)
			os.Exit(1)
		}
		close(done)
	}()

	<-done
}
