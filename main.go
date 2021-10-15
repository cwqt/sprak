package main

import (
	"fmt"
	Anki "sprak/anki"
	Bus "sprak/bus"
	Data "sprak/data"
)

func main() {
	Data.Connect()

	Bus.Subscribe("card:upserted", func(cardId int) {
		fmt.Printf("Upserted card: %d!\n", cardId)
	})

	Anki.ImportApkg("Bokmål.apkg")

}
