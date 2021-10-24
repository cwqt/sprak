package UI

type Route struct {
	Create   func() Component
	Children map[string]Route
}

// Private to Create() closure -- its internal model
type routerModel struct {
	Path   []string         // active router path
	Routes map[string]Route // routing structure with factory functions
}

// Public methods
type Router struct {
	GetPath  func() *[]string
	Outlet   Outlet
	Navigate func(paths ...string)
}

type Outlet struct {
	Create func() Component
}

// Router instance -- one per Program
func CreateRouter(routes map[string]Route) Router {
	router := routerModel{
		Path:   make([]string, 0),
		Routes: routes,
	}

	return Router{
		GetPath: func() *[]string {
			return &router.Path
		},
		Navigate: func(path ...string) {
			router.Path = path
		},
		Outlet: Outlet{
			Create: func() Component {
				return CreateOutlet(router)
			},
		},
	}
}

// Given a router with a root node, return a new root node containing only
// the nodes listed in the path
//     __1__
//    /     \
//   2      _3_
//   |     / | \
//   4    5  6  7
//
// path  -> [1,3]
// route –> [1,[3,[5,6,7]]]

// func construct(router *Router) Route {
// 	// the complete route path
// 	var active = Route{
// 		children: map[string]*Route{},
// 	}

// 	// track current tail of the route in construction
// 	var tail *Route = &active

// 	var find func(route map[string]*Route, path []string)
// 	find = func(route map[string]*Route, path []string) {
// 		head := path[0]

// 		// lookup to see if head exists in current set of children
// 		if r, ok := route[head]; ok {
// 			// last element in path -- end of tree
// 			if head == path[len(path)-1] {
// 				// end of the tree
// 				tail.children[head] = r // the rest of the route
// 				return
// 			} else {
// 				tail.children[head] = r.children[head]

// 				// tail.children[head]
// 				find(tail.children, path[1:])
// 			}
// 		} else {
// 			panic("Couldn't find node in current router tree!")
// 		}
// 	}

// 	root := map[string]*Route{
// 		"root": router.root,
// 	}

// 	find(root, append([]string{"root"}, router.state...))

// 	return *active.children["root"]
// }
