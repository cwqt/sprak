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

var logStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("8"))
var rerenderStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("163"))

func Log(props UI.Props) *UI.Component {
	m := logModel{
		logs: make([]string, 0),
	}

	unsubscribe := Bus.Subscribe("log", func(event Bus.Event) {
		if len(m.logs) > 30 {
			m.logs = m.logs[1:]
		}

		if event.Topic == "log" {
			// muted colours for log lines
			m.logs = append(m.logs, logStyle.Render(fmt.Sprintf("%s", event.Data)))
		} else if event.Topic == "re:render" {
			m.logs = append(m.logs, rerenderStyle.Render("RE-RENDER!"))

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
			s := "\n"

			if len(m.logs) == 0 {
				s += "No logs (yet)"
			} else {
				// Add all logs onto s, reverse order: newest at top
				for i := len(m.logs) - 1; i >= 0; i-- {
					s += fmt.Sprintf("%s\n", m.logs[i])
				}
			}

			return s
		},
		Destroy: func() {
			unsubscribe()
		},
	}
}
