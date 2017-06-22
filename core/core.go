package core

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/Sirupsen/logrus"
	"time"
)

type StateMachine struct {
	Version        string
	Comment        string
	TimeoutSeconds int32
	StartAt        string
	States         map[string]*json.RawMessage
}

func init() {
	logrus.Infoln("init Afanty core")
}

func NewStateMachineFromJSON(j []byte) (StateMachine, error) {
	sm := StateMachine{
		Version:        "1.0",
		Comment:        "Default StateMachine Comment",
		TimeoutSeconds: 10,
		StartAt:        "InitState",
	}
	err := json.Unmarshal(j, &sm)
	if err != nil {
		return StateMachine{}, err
	}

	return sm, nil
}

func (sm *StateMachine) Execute() error {
	logrus.WithField("name", sm.Comment).Infoln("execute sm start")

	ctx, cancel := context.WithTimeout(context.Background(),
		time.Second*time.Duration(sm.TimeoutSeconds))

	errChan := make(chan error)
	go func(sm *StateMachine, done context.CancelFunc) {
		defer done()

		var s State
		var err error
		name := sm.StartAt
		for {
			stateJSON, exist := sm.States[name]
			if !exist {
				err = errors.New("can not find state " + name)
				break
			}
			s, err = ExecuteStateJSON(stateJSON)
			if err != nil || s.End {
				break
			}
			name = s.Next
		}
		if err != nil {
			errChan <- err
		}
	}(sm, cancel)

	select {
	case e := <-errChan:
		logrus.Errorln("StateMachine execute State error", e)
	case <-ctx.Done():
		if e := ctx.Err(); e != context.Canceled {
			logrus.Warningln("sm execute timeout: ", sm.Comment)
		}
	}

	logrus.WithField("name", sm.Comment).Infoln("execute sm end")
	return nil
}

func ExecuteStateJSON(stateJSON *json.RawMessage) (State, error) {
	if stateJSON == nil {
		return State{}, errors.New("input state json is nil")
	}
	stateInterface, err := BuildState(*stateJSON)
	if err != nil {
		return State{}, err
	}
	switch s := stateInterface.(type) {
	case *PassState:
		logrus.Infof("%#v", s)
		return s.State, nil
	case *TaskState:
		logrus.Infof("%#v", s)
		return s.State, nil
	default:
		logrus.WithField("state", s).Error("the logic of state has not been implement")
		return State{}, errors.New("the invoke logic of state has not been implement")
	}
}
