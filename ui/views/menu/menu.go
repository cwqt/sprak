package Menu

import (
	"fmt"
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

type model struct {
	cursor int
	items  map[int]menuItem
}

func initialModel(items map[int]menuItem) model {
	return model{
		cursor: 0,
		items:  items,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			view, ok := m.items[m.cursor]
			if ok {
				Views.SwitchTo(view.view)
			}
		}

	}

	return m, nil
}

func (m model) View() string {
	s := titleText

	for i, item := range m.items {
		cursor := " "

		if i == m.cursor {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s\n", cursor, item.label)
	}

	return s
}
