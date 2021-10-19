package Menu

import (
	"fmt"
	Bus "sprak/bus"
	Component "sprak/ui/component"
	Views "sprak/ui/views"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// ensure each line is 26 chars across
const titleText string = `
               ( )   |   
(_-<  _ \   _| _` + "`" + ` |  | / 
___/ .__/ _| \__,_| _\_\ 
    _|                   
`

type menuItem struct {
	label  string
	view   Views.View
	action func()
}

type Model struct {
	cursor int
	items  [2]menuItem
}

func Create() Component.Component {
	m := func() Model {
		items := [2]menuItem{{
			label: "Lesson",
			view:  Views.Lesson,
			action: func() {
				Views.SwitchTo(Views.Lesson)
			},
		}, {
			label: "Refresh deck",
			view:  Views.Menu,
			action: func() {
				Views.SwitchTo(Views.Menu)
			},
		}}

		return Model{cursor: 0, items: items}
	}()

	return Component.Component{
		Name: "MenuComponent",
		Init: func() tea.Cmd {
			return nil
		},
		Update: func(msg tea.Msg) tea.Cmd {
			switch msg := msg.(type) {
			case tea.KeyMsg:
				switch msg.String() {
				case "up":
					Bus.Publish("log", "Moved cursor up on menu")
					if m.cursor > 0 {
						m.cursor--
					}
				case "down":
					Bus.Publish("log", "Moved cursor down on menu")
					if m.cursor < len(m.items)-1 {
						m.cursor++
					}
				case "enter":
					Bus.Publish("log", "Selected menu at position "+strconv.Itoa(m.cursor))
					m.items[m.cursor].action()
				}
			}

			return nil
		},
		View: func() string {
			var style = lipgloss.NewStyle().
				Width(60).
				Align(lipgloss.Center)

			s := style.Render(titleText)

			s += "\nspråk, duolingo on the cli\n\n"

			for i, item := range m.items {
				cursor := " "
				if i == m.cursor {
					cursor = ">"
				}

				s += fmt.Sprintf("%s %s\n", cursor, item.label)
			}

			return s
		},
	}
}
