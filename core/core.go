package core

import (
	"context"
	"errors"
	"github.com/Sirupsen/logrus"
	"time"
)

type StateMachine struct {
	Version        string               `json:"version,omitempty"`
	Comment        string               `json:"comment,omitempty"`
	TimeoutSeconds int32                `json:"timeoutSeconds,omitempty"`
	StartAt        string               `json:"startAt,omitempty"`
	States         map[string]BaseState `json:"states,omitempty"  `
}

func init() {
	logrus.Infoln("init Afanty core")
}

func NewStateMachine() StateMachine {
	return StateMachine{
		Version:        "1.0",
		Comment:        "Default StateMachine Comment",
		TimeoutSeconds: 10,
	}
}

func (sm StateMachine) Execute() error {
	logrus.Infoln("start execute sm: ", sm.Comment)

	mCtx, _ := context.WithTimeout(context.Background(),
		time.Second*time.Duration(sm.TimeoutSeconds))

	var (
		baseState BaseState
		//state     State
		exist  bool
		isDone chan bool
	)
	if baseState, exist = sm.States[sm.StartAt]; !exist {
		logrus.Warnln()
		return errors.New("cann't find state " + sm.StartAt)
	}
	logrus.Debug(baseState)

	isDone = make(chan bool, 1)

	//go func(state State, isDone chan bool) {
	//	for s := state; ; {
	//
	//	}
	//
	//}(state, isDone)

	for {
		select {
		case <-isDone:
			break
		case <-mCtx.Done():
			logrus.Warningln("sm execute timeout: ", sm.Comment)
			return mCtx.Err()
		}
	}

	logrus.Infoln("end execute sm: ", sm.Comment)
	return nil
}
