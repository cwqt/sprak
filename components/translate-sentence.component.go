package Component

import (
	UI "sprak/ui"

	tea "github.com/charmbracelet/bubbletea"
)

func TranslateSentenceComponent(props *UI.Props) *UI.Component {
	return &UI.Component{
		Init: func() tea.Cmd {
			return nil
		},
		Update: func(msg tea.Msg) tea.Cmd {
			return nil
		},
		View: func() string {
			return "translate the sentence!"
		},
		Destroy: func() {
		},
	}
}
