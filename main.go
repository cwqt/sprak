package main

import (
	"fmt"
	"os"
	Bus "sprak/bus"
	Events "sprak/bus/events"
	UI "sprak/ui"
	Views "sprak/ui/views"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type X struct{ T int }

func main() {
	done := make(chan struct{})

	p := tea.NewProgram(UI.InitialModel())

	Bus.Subscribe("view:change", func(event Bus.Event) {
		fmt.Printf("sending event to tea %+v\n", event)
		p.Send(event)
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
			Bus.Publish("view:change", Events.ChangeView{To: Views.Home})
		}
	}()

	<-done
}
