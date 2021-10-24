package UI

import (
	"fmt"
	Bus "sprak/bus"

	tea "github.com/charmbracelet/bubbletea"
)

type outletModel struct {
	Routes map[string]Component // active routed component tree
}

func CreateOutlet(router routerModel) Component {
	m := outletModel{
		Routes: make(map[string]Component),
	}

	return Component{
		Init: func() tea.Cmd {
			Bus.Log(fmt.Sprintf("%+v", m))
			return nil
		},
		Update: func(msg tea.Msg) tea.Cmd {
			path := router.Path

			Bus.Log(fmt.Sprint(path))

			if len(path) > 0 {
				head := path[0]

				if route, ok := m.Routes[head]; ok {
					route.Update(msg)
				} else {
					if component, ok := router.Routes[head]; ok {
						m.Routes[head] = component.Create()
						m.Routes[head].Init()
					}
				}
			}

			return nil
		},
		View: func() string {
			s := ""

			for _, path := range router.Path {
				if component, ok := m.Routes[path]; ok {
					s += component.View()
				}
			}

			return s
		},
		Destroy: func() {
			for _, component := range m.Routes {
				component.Destroy()
			}
		},
	}
}
