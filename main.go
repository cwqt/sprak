package main

import (
	"fmt"
	"os"
	Bus "sprak/bus"
	Component "sprak/components"
	UI "sprak/ui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	fmt.Println("Starting språk..")

	Router := UI.CreateRouter(UI.RoutingTable{
		"index": {
			Create: Component.App,
			Children: UI.RoutingTable{
				"menu": {
					Create: Component.Menu,
				},
				"lesson": {
					Create: Component.Lesson,
					Children: UI.RoutingTable{
						"translate-sentence": {
							Create: Component.TranslateSentenceComponent,
						},
						"listening": {
							Create: Component.TranslateSentenceComponent,
						},
						"multiple-choice": {
							Create: Component.TranslateSentenceComponent,
						},
					},
				},
				"import": {
					Create: Component.Import,
				},
			},
		},
	}, "index", "menu")

	p := tea.NewProgram(UI.Create(&Router), tea.WithAltScreen())
	Bus.AttachToProgram(p)

	if err := p.Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
