package Component

import (
	UI "sprak/ui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type appModel struct {
	logger *UI.Component
}

var Wrapper = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("8")).
	Align(lipgloss.Center)

func App(props *UI.Props) *UI.Component {
	m := appModel{
		logger: Log(UI.Props{}),
	}

	return &UI.Component{
		Init: props.Outlet.Init,
		Update: func(msg tea.Msg) tea.Cmd {
			cmds := make([]tea.Cmd, 0)

			if cmd := m.logger.Update(msg); cmd != nil {
				cmds = append(cmds, cmd)
			}

			return tea.Batch(append(cmds, props.Outlet.Update(msg))...)
		},
		View: func() string {
			s := "språk 0.1\n"
			s += Wrapper.Width(PROGRAM_WIDTH).Render(props.Outlet.View())
			s += m.logger.View()
			return s
		},
		Destroy: props.Outlet.Destroy,
	}
}
