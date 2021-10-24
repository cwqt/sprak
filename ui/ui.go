package UI

import (
	"fmt"
	Bus "sprak/bus"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	router *Router
	outlet *Component
}

func Create(router *Router) model {
	Bus.Publish("log", "Program loaded!")
	Bus.Subscribe("router.navigate", func(event Bus.Event) {
		if paths, ok := event.Data.([]string); ok {
			router.Navigate(paths...)
		}

	})

	return model{
		router: router,
		outlet: router.Outlet,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "backspace":
			Bus.Publish("router.navigate", []string{"index", "menu"})
			return m, func() tea.Msg {
				return 0
			}
		case "ctrl+c", "q":
			fmt.Println("Goodbye!")
			return m, tea.Quit
		}
	}

	return m, m.outlet.Update(msg)
}

func (m model) View() string {
	s := ""

	s += "[ "
	path := *m.router.GetPath()
	for index, segment := range path {
		s += segment
		if index != len(path)-1 {
			s += " / "
		} else {
			s += " ]\n"
		}
	}

	return s + m.outlet.View()
}
