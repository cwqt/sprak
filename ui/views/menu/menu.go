package Menu

import (
	"fmt"
	Component "sprak/ui/component"
	Views "sprak/ui/views"

	tea "github.com/charmbracelet/bubbletea"
)

const titleText string = `
                ( )   |   
 (_-<  _ \   _| _` + "`" + ` |  | / 
 ___/ .__/ _| \__,_| _\_\ 
     _|                            
`

type menuItem struct {
	label string
	view  Views.View
}

type Model struct {
	cursor int
	items  [2]menuItem
}

func Create() Component.Component {
	m := func() Model {
		items := [2]menuItem{{
			label: "Menu",
			view:  Views.Menu,
		}, {
			label: "Lesson",
			view:  Views.Lesson,
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
					if m.cursor > 0 {
						m.cursor--
					}
				case "down":
					if m.cursor < len(m.items)-1 {
						m.cursor++
					}
				case "enter":
					view := m.items[m.cursor]
					Views.SwitchTo(view.view)
				}
			}

			return nil
		},
		View: func() string {
			s := titleText

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
