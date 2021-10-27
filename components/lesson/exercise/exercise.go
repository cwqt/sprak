package Exercise

// Current state of the exercise
type ExerciseState int

const (
	Idle ExerciseState = iota
	Failed
	Completed
)

type ExerciseType int

const (
	TranslateSentence ExerciseType = iota
	Listening
	MultipleChoice
)

type Exercise struct {
	State ExerciseState
	Type  ExerciseType
	Data  interface{}
}
