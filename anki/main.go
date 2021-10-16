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
	if err != nil {
		fmt.Println("Failed to load .apkg file")
		return nil, err
	}

	// Get a count of total cards so we have a point to get to in the loading bar indicator
	var totalCardsCount int
	if err := apkg.Db.QueryRow("select count(*) from notes").Scan(&totalCardsCount); err != nil {
		fmt.Println("Failed to get total count of Notes")
		return nil, err
	}

	Bus.Publish("cards:total", totalCardsCount)

	notes, _ := apkg.Notes()

	var i = 0
	for notes.Next() {
		i++

		note, err := notes.Note()
		if err != nil {
			fmt.Println("Failed to get note", fmt.Errorf(err.Error()))
		}

		card, err := Data.UpsertCard(note)
		if err != nil {
			return nil, err
		}

		Bus.Publish("card:upserted", card.ID)
	}

	apkg.Close()
	fmt.Print("SUCCESS\n")
	return apkg, nil
}
