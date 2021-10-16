package UI

import (
	"fmt"
	// Bus "sprak/bus"

	Views "sprak/ui/views"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	currentView Views.View
}

func InitialModel() model {
	return model{
		currentView: Views.Menu,
	}
}

type ViewChange struct {
	To Views.View
}

func (m model) Init() tea.Cmd {

	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	fmt.Println("printing main")

	switch msg := msg.(type) {

	// case Bus.Event:
	// 	fmt.Printf("%s, %+v\n", "topic hit", msg)
	// 	switch msg.Topic {
	// 	case "view:change":
	// 		if viewChange, ok := msg.Data.(ViewChange); ok {
	// 			fmt.Println("changing view to", viewChange.To)
	// 			m.currentView = viewChange.To
	// 		} else {
	// 			fmt.Println("not ok!")
	// 		}
	// 	}

	case tea.KeyMsg:
		switch msg.String() {
		case "1":
			// Bus.Publish("view:change", ViewChange{To: Views.Menu})
		// case "2":
		// Bus.Publish("view:change", ViewChange{To: Views.Home})
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	s := ""

	if m.currentView == Views.Home {
		s += "Welcome home! 💚 "
	}

	if m.currentView == Views.Menu {
		s += "this is the menu "
	}

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}
