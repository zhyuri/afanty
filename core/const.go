package core

import (
	"encoding/json"
	"reflect"
)

const (
	Name_PassState     = "Pass"
	Name_TaskState     = "Task"
	Name_ChoiceState   = "Choice"
	Name_WaitState     = "Wait"
	Name_ParallelState = "Parallel"
	Name_SucceedState  = "Succeed"
	Name_FailState     = "Fail"

	Errors_All           = "States.ALL"
	Errors_Timeout       = "States.Timeout"
	Errors_Failed        = "States.TaskFailed"
	Errors_Permissions   = "States.Permissions"
	Errors_ProcessFailed = "States.ProcessFailed"
)

var (
	stateType = make(map[string]reflect.Type)
)

func init() {
	stateType[Name_PassState] = reflect.TypeOf(PassState{})
	stateType[Name_TaskState] = reflect.TypeOf(TaskState{})
	stateType[Name_ChoiceState] = reflect.TypeOf(ChoiceState{})
	stateType[Name_WaitState] = reflect.TypeOf(WaitState{})
	stateType[Name_ParallelState] = reflect.TypeOf(ParallelState{})
	stateType[Name_SucceedState] = reflect.TypeOf(SucceedState{})
	stateType[Name_FailState] = reflect.TypeOf(FailState{})
}

type (
	BaseState struct {
		Type string
	}

	State struct {
		BaseState
		Next       string
		End        bool
		Comment    string
		InputPath  string
		OutputPath string
	}

	PassState struct {
		State
		Result     *json.RawMessage
		ResultPath string
	}

	TaskState struct {
		State
		Resource         string
		ResultPath       string
		Retry            []*Retry
		Catch            []*Catcher
		TimeoutSeconds   int32
		HeartbeatSeconds int32
	}

	ChoiceRule struct {
		Variable string
		Type     string
		// String, number, boolean, Timestamp or non-empty Array of ChoiceRule
		// which means ChoiceRule can be nested
		Target interface{}
		Next   string
	}

	ChoiceState struct {
		// Choice states do not support the End field
		State
		Choices []*ChoiceRule
		Default string
	}

	WaitState struct {
		State
		Seconds       int32
		Timestamp     string
		SecondsPath   string
		TimestampPath string
	}

	Branches struct {
		StartAt string
		States  map[string]*json.RawMessage
	}

	ParallelState struct {
		State
		Branches   []*Branches
		ResultPath string
		Retry      []*Retry
		Catch      []*Catcher
	}

	SucceedState struct {
		BaseState
		Comment string
	}

	FailState struct {
		BaseState
		Comment string
		Cause   string
		Error   string
	}
)
