package UI

import (
	"fmt"
	Bus "sprak/bus"
	Component "sprak/ui/component"
	Views "sprak/ui/views"
	Lesson "sprak/ui/views/lesson"
	Logs "sprak/ui/views/log"
	Menu "sprak/ui/views/menu"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"syscall"
	"unsafe"
)

type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

func getWidth() uint {
	ws := &winsize{}
	retCode, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)))

	if int(retCode) == -1 {
		panic(errno)
	}
	return uint(ws.Col)
}

var PROGRAM_WIDTH = getWidth() - 2
var PROGRAM_HEIGHT = 32

type model struct {
	view Views.View

	menu   *Component.Component
	lesson *Component.Component
	logs   *Component.Component
}

func Create() model {
	menu := Menu.Create()
	logs := Logs.Create()

	Bus.Publish("log", "Program loaded!")

	return model{
		view:   Views.Menu,
		menu:   &menu,
		lesson: nil,
		logs:   &logs,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case Bus.Event:
		switch msg.Topic {
		case "view:change":
			if event, ok := msg.Data.(Views.ChangeViewEvent); ok {
				m.view = event.To
			}
		}

	case tea.KeyMsg:
		switch msg.String() {
		case "1":
			Views.SwitchTo(Views.Menu)
		case "ctrl+c", "q":
			fmt.Println("Goodbye!")
			return m, tea.Quit
		}
	}

	m.logs.Update(msg)

	switch m.view {
	case Views.Menu:
		if m.menu == nil {
			menu := Menu.Create()
			m.menu = &menu
		}
		m.menu.Update(msg)
	case Views.Lesson:
		if m.lesson == nil {
			lesson := Lesson.Create()
			m.lesson = &lesson
		}
		m.lesson.Update(msg)
	}

	return m, nil
}

func (m model) View() string {
	var style = lipgloss.NewStyle().
		Width(int(PROGRAM_WIDTH)).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("8")).
		Align(lipgloss.Center)

	s := ""

	// s += m.menu.View()

	switch m.view {
	case Views.Menu:
		s += m.menu.View()
	case Views.Lesson:
		s += m.lesson.View()
	}

	// s += m.children.View()

	// for _, child := range m.children {
	// 	if child.Name == m.currentView {
	// 		s += child.View()
	// 	}
	// }

	s += "\nPress q to quit.\n"

	ret := style.Render(s)

	ret += "\n\n" + m.logs.View()

	return ret
}
