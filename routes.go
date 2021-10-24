package main

import (
	Component "sprak/components"
	UI "sprak/ui"
)

var Routes = map[string]UI.Route{
	"": {
		Create: Component.App,
		Children: map[string]UI.Route{
			"menu": {
				Create: Component.Menu,
			},
			"lesson": {
				Create: Component.Lesson,
				Children: map[string]UI.Route{
					"translate-sentence": {
						Create: Component.TranslateSentenceComponent,
					},
					"listening": {
						Create: Component.TranslateSentenceComponent,
					},
					"multiple-choice": {
						Create: Component.TranslateSentenceComponent,
					},
				},
			},
			"import": {
				Create: Component.Import,
			},
		},
	},
}
