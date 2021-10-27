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
		Bus.Err("Failed to read file")
		Bus.Err(err.Error())
	}

	defer func() {
		apkg.Close()
	}()

	if err != nil {
		Bus.Err("Failed to load .apkg file")
		return nil, err
	}

	notes, err := apkg.Notes()
	if err != nil {
		Bus.Err("Failed to get notes")
	}

	i := 0
	// emit an event every CHUNK_SIZE, prevents bus from being overwhelmed
	CHUNK_SIZE := 250
	chunk := make([]int, CHUNK_SIZE)
	for notes.Next() {
		i++

		note, err := notes.Note()
		if err != nil {
			Bus.Err("Failed to get note")
			fmt.Println("Failed to get note", fmt.Errorf(err.Error()))
		}

		card, err := Data.UpsertCard(note)
		if err != nil {
			return nil, err
		}

		if i%CHUNK_SIZE == 0 {
			Bus.Publish("cards:upserted", chunk)
		} else {
			chunk[i%CHUNK_SIZE] = card.ID
		}
	}

	Bus.Log("SUCCESS\n")
	return apkg, nil
}
