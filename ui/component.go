package UI

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Component struct {
	Model   interface{}
	Init    func() tea.Cmd
	Update  func(msg tea.Msg) tea.Cmd
	View    func(width int) string
	Destroy func()
}

type Props struct {
	Outlet *Component
}
