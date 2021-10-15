package Data

import (
	"context"
	"fmt"
	"sprak/db"

	"github.com/flimzy/anki"
)

var Client = db.NewClient()
var ctx = context.Background()

func Connect() *db.PrismaClient {
	fmt.Print("Connecting to SQLite via Prisma... ")

	if err := Client.Prisma.Connect(); err != nil {
		fmt.Print("FAILURE\n")
		fmt.Errorf(err.Error())
	}

	fmt.Print("SUCCESS\n")

	return Client
}

func Disconnect() {
	fmt.Println("Disconnecting from Prisma...")
	if err := Client.Prisma.Disconnect(); err != nil {
		panic(err)
	}
}

func UpsertCard(note *anki.Note) (*db.CardModel, error) {
	card, err := Client.Card.UpsertOne(
		db.Card.ID.Equals(int(note.ID)),
	).Create(
		db.Card.ID.Set(int(note.ID)),
		db.Card.Mapping.Set("no:en"),
		db.Card.Target.Set(note.FieldValues[0]),
		db.Card.Source.Set(note.FieldValues[1]),
		db.Card.Tags.Set(note.Tags),
	).Update(
		db.Card.Target.Set(note.FieldValues[0]),
		db.Card.Source.Set(note.FieldValues[1]),
		db.Card.Tags.Set(note.Tags),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return card, nil
}
