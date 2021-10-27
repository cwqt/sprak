package Lesson

import (
	Exercise "sprak/components/lesson/exercise"
	Style "sprak/style"
	UI "sprak/ui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func ProgressBar(model *lessonModel) *UI.Component {
	return &UI.Component{
		Init: func() tea.Cmd {
			return nil
		},
		Update: func(msg tea.Msg) tea.Cmd {
			return nil
		},
		View: func(width int) string {
			totalCorrect := 0
			for _, exercise := range (*model).Exercises {
				if exercise.State == Exercise.Completed {
					totalCorrect++
				}
			}

			// Change the colour of the filled progress bar depending on the current
			// state of the current exercise in progress
			highlight := Style.DarkGray

			if (*model).Current != nil {
				if (*model).Current.State == Exercise.Completed {
					highlight = Style.Green
				} else if (*model).Current.State == Exercise.Failed {
					highlight = Style.Green
				}
			}

			// The completed amount
			s := lipgloss.NewStyle().
				Width(totalCorrect).
				Background(highlight).
				Render("")

			// The remainder
			s += lipgloss.NewStyle().
				Width(width - totalCorrect).
				Background(Style.DarkGray).
				Render("")

			return s
		},
		Destroy: func() {},
	}
}
