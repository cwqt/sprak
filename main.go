package main

import (
	"context"
	"fmt"
	"os"

	"sprak/db"

	tea "github.com/charmbracelet/bubbletea"
	lipgloss "github.com/charmbracelet/lipgloss"
	anki "github.com/flimzy/anki"

	ViewPage "sprak/view-page"
)

type model struct {
	choices  []string         // items on the to-do list
	cursor   int              // which to-do list item our cursor is pointing at
	selected map[int]struct{} // which to-do items are selected

	currentView ViewPage.ViewPage
}

func initialModel() model {
	return model{
		// Our shopping list is a grocery list
		choices: []string{"Buy carrots", "Buy celery", "Buy kohlrabi"},

		// A map which indicates which choices are selected. We're using
		// the  map like a mathematical set. The keys refer to the indexes
		// of the `choices` slice, above.
		selected: make(map[int]struct{}),

		currentView: ViewPage.Home,
	}
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		case "1":
			m.currentView = ViewPage.Main

		case "2":
			m.currentView = ViewPage.Home

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() string {
	// The header
	s := "Hello!\n"

	var style = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("53"))

	if m.currentView == ViewPage.Main {
		// Iterate over our choices
		for i, choice := range m.choices {

			// Is the cursor pointing at this choice?
			cursor := " " // no cursor
			if m.cursor == i {
				cursor = ">" // cursor!
			}

			// Is this choice selected?
			checked := " " // not selected
			if _, ok := m.selected[i]; ok {
				checked = "x" // selected!

				// Render the row
				s += style.Render(fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)) + "\n"
			} else {
				// Render the row
				s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
			}
		}
	}

	if m.currentView == ViewPage.Home {
		s += "Welcome home! 💚 "
	}

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}

func main() {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		fmt.Errorf(err.Error())
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	ctx := context.Background()

	createCard := func(note *anki.Note) (*db.CardModel, error) {
		card, err := client.Card.CreateOne(
			db.Card.ID.Set(int(note.ID)),
			db.Card.Mapping.Set("no:en"),
			db.Card.Target.Set(note.FieldValues[0]),
			db.Card.Source.Set(note.FieldValues[1]),
			db.Card.Tags.Set(""),
		).Exec(ctx)

		if err != nil {
			return nil, err
		}

		return card, nil
	}

	apkg, err := anki.ReadFile("Bokmål.apkg")
	if err != nil {
		fmt.Println("Failed to load .apkg file")
		os.Exit(1)
	}

	notes, _ := apkg.Notes()

	for notes.Next() {
		note, _ := notes.Note()
		if err != nil {
			fmt.Println("Failed on note")
		}

		createdNote, _ := createCard(note)
		fmt.Printf("createdNote: %v\n", createdNote)
	}

	apkg.Close()

	// p := tea.NewProgram(initialModel())
	// if err := p.Start(); err != nil {
	// 	fmt.Printf("Alas, there's been an error: %v", err)
	// 	os.Exit(1)
	// }
}
