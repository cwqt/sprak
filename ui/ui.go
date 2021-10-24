package UI

import (
	"fmt"
	Bus "sprak/bus"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	router Router
	outlet Component
	logs   Component
}

func Create(router Router) model {
	Bus.Publish("log", "Program loaded!")
	Bus.Subscribe("router.navigate", func(event Bus.Event) {
		if paths, ok := event.Data.([]string); ok {
			router.Navigate(paths...)
		}
	})

	// Set the router up to be in Menu straight away
	router.Navigate("")

	return model{
		router: router,
		outlet: router.Outlet.Create(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			fmt.Println("Goodbye!")
			return m, tea.Quit
		}
	}

	m.logs.Update(msg)
	m.outlet.Update(msg)

	return m, nil
}

func (m model) View() string {
	var style = lipgloss.NewStyle().
		Width(PROGRAM_WIDTH).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("8")).
		Align(lipgloss.Center)

	s := ""

	path := *m.router.GetPath()
	for index, segment := range path {
		s += segment
		if index != len(path)-1 {
			s += " / "
		} else {
			s += "\n"
		}
	}

	s += m.outlet.View()

	s += "\nPress q to quit.\n"

	ret := style.Render(s)

	ret += "\n\n" + m.logs.View()

	return ret
}
