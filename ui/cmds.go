package UI

import tea "github.com/charmbracelet/bubbletea"

type cmd struct {
	Append func(tea.Cmd)
	AsCmd  func() tea.Cmd
}

func Cmds(initialCmds ...tea.Cmd) cmd {
	cmds := make([]tea.Cmd, len(initialCmds))
	for i := range initialCmds {
		if initialCmds[i] != nil {
			cmds[i] = initialCmds[i]
		}
	}

	return cmd{
		Append: func(cmd tea.Cmd) {
			if cmd != nil {
				cmds = append(cmds, cmd)
			}
		},
		AsCmd: func() tea.Cmd {
			if len(cmds) > 0 {
				return tea.Batch(cmds...)
			} else {
				return nil
			}
		},
	}
}
