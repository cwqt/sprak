package Views

type View string

const (
	Menu View = "menu"
	Home View = "home"
)

func SwitchTo(view View) {
	// Bus.Publish("view:change", view)
}
