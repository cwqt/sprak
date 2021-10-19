package Lesson

import (
	Bus "sprak/bus"
	Data "sprak/data"
	"sprak/db"
	"strconv"

	Component "sprak/ui/component"
	Exercise "sprak/ui/views/lesson/exercises"

	tea "github.com/charmbracelet/bubbletea"
)

type State int

// Current state of the lesson
const (
	Idle State = iota
	Failure
	Success
)

type Model struct {
	state    State
	cards    []db.CardModel
	exercise Exercise.Current
}

func Create() Component.Component {
	Bus.Publish("log", "Creating Lesson component")

	// Get first 20 cards
	cards, err := Data.GetCards(20)
	if err != nil {
		// Return an empty set if errored
		cards = make([]db.CardModel, 0)
	}

	m := Model{
		state: Idle,
		cards: cards,
	}

	return Component.Component{
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
