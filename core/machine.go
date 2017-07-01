package core

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/Sirupsen/logrus"
)

type StateMachine struct {
	Version        string
	Comment        string
	TimeoutSeconds int32
	StartAt        string
	States         map[string]*json.RawMessage
}

func NewStateMachineFromJSON(j []byte) (StateMachine, error) {
	sm := StateMachine{
		Version:        "1.0",
		Comment:        "Default StateMachine Comment",
		TimeoutSeconds: 10,
		StartAt:        "InitState",
	}
	err := json.Unmarshal(j, &sm)
	return sm, err
}

func (sm *StateMachine) Execute(data *json.RawMessage) error {
	start := time.Now()
	logEntry := logrus.WithField("name", sm.Comment)
	logEntry.Infoln("execute sm start")
	defer logEntry.Infoln("cost", strconv.FormatFloat(time.Since(start).Seconds(), 'f', 3, 64))

	ctx, done := context.WithTimeout(context.Background(),
		time.Second*time.Duration(sm.TimeoutSeconds))

	errChan := make(chan error)
	go func() {
		defer done()

		var (
			s   State
			err error
		)
		name := sm.StartAt
		for {
			stateJSON, exist := sm.States[name]
			if !exist {
				err = errors.New("can not find state " + name)
				break
			}
			s, err = ExecuteStateJSON(*stateJSON, data)
			if err != nil || s.End {
				break
			}
			select {
			case <-ctx.Done():
				return
			}
		}
		errChan <- err
	}()

	select {
	case e := <-errChan:
		if e != nil {
			logrus.Warning("StateMachine execute State error, %v", e)
		}
		return e
	case <-ctx.Done():
		if e := ctx.Err(); e != context.Canceled {
			logrus.Warning("sm execute timeout: ", sm.Comment)
		}
	}

	return nil
}

func ExecuteStateJSON(stateJSON json.RawMessage, data *json.RawMessage) (State, error) {
	if stateJSON == nil {
		return State{}, errors.New("input state json is nil")
	}
	state, err := BuildState(stateJSON)
	if err != nil {
		return State{}, err
	}
	switch s := state.(type) {
	case *TaskState:
		retryContext := &RetryContext{}
		fallback := false
		for {
			state, err := s.Call((*[]byte)(data))
			if err == nil {
				return state, nil
			}
			retryContext.StateErr = err
			// client defined error
			if retryContext, fallback = s.RetryWait(retryContext); fallback {
				// catch
				return s.CatchFail(err, (*[]byte)(data))
			}
		}
	case *SucceedState:
		logrus.Infof("%#v", s)
		return State{}, nil
	case *FailState:
		logrus.Infof("%#v", s)
		return State{}, nil
	default:
		logrus.WithField("state", s).Errorln("the logic of state has not been implement")
		return State{}, errors.New("the invoke logic of state has not been implement")
	}
}
