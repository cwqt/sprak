package Lesson

import (
	"fmt"
	Bus "sprak/bus"
	Exercise "sprak/components/lesson/exercise"
	Data "sprak/data"
	"sprak/db"
	UI "sprak/ui"

	tea "github.com/charmbracelet/bubbletea"
)

type lessonModel struct {
	Exercises []Exercise.Exercise
	Current   *Exercise.Exercise
}

func Lesson(props *UI.Props) *UI.Component {
	Bus.Publish("log", "Creating Lesson component")

	// Get first 20 cards
	cards, err := Data.GetCards(20)
	if err != nil {
		// Return an empty set if errored
		cards = make([]db.CardModel, 0)
	}

	exercises := make([]Exercise.Exercise, len(cards))
	for index, card := range cards {
		exercise := Exercise.Exercise{
			State: Exercise.Idle,
			Data: Exercise.TranslateSentenceData{
				Expected: card.Target,
			},
		}

		exercises[index] = exercise
	}

	m := lessonModel{
		Exercises: exercises,
		Current:   nil,
	}

	progressBar := ProgressBar(&m)

	return &UI.Component{
		Init: func() tea.Cmd {
			start := func() tea.Msg {
				for _, exercise := range m.Exercises {
					Bus.Log(fmt.Sprintf("%+v", exercise.Data))
				}

				Bus.Publish("exercise:started", &exercises[0])
				return nil
			}

			return UI.Cmds(progressBar.Init(), props.Outlet.Init(), start).AsCmd()
		},
		Update: func(msg tea.Msg) tea.Cmd {
			cmds := UI.Cmds()

			switch msg := msg.(type) {
			case tea.KeyMsg:
				switch msg.String() {
				case "1":
					Bus.Publish("router.navigate", []string{"index", "lesson", "listening"})
				}
			}

			cmds.Append(props.Outlet.Update(msg))

			return cmds.AsCmd()
		},
		View: func(width int) string {
			s := progressBar.View(width) + "\n\n"

			s += props.Outlet.View(width)

			return s
		},
		Destroy: func() {
			props.Outlet.Destroy()
		},
	}
}
