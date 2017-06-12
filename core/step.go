package core

type stepFunc func(Context) error

type Step struct {
	name    string
	execute stepFunc
	headCon Condition
	tailCon Condition

	prev *Step
	next *Step
}

type option func(*Step)

func NewStep(exe stepFunc, opts ...option) *Step {
	s := &Step{
		name:    "default",
		execute: exe,
	}

	for _, opt := range opts {
		opt(s)
	}
	return s
}

func WithHeadCondition(c Condition) option {
	return func(s *Step) { s.headCon = c }
}
func WithTailCondition(c Condition) option {
	return func(s *Step) { s.tailCon = c }
}
