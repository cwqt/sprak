package UI

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Component struct {
	Init    func() tea.Cmd
	Update  func(msg tea.Msg) tea.Cmd
	View    func() string
	Destroy func()
}
