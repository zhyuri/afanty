package core

type BaseState struct{}

type State struct {
	BaseState
	Type       string `json:"type,omitempty"`
	Next       string `json:"next,omitempty"`
	End        bool   `json:"end,omitempty"`
	Comment    string `json:"comment,omitempty"`
	InputPath  string `json:"inputPath,omitempty"`
	OutputPath string `json:"outputPath,omitempty"`
}

type PassState struct {
	State
	Result     string `json:"result,omitempty"`
	ResultPath string `json:"resultPath,omitempty"`
}

type TaskState struct {
	State
	Resource         string     `json:"resource,omitempty"`
	ResultPath       string     `json:"resultPath,omitempty"`
	Retry            []*Retry   `json:"retry,omitempty"`
	Catch            []*Catcher `json:"catch,omitempty"`
	TimeoutSeconds   int32      `json:"timeoutSeconds,omitempty"`
	HeartbeatSeconds int32      `json:"heartbeatSeconds,omitempty"`
}

type ChoiceRule struct {
	Variable string `json:"variable,omitempty"`
	Type     string `json:"type,omitempty"`
	// String, number, boolean, Timestamp or non-empty Array of ChoiceRule
	// which means ChoiceRule can be nested
	Target interface{} `json:"target,omitempty"`
	Next   string      `json:"next,omitempty"`
}

type ChoiceState struct {
	// Choice states do not support the End field
	State
	Choices []*ChoiceRule `json:"choices,omitempty"`
	Default string        `json:"default,omitempty"`
}

type WaitState struct {
	State
	Seconds       int32  `json:"seconds,omitempty"`
	Timestamp     string `json:"timestamp,omitempty"`
	SecondsPath   string `json:"secondsPath,omitempty"`
	TimestampPath string `json:"timestampPath,omitempty"`
}

type Branches struct {
	StartAt string               `json:"startAt,omitempty"`
	States  map[string]BaseState `json:"states,omitempty"`
}

type ParallelState struct {
	State
	Branches   []*Branches `json:"branches,omitempty"`
	ResultPath string      `json:"resultPath,omitempty"`
	Retry      []*Retry    `json:"retry,omitempty"`
	Catch      []*Catcher  `json:"catch,omitempty"`
}

type SucceedState struct {
	BaseState
	Type    string `json:"type,omitempty"`
	Comment string `json:"comment,omitempty"`
}

type FailState struct {
	BaseState
	Type    string `json:"type,omitempty"`
	Comment string `json:"comment,omitempty"`
	Cause   string `json:"cause,omitempty"`
	Error   string `json:"error,omitempty"`
}
