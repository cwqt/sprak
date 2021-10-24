package Component

import (
	UI "sprak/ui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type appModel struct {
	logger *UI.Component
}

var style = lipgloss.NewStyle().
	Width(PROGRAM_WIDTH).
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("8")).
	Align(lipgloss.Center)

func App(props *UI.Props) *UI.Component {
	m := appModel{
		logger: Log(UI.Props{}),
	}

	return &UI.Component{
		Init: func() tea.Cmd {
			return nil
		},
		Update: func(msg tea.Msg) tea.Cmd {
			m.logger.Update(msg)

			return props.Outlet.Update(msg)
		},
		View: func() string {
			s := "språk 0.1\n"
			s += style.Render(props.Outlet.View())
			s += m.logger.View()
			return s
		},
		Destroy: func() {
		},
	}
}
