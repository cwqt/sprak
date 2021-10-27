package Component

import (
	"fmt"
	anki "sprak/anki"
	Bus "sprak/bus"
	UI "sprak/ui"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/lipgloss"

	tea "github.com/charmbracelet/bubbletea"
)

type importModel struct {
	importStatus       int // 0 importing, -1 err, 1 done
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
		importStatus:       0,
		currentCardsLoaded: 0,
		spinner:            s,
	}

	return &UI.Component{
		Model: m,
		Init: func() tea.Cmd {
			startImport := func() tea.Msg {
				Bus.Log("starting import...")
				if _, err := anki.ImportApkg("/Users/cass/code/sprak/deck.apkg"); err != nil {
					Bus.Err(err.Error())
					return importError{}
				} else {
					Bus.Log("import success...")
					return importSuccess{}
				}
			}

			return UI.Cmds(startImport, spinner.Tick).AsCmd()
		},
		Update: func(msg tea.Msg) tea.Cmd {
			cmds := UI.Cmds()

			switch msg := msg.(type) {
			case Bus.Event:
				switch msg.Topic {
				case "cards:upserted":
					if cardIds, ok := msg.Data.([]int); ok {
						m.currentCardsLoaded += len(cardIds)
					}
					if m.currentCardsLoaded%250 == 0 {
						cmds.Append(spinner.Tick)
					}
				}
			case importSuccess:
				m.importStatus = 1
			case importError:
				m.importStatus = -1
			case spinner.TickMsg:
				// keep spinning & emitting events until finished or errored
				if m.importStatus == 0 {
					var cmd tea.Cmd
					m.spinner, cmd = m.spinner.Update(msg)
					cmds.Append(cmd)
				}
			}

			return cmds.AsCmd()
		},
		View: func(width int) string {
			s := ""

			if m.importStatus == 0 {
				s += fmt.Sprintf("%s Importing %d cards", m.spinner.View(), m.currentCardsLoaded)
			} else if m.importStatus == 1 {
				s += fmt.Sprintf("Imported %d cards!", m.currentCardsLoaded)
			} else if m.importStatus == -1 {
				s += "Failed to import cards"
			}

			return s
		},
		Destroy: func() {
		},
	}
}
