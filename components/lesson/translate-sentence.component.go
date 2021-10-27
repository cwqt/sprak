package Lesson

import (
	"fmt"
	Bus "sprak/bus"
	Exercise "sprak/components/lesson/exercise"
	UI "sprak/ui"

	tea "github.com/charmbracelet/bubbletea"
)

type translateSentenceModel struct {
	Expected string
	Current  string
}

func TranslateSentenceComponent(props *UI.Props) *UI.Component {
	m := translateSentenceModel{}

	return &UI.Component{
		Init: func() tea.Cmd {
			return nil
		},
		Update: func(msg tea.Msg) tea.Cmd {
			switch msg := msg.(type) {
			case Bus.Event:
				switch msg.Topic {
				case "exercise:started":
					if exercise, ok := msg.Data.(*Exercise.Exercise); ok {
						if data, ok := exercise.Data.(Exercise.TranslateSentenceData); ok {
							m.Expected = data.Expected
						}
					}
				}
			}

			return nil
		},
		View: func(width int) string {
			s := "Translate the following sentence:\n\n"

			s += fmt.Sprintf("\t%s\n", m.Expected)

			return s
		},
		Destroy: func() {
		},
	}
}
