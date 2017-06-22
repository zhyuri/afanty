package core

import "encoding/json"

type BaseState struct {
	Type string
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
	Type    string
	Comment string
}

type FailState struct {
	BaseState
	Type    string
	Comment string
	Cause   string
	Error   string
}
