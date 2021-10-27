package main

import (
	"fmt"
	"os"
	Bus "sprak/bus"
	Component "sprak/components"
	Lesson "sprak/components/lesson"
	Data "sprak/data"
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
					Create: Lesson.Lesson,
					Children: UI.RoutingTable{
						"translate-sentence": {
							Create: Lesson.TranslateSentenceComponent,
						},
						"listening": {
							Create: Lesson.TranslateSentenceComponent,
						},
						"multiple-choice": {
							Create: Lesson.TranslateSentenceComponent,
						},
					},
				},
				"import": {
					Create: Component.Import,
				},
			},
		},
	}, "index", "menu")

	Data.Connect()

	p := tea.NewProgram(UI.Create(&Router), tea.WithAltScreen())
	Bus.AttachToProgram(p)

	if err := p.Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
