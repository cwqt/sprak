package UI

import (
	"fmt"
	Bus "sprak/bus"
)

type RoutingTable map[string]Route

type Route struct {
	Create   func(props *Props) *Component
	Children RoutingTable
}

// Private to Create() closure -- its internal model
type routerModel struct {
	Path   []string     // active router path
	Routes RoutingTable // routing structure with factory functions
}

// Public methods
type Router struct {
	GetRoute func(path ...string) *Route
	GetPath  func() *[]string
	Navigate func(paths ...string)
	Outlet   *Component // the root outlet
}

// Router instance -- one per Program
func CreateRouter(routes RoutingTable, initialPath ...string) Router {
	router := routerModel{
		Path:   make([]string, 0),
		Routes: routes,
	}

	if len(initialPath) > 0 {
		router.Path = append(router.Path, initialPath...)
	}

	return Router{
		Outlet: CreateOutlet(router.Routes, router.Path...),
		GetRoute: func(paths ...string) *Route {
			var find func(routes RoutingTable, paths ...string) *Route
			find = func(routes RoutingTable, paths ...string) *Route {
				head := paths[0]

				if len(paths) == 0 {
					if route, ok := routes[head]; ok {
						return &route
					}
				}

				return find(routes[head].Children, paths[1:]...)
			}

			return find(router.Routes, paths...)
		},
		GetPath: func() *[]string {
			return &router.Path
		},
		Navigate: func(path ...string) {
			Bus.Log(fmt.Sprintf("%+v", path))
			copy(router.Path, path)
		},
	}
}
