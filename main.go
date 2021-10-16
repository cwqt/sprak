package main

import (
	"fmt"
	"os"

	UI "sprak/ui"

	tea "github.com/charmbracelet/bubbletea"
	// tea "github.com/charmbracelet/bubbletea"
)

type X struct {
	T string
}

func main() {
	p := tea.NewProgram(UI.InitialModel())

	// Bus.Subscribe("test", func(event Bus.Event) {
	// 	fmt.Println("%s %+v\n", event.Topic, event.Data)
	// 	fmt.Print(p)
	// 	p.Send(&event)
	// })

	// Bus.Publish("test", X{T: "hello"})

	// Bus.Subscribe("view:change", func(event *Bus.Event) {
	// 	p.Send(event)
	// })

	// Bus.Publish("view:change", UI.ViewChange{To: Views.Menu})

	// p.Send("hello")

	// Bus.Subscribe("view:change", func(value UI.ViewChange) {
	// })

	if err := p.Start(); err != nil {
		fmt.Println("Uh oh", err)
		os.Exit(1)
	}

	p.Send(&X{T: "sting"})
	p.Send(&X{T: "sting"})
	p.Send(&X{T: "sting"})
	p.Send(&X{T: "sting"})

	// Bus.Publish("test", X{T: "hello"})

	// p.Send(tea.Quit())
	//
	// if err := Data.Connect(); err != nil {
	// 	os.Exit(1)
	// }

	// if _, err := Anki.ImportApkg("Bokmål.apkg"); err != nil {
	// 	os.Exit(1)
	// }

}
