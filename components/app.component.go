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
			cmds := UI.Cmds()

			cmds.Append(m.logger.Update(msg))
			cmds.Append(props.Outlet.Update(msg))

			return cmds.AsCmd()
		},
		View: func(width int) string {
			s := "språk 0.1\n"
			s += Wrapper.Width(width).Render(props.Outlet.View(width - 4))
			s += m.logger.View(width)
			return s
		},
		Destroy: props.Outlet.Destroy,
	}
}
