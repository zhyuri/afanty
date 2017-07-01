package core

import (
	"context"
	"errors"
	"plugin"
	"time"

	"github.com/Sirupsen/logrus"
	. "github.com/zhyuri/afanty/api"
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

func (t *TaskState) Call(data *[]byte) (State, *StateError) {
	logEntry := logrus.WithField("name", t.Comment)
	logEntry.Info("start execute")

	run, err := getPlugin(t.Resource)
	if err != nil {
		logEntry.Errorln("")
		return State{}, &StateError{Name: Errors_Failed, Err: err}
	}
	type Result struct {
		Resp     MOutput
		StateErr StateError
	}
	resultChan := make(chan Result)
	ctx, _ := context.WithTimeout(context.Background(), time.Second*time.Duration(t.TimeoutSeconds))
	go func() {
		resp, stateErr := run(MInput{Input: *data})
		select {
		case <-ctx.Done():
			resultChan <- Result{
				resp,
				StateError{Name: Errors_Timeout, Err: ctx.Err()},
			}
			return
		default:
			resultChan <- Result{resp, stateErr}
		}
	}()

	for {
		select {
		case <-ctx.Done():
			switch err := ctx.Err(); err {
			case context.Canceled:
				// handle timeout
				logEntry.Warningln("taskstate execute timeout: ", err)
				return t.State, &StateError{Name: Errors_Timeout, Err: err}
			default:
				// other unknown context error
				logEntry.Errorln("taskstate execute error: ", err)
				return t.State, &StateError{Name: Errors_Failed, Err: err}
			}
		case r := <-resultChan:
			if r.StateErr.Err != nil {
				// client defined error
				return State{}, &r.StateErr
			}
			output, err := Process(r.Resp.Output, "")
			if err != nil {
				logEntry.Warnln("process output failed", err)
				return State{}, &StateError{Name: Errors_ProcessFailed, Err: err}
			}
			data = &output
			return t.State, nil
		}
	}

}

type RetryContext struct {
	StateErr  *StateError
	Times     int32
	StartTime time.Time
	Interval  time.Duration
	Fallback  bool
}

func (t *TaskState) RetryWait(ctx *RetryContext) (*RetryContext, bool) {
	for _, retry := range t.Retry {
		for _, errorEqual := range retry.ErrorEquals {
			if errorEqual == ctx.StateErr.Name || errorEqual == Errors_All {
				if retry.MaxAttempts <= 0 {
					continue
				}
				if ctx.Times >= retry.MaxAttempts {
					return ctx, true
				}
				sleepTime := (retry.IntervalSeconds + float32(ctx.Times)*retry.BackoffRate) * float32(time.Second)
				time.Sleep(time.Duration(sleepTime))

				ctx.Times++
				return ctx, false
			}
		}
	}
	return ctx, true
}

func (t *TaskState) CatchFail(err *StateError, data *[]byte) (State, error) {
	for _, catch := range t.Catch {
		for _, catchEqual := range catch.ErrorEquals {
			if catchEqual == err.Name || catchEqual == Errors_All {
				output, err := Process(*data, catch.ResultPath)
				if err != nil {
					return State{}, StateError{Name: Errors_ProcessFailed, Err: err}
				}
				data = &output
				return State{Next: catch.Next}, nil
			}
		}
	}
	return State{}, errors.New("cannot find matched catcher for task: " + t.Comment + " err: " + err.Error())
}
