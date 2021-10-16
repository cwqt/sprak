package main

import (
	"fmt"
	"os"
	"time"

	Bus "sprak/bus"
	UI "sprak/ui"

	tea "github.com/charmbracelet/bubbletea"
	// tea "github.com/charmbracelet/bubbletea"
)

func main() {

	done := make(chan struct{})

	p := tea.NewProgram(UI.InitialModel())

	Bus.Subscribe("test", func() {
		p.Send("pushed!")
	})

	go func() {
		if err := p.Start(); err != nil {
			fmt.Println("Error running program:", err)
			os.Exit(1)
		}
		close(done)
	}()

	// Simulate activity
	go func() {
		for {
			time.Sleep(1e9)
			Bus.Publish("test", 1)
		}
	}()

	<-done
}
