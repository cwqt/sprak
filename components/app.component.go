package Component

import (
	UI "sprak/ui"

	"github.com/charmbracelet/bubbles/spinner"

	tea "github.com/charmbracelet/bubbletea"
)

type appModel struct {
	importStatus       int // 0 ok, 1 err, nil idle
	totalCards         int
	currentCardsLoaded int
	spinner            spinner.Model
}

func App() UI.Component {

	return UI.Component{
		Init: func() tea.Cmd {
			return nil
		},
		Update: func(msg tea.Msg) tea.Cmd {
			return nil
		},
		View: func() string {

			return ""
		},
		Destroy: func() {
		},
	}
}
