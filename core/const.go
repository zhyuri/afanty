package core

import "encoding/json"

type BaseState struct {
	Type string `json:"omitempty"`
}

type State struct {
	BaseState
	Next       string `json:"omitempty"`
	End        bool   `json:"omitempty"`
	Comment    string `json:"omitempty"`
	InputPath  string `json:"omitempty"`
	OutputPath string `json:"omitempty"`
}

type PassState struct {
	State
	Result     *json.RawMessage `json:"omitempty"`
	ResultPath string           `json:"omitempty"`
}

type TaskState struct {
	State
	Resource         string     `json:"omitempty"`
	ResultPath       string     `json:"omitempty"`
	Retry            []*Retry   `json:"omitempty"`
	Catch            []*Catcher `json:"omitempty"`
	TimeoutSeconds   int32      `json:"omitempty"`
	HeartbeatSeconds int32      `json:"omitempty"`
}

type ChoiceRule struct {
	Variable string `json:"omitempty"`
	Type     string `json:"omitempty"`
	// String, number, boolean, Timestamp or non-empty Array of ChoiceRule
	// which means ChoiceRule can be nested
	Target interface{} `json:"omitempty"`
	Next   string      `json:"omitempty"`
}

type ChoiceState struct {
	// Choice states do not support the End field
	State
	Choices []*ChoiceRule `json:"omitempty"`
	Default string        `json:"omitempty"`
}

type WaitState struct {
	State
	Seconds       int32  `json:"omitempty"`
	Timestamp     string `json:"omitempty"`
	SecondsPath   string `json:"omitempty"`
	TimestampPath string `json:"omitempty"`
}

type Branches struct {
	StartAt string               `json:"omitempty"`
	States  map[string]BaseState `json:"omitempty"`
}

type ParallelState struct {
	State
	Branches   []*Branches `json:"omitempty"`
	ResultPath string      `json:"omitempty"`
	Retry      []*Retry    `json:"omitempty"`
	Catch      []*Catcher  `json:"omitempty"`
}

type SucceedState struct {
	BaseState
	Type    string `json:"omitempty"`
	Comment string `json:"omitempty"`
}

type FailState struct {
	BaseState
	Type    string `json:"omitempty"`
	Comment string `json:"omitempty"`
	Cause   string `json:"omitempty"`
	Error   string `json:"omitempty"`
}
