package Log

import (
	"fmt"
	Bus "sprak/bus"
	Component "sprak/ui/component"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	logs []string
}

func Create() Component.Component {
	m := Model{
		logs: make([]string, 0),
	}

	unsubsribe := Bus.Subscribe("log", func(event Bus.Event) {
		m.logs = append(m.logs, fmt.Sprintf("%+v", event))
	})

	return Component.Component{
		Init: func() tea.Cmd { return nil },
		Update: func(msg tea.Msg) tea.Cmd {
			return nil
		},
		View: func() string {
			s := ""

			if len(m.logs) == 0 {
				s += "No logs (yet)"
			} else {
				// Add all logs onto
				for _, v := range m.logs {
					s += fmt.Sprintf("%s\n", v)
				}
			}

			return s
		},
		Destroy: func() {
			unsubsribe()
		},
	}
}
