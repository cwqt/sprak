package Exercise

import "sprak/db"

type Type int

const (
	TranslateSentence Type = iota
)

type Current struct {
	Type Type
	Card *db.CardModel
}
