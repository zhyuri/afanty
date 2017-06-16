package core

import (
	"context"
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

func Execute(sm StateMachine) error {
	logrus.Infoln("start execute sm: ", sm.Comment)

	ctx, _ := context.WithTimeout(context.Background(),
		time.Second*time.Duration(sm.TimeoutSeconds))

	for {
		select {
		case <-ctx.Done():
			logrus.Warningln("sm execute timeout: ", sm.Comment)
			return ctx.Err()
		}
	}

	logrus.Infoln("end execute sm: ", sm.Comment)
	return nil
}
