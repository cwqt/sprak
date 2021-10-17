package Views

import (
	Bus "sprak/bus"
)

type View string

const (
	Menu   View = "menu"
	Lesson View = "lesson"
)

type ChangeViewEvent struct {
	To View
}

func SwitchTo(view View) {
	Bus.Publish("view:change", ChangeViewEvent{To: view})
}
