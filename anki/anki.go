package Anki

import (
	"fmt"

	Bus "sprak/bus"
	Data "sprak/data"

	"github.com/flimzy/anki"
)

func ImportApkg(path string) (*anki.Apkg, error) {
	fmt.Println("Attempting to load .apkg at path:", path)

	apkg, err := anki.ReadFile(path)
	if err != nil {
		fmt.Println("Failed to load .apkg file")
	}

	notes, _ := apkg.Notes()

	var i = 0
	for notes.Next() {
		i++

		note, err := notes.Note()
		if err != nil {
			fmt.Println("Failed on note", note.ID, fmt.Errorf(err.Error()))
		}

		card, err := Data.UpsertCard(note)
		if err != nil {
			fmt.Print(err)
		}

		Bus.Publish("card:upserted", card.ID)
	}

	apkg.Close()
	return apkg, nil
}
