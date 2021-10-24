package Component

import (
	"fmt"
	Bus "sprak/bus"
	UI "sprak/ui"

	"github.com/charmbracelet/lipgloss"

	tea "github.com/charmbracelet/bubbletea"
)

type logModel struct {
	logs []string
}

var logStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("81"))

func Log(props UI.Props) *UI.Component {
	m := logModel{
		logs: make([]string, 0),
	}

	unsubscribe := Bus.Subscribe("log", func(event Bus.Event) {
		if event.Topic == "log" {
			m.logs = append(m.logs, logStyle.Render(fmt.Sprintf("%s", event.Data)))
		} else {
			m.logs = append(m.logs, fmt.Sprintf("%+v", event))
		}
	})

	return &UI.Component{
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
			unsubscribe()
		},
	}
}
