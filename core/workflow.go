package core

import "fmt"

type Workflow struct {
	name string
	// maintain some kind of tree to store the step map
	root    *Step
	stepMap map[string]*Step
}

func init() {
	fmt.Println("Hello World")
}

func NewWorkflow(n string) *Workflow {
	w := &Workflow{
		name:    n,
		stepMap: make(map[string]*Step),
	}

	return w
}

func (w *Workflow) AddStep(step Step) *Workflow {

	return w
}

func (w *Workflow) Run() error {
	return nil
}
