package UI

import (
	"fmt"
	Bus "sprak/bus"
	Events "sprak/bus/events"
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

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// fmt.Println("updating...")

	switch msg := msg.(type) {

	case Bus.Event:
		switch msg.Topic {
		case "view:change":

			if viewChange, ok := msg.Data.(Events.ChangeView); ok {
				// fmt.Println("changing view to", viewChange.To)
				m.currentView = viewChange.To
			}
		}

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			fmt.Println("Goodbye!")
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	s := "\nPress q to quit.\n"

	if m.currentView == Views.Home {
		s += "currently home!"
	} else {
		s += "currently away!"
	}
	return s
}
