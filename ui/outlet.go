package UI

import (
	"fmt"
	Bus "sprak/bus"

	tea "github.com/charmbracelet/bubbletea"
)

type outletModel struct {
	active map[string]*Component
}

type Outlet struct {
	Create func() *Component
}

func CreateOutlet(routing RoutingTable, paths *[]string, depth int) *Component {
	m := outletModel{
		active: map[string]*Component{},
	}

	return &Component{
		Model: &m,
		Init: func() tea.Cmd {
			Bus.Log(fmt.Sprintf("%+v", m))
			return nil
		},
		Update: func(msg tea.Msg) tea.Cmd {
			cmds := make([]tea.Cmd, 0)

			if len(*paths) > 0 {
				head := (*paths)[depth]

				if route, ok := m.active[head]; ok {
					if cmd := route.Update(msg); cmd != nil {
						cmds = append(cmds, cmd)
					}
				} else {
					for path, route := range m.active {
						if path != head {
							Bus.Log(fmt.Sprintf("destroying: %s", path))
							route.Destroy()
							delete(m.active, path)
						}
					}

					if route, ok := routing[head]; ok {
						m.active[head] = route.Create(&Props{
							Outlet: CreateOutlet(route.Children, paths, depth+1),
						})

						if cmd := m.active[head].Init(); cmd != nil {
							cmds = append(cmds, cmd)
						}
					}
				}
			}

			return tea.Batch(cmds...)
		},
		View: func() string {
			s := ""

			for _, path := range *paths {
				if component, ok := m.active[path]; ok {
					s += component.View()
				}
			}

			return s
		},
		Destroy: func() {
			for _, component := range m.active {
				component.Destroy()
			}
		},
	}
}
