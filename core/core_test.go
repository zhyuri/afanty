package core

import (
	"testing"
)

var exampleStateMachineJSON = []byte(`{
  "Comment": "An example that adds two numbers.",
  "StartAt": "Add",
  "Version": "1.0",
  "TimeoutSeconds": 10,
  "States":
    {
        "Add": {
          "Type": "Task",
          "Resource": "arn:aws:lambda:us-east-1:123456789012:function:Add",
          "End": true
        }
    }
}`)

func TestNewStateMachineFromJSON(t *testing.T) {
	got, err := NewStateMachineFromJSON(exampleStateMachineJSON)
	if err != nil {
		t.Errorf("NewStateMachineFromJSON return err, %#v", err)
	}
	if _, exist := got.States[got.StartAt]; !exist {
		t.Errorf("can not find states on start, need %v, found %#v", got.StartAt, got.States)
	}
}

func TestExecute(t *testing.T) {
	sm, err := NewStateMachineFromJSON(exampleStateMachineJSON)
	if err != nil {
		t.Errorf("NewStateMachineFromJSON return err, %#v", err)
	}
	if err = sm.Execute(); err != nil {
		t.Errorf("Execute return err, %#v", err)
	}
}
