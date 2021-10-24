package Component

import (
	"fmt"
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

func Lesson() UI.Component {
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

	Bus.Publish("log", fmt.Sprintf("%+v\n", m))

	return UI.Component{
		Init: func() tea.Cmd {
			return nil
		},
		Update: func(msg tea.Msg) tea.Cmd {
			return nil
		},
		View: func() string {
			return "lesson! in state " + strconv.Itoa(int(m.state))
		},
	}
}
