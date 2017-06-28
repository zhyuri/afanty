package core

import (
	"context"
	"github.com/Sirupsen/logrus"
	. "github.com/zhyuri/afanty/api"
	"plugin"
	"time"
)

func getPlugin(name string) (Run, error) {
	p, err := plugin.Open(name)
	if err != nil {
		return nil, err
	}
	symbolRun, err := p.Lookup("Run")
	if err != nil {
		return nil, err
	}
	run, ok := symbolRun.(Run)
	if !ok {

	}
	return run, nil
}

func (t *TaskState) Call(data *[]byte) (State, error) {
	logrus.Info("")

	run, err := getPlugin(t.Resource)
	if err != nil {
		logrus.Errorln("")
		return State{}, err
	}
	respChan := make(chan MOutput)
	errChan := make(chan StateError)
	ctx, _ := context.WithTimeout(context.Background(), time.Second*time.Duration(t.TimeoutSeconds))
	go func() {
		resp, stateErr := run(MInput{})
		respChan <- resp
		if stateErr.Err != nil {
			errChan <- stateErr
		}
	}()

	select {
	case <-ctx.Done():
		if e := ctx.Err(); e == context.DeadlineExceeded {
			logrus.Warningln("taskstate execute timeout: ", t.Comment, e)
			return t.State, e
		} else if e != context.Canceled {
			logrus.Errorln("taskstate execute error: ", t.Comment, e)
			return t.State, e
		}
	case resp := <-respChan:
		output, err := Process(resp.Output, "")
		if err == nil {
			data = &output
		} else {
			logrus.Warnln("")
		}
	}
	return t.State, nil
}

func (t *TaskState) DoRetry(name string) error {

	return nil
}

func (t *TaskState) DoCatch() error {

	return nil
}
