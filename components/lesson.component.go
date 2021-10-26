package Component

import (
	Bus "sprak/bus"
	Data "sprak/data"
	"sprak/db"
	UI "sprak/ui"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
)

type State int

// Current state of the lesson
const (
	Idle State = iota
	Failure
	Success
)

type lessonModel struct {
	state State
	cards []db.CardModel
	// exercise Exercise
}

func Lesson(props *UI.Props) *UI.Component {
	Bus.Publish("log", "Creating Lesson component")

	// Get first 20 cards
	cards, err := Data.GetCards(20)
	if err != nil {
		// Return an empty set if errored
		cards = make([]db.CardModel, 0)
	}

	m := lessonModel{
		state: Idle,
		cards: cards,
	}

	return &UI.Component{
		Init: func() tea.Cmd {
			return props.Outlet.Init()
		},
		Update: func(msg tea.Msg) tea.Cmd {
			switch msg := msg.(type) {
			case tea.KeyMsg:
				switch msg.String() {
				case "1":
					Bus.Publish("router.navigate", []string{"index", "lesson", "listening"})
				}
			}

			return props.Outlet.Update(msg)
		},
		View: func() string {
			s := Wrapper.Render(props.Outlet.View()) + "\n"
			s += "lesson! in state " + strconv.Itoa(int(m.state))

			return s
		},
		Destroy: func() {
			props.Outlet.Destroy()
		},
	}
}
