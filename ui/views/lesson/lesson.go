package Lesson

import (
	"sprak/db"

	Exercise "sprak/ui/views/lesson/exercises"
)

type State int

// Current state of the lesson
const (
	Idle State = iota
	Failure
	Success
)

type model struct {
	state    State
	cards    *[]db.CardModel
	exercise Exercise.Current
}
