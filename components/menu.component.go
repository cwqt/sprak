package Component

import (
	"fmt"
	Bus "sprak/bus"
	UI "sprak/ui"
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
	action func()
}

type menuModel struct {
	cursor int
	items  [2]menuItem
}

func Menu() UI.Component {
	m := func() menuModel {
		items := [2]menuItem{{
			label: "Lesson",
			action: func() {
				Bus.Publish("router.navigate", []string{"lesson"})
			},
		}, {
			label: "Import deck",
			action: func() {
				Bus.Publish("router.navigate", []string{"import"})
			},
		}}

		return menuModel{cursor: 0, items: items}
	}()

	return UI.Component{
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
