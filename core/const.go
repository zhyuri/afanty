package core

import (
	"encoding/json"
	"reflect"
)

const (
	NamePassState     = "Pass"
	NameTaskState     = "Task"
	NameChoiceState   = "Choice"
	NameWaitState     = "Wait"
	NameParallelState = "Parallel"
	NameSucceedState  = "Succeed"
	NameFailState     = "Fail"
)

var (
	stateType = make(map[string]reflect.Type)
)

func init() {
	stateType[NamePassState] = reflect.TypeOf(PassState{})
	stateType[NameTaskState] = reflect.TypeOf(TaskState{})
	stateType[NameChoiceState] = reflect.TypeOf(ChoiceState{})
	stateType[NameWaitState] = reflect.TypeOf(WaitState{})
	stateType[NameParallelState] = reflect.TypeOf(ParallelState{})
	stateType[NameSucceedState] = reflect.TypeOf(SucceedState{})
	stateType[NameFailState] = reflect.TypeOf(FailState{})
}

type RunnableState interface {
	Call(data *json.RawMessage) *json.RawMessage
}

type BaseState struct {
	Type string
	Data *json.RawMessage
}

type State struct {
	BaseState
	Next       string
	End        bool
	Comment    string
	InputPath  string
	OutputPath string
}

type PassState struct {
	State
	Result     *json.RawMessage
	ResultPath string
}

type TaskState struct {
	State
	Resource         string
	ResultPath       string
	Retry            []*Retry
	Catch            []*Catcher
	TimeoutSeconds   int32
	HeartbeatSeconds int32
}

type ChoiceRule struct {
	Variable string
	Type     string
	// String, number, boolean, Timestamp or non-empty Array of ChoiceRule
	// which means ChoiceRule can be nested
	Target interface{}
	Next   string
}

type ChoiceState struct {
	// Choice states do not support the End field
	State
	Choices []*ChoiceRule
	Default string
}

type WaitState struct {
	State
	Seconds       int32
	Timestamp     string
	SecondsPath   string
	TimestampPath string
}

type Branches struct {
	StartAt string
	States  map[string]*json.RawMessage
}

type ParallelState struct {
	State
	Branches   []*Branches
	ResultPath string
	Retry      []*Retry
	Catch      []*Catcher
}

type SucceedState struct {
	BaseState
	Comment string
}

type FailState struct {
	BaseState
	Comment string
	Cause   string
	Error   string
}
