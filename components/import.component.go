package Component

import (
	"fmt"
	Bus "sprak/bus"
	UI "sprak/ui"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/lipgloss"

	anki "sprak/anki"

	tea "github.com/charmbracelet/bubbletea"
)

type importModel struct {
	importStatus       int // 0 ok, 1 err, nil idle
	totalCards         int
	currentCardsLoaded int
	spinner            spinner.Model
}

type importError struct{}
type importSuccess struct{}

func Import(props *UI.Props) *UI.Component {
	Bus.Publish("log", "Creating Refresh Deck component")

	// Create the spinner to be shown
	s := spinner.NewModel()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	m := importModel{
		totalCards:         0,
		currentCardsLoaded: 0,
		spinner:            s,
	}

	Bus.Log(fmt.Sprintf("%+v\n", m))

	return &UI.Component{
		Init: func() tea.Cmd {
			return func() tea.Msg {
				Bus.Log("starting import...")
				// In here we'll make a call to import the .apkg
				// for now it depends on having a file in the root called deck.apkg
				if _, err := anki.ImportApkg("deck.apkg"); err != nil {
					Bus.Publish("log", err.Error())
					Bus.Log("import error...")
					return importError{}
				} else {
					Bus.Log("import success...")
					return importSuccess{}
				}
			}
		},
		Update: func(msg tea.Msg) tea.Cmd {
			Bus.Log(fmt.Sprintf("%+v", msg))

			switch msg := msg.(type) {

			case Bus.Event:
				switch msg.Topic {
				case "cards:total":
					if count, ok := msg.Data.(int); ok {
						m.totalCards = count
					}
					return spinner.Tick
				case "cards:upserted":
					if _, ok := msg.Data.(int); ok {
						m.currentCardsLoaded += 1
					}
					return spinner.Tick
				}
			case importSuccess:
				m.importStatus = 1
			case importError:
				m.importStatus = 0
			}

			return nil
		},
		View: func() string {
			s := ""

			if m.importStatus == 1 {
				s += fmt.Sprintf("%s Importing %d/%d", m.spinner.View(), m.currentCardsLoaded, m.totalCards)
			} else if m.importStatus == 0 {
				s += "An error occured while importing deck"
			} else {
				s += fmt.Sprintf("%s Waiting to start import...", m.spinner.View())
			}

			return s
		},
		Destroy: func() {
		},
	}
}
