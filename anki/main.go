package Anki

import (
	"fmt"

	Bus "sprak/bus"
	Data "sprak/data"

	"github.com/flimzy/anki"
)

func ImportApkg(path string) (*anki.Apkg, error) {
	fmt.Print("Attempting to load .apkg at path: ", path, "... ")

	apkg, err := anki.ReadFile(path)

	defer func() {
		apkg.Close()
	}()

	if err != nil {
		Bus.Log("Failed to load .apkg file")
		return nil, err
	}

	// Get a count of total cards so we have a point to get to in the loading bar indicator
	var totalCardsCount int
	if err := apkg.Db.QueryRow("select count(*) from notes").Scan(&totalCardsCount); err != nil {
		Bus.Log("Failed to get total count of Notes")
		return nil, err
	}

	Bus.Publish("cards:total", totalCardsCount)

	notes, err := apkg.Notes()
	if err != nil {
		Bus.Log("Failed to get notes")
	}

	var i = 0
	for notes.Next() {
		i++

		note, err := notes.Note()
		if err != nil {
			fmt.Println("Failed to get note", fmt.Errorf(err.Error()))
		}

		if card, err := Data.UpsertCard(note); err != nil {
			return nil, err
		} else {
			Bus.Publish("card:upserted", card.ID)
		}
	}

	Bus.Log("SUCCESS\n")
	return apkg, nil
}
