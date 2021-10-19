package Component

import tea "github.com/charmbracelet/bubbletea"

type Component struct {
	Name    string
	Init    func() tea.Cmd
	Update  func(msg tea.Msg) tea.Cmd
	View    func() string
	Destroy func()
}
