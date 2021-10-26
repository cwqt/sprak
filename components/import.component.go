package Component

import (
	"fmt"
	Bus "sprak/bus"
	UI "sprak/ui"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/lipgloss"

	tea "github.com/charmbracelet/bubbletea"
)

type importModel struct {
	importStatus       int // 0 ok, 1 err, nil idle
	currentCardsLoaded int
	spinner            spinner.Model
}

type importError struct{}
type importSuccess struct{}

func Import(props *UI.Props) *UI.Component {
	Bus.Publish("log", "Creating Refresh Deck component")

	// Create the spinner to be shown
	s := spinner.NewModel()
	s.Spinner = spinner.Line
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	m := importModel{
		currentCardsLoaded: 0,
		spinner:            s,
	}

	Bus.Log(fmt.Sprintf("%+v\n", m))

	return &UI.Component{
		Model: m,
		Init: func() tea.Cmd {
			// startImport := func() tea.Msg {
			// 	Bus.Log("starting import...")
			// 	if _, err := anki.ImportApkg("/Users/cass/code/sprak/deck.apkg"); err != nil {
			// 		Bus.Err(err.Error())
			// 		return importError{}
			// 	} else {
			// 		Bus.Log("import success...")
			// 		return importSuccess{}
			// 	}
			// }
			tea.Every(time.Second, func(t time.Time) tea.Msg {
				return spinner.Tick
			})
			return tea.Batch(spinner.Tick)
		},
		Update: func(msg tea.Msg) tea.Cmd {
			switch msg := msg.(type) {
			case Bus.Event:
				switch msg.Topic {
				case "card:upserted":
					if _, ok := msg.Data.(int); ok {
						m.currentCardsLoaded += 1
					}
				}
			case importSuccess:
				m.importStatus = 1
			case importError:
				m.importStatus = 0
			case spinner.TickMsg:
				var cmd tea.Cmd
				m.spinner, cmd = m.spinner.Update(msg)
				Bus.Log(fmt.Sprintf("%+v", m.spinner))
				return cmd
			}

			return nil
		},
		View: func() string {
			s := ""
			s += fmt.Sprintf("%s Importing %d cards", m.spinner.View(), m.currentCardsLoaded)
			if m.importStatus == 1 {
				s += "COMPLETE"
			}
			return s
		},
		Destroy: func() {
		},
	}
}
