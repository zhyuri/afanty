package core

func NewPassState() PassState {
	return PassState{State: State{Type: "Pass"}}
}

func NewTaskState() TaskState {
	return TaskState{State: State{Type: "Task"}}
}

func NewChoiceState() ChoiceState {
	return ChoiceState{State: State{Type: "Choice"}}
}

func NewWaitState() WaitState {
	return WaitState{State: State{Type: "Wait"}}
}

func NewParallelState() ParallelState {
	return ParallelState{State: State{Type: "Parallel"}}
}

func NewSucceedState() SucceedState {
	return SucceedState{Type: "Succeed"}
}

func NewFailState() FailState {
	return FailState{Type: "Fail"}
}
