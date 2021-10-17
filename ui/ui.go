package UI

import (
	"fmt"
	Bus "sprak/bus"
	Component "sprak/ui/component"
	Views "sprak/ui/views"
	Menu "sprak/ui/views/menu"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	view Views.View

	menu Component.Component
}

func Create() model {
	return model{
		view: Views.Menu,
		menu: Menu.Create(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case Bus.Event:
		switch msg.Topic {
		case "view:change":
			if event, ok := msg.Data.(Views.ChangeViewEvent); ok {
				m.view = event.To
			}
		}

	case tea.KeyMsg:
		switch msg.String() {
		case "1":
			Views.SwitchTo(Views.Menu)
		case "ctrl+c", "q":
			fmt.Println("Goodbye!")
			return m, tea.Quit
		}
	}

	m.menu.Update(msg)

	// m.children.Update(msg)

	// for _, child := range m.children {
	// 	if child.Name == m.currentView {
	// 		child.Update(msg)
	// 	}
	// }

	return m, nil
}

func (m model) View() string {
	s := ""

	s += m.menu.View()

	// s += m.children.View()

	// for _, child := range m.children {
	// 	if child.Name == m.currentView {
	// 		s += child.View()
	// 	}
	// }

	s += "\nPress q to quit.\n"

	return s
}
