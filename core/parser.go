package core

import (
	"encoding/json"
	"errors"
	"github.com/Sirupsen/logrus"
	"reflect"
)

const (
	NamePassState     = "Pass"
	NameTaskState     = "Task"
	NameChoiceState   = "Choice"
	NameWaitState     = "Wait"
	NameParallelState = "Parallel"
	NameSucceedState  = "Succeed"
	NameFailState     = "Fail"
)

var (
	stateType = make(map[string]reflect.Type)
)

func init() {
	stateType[NamePassState] = reflect.TypeOf(PassState{})
	stateType[NameTaskState] = reflect.TypeOf(TaskState{})
	stateType[NameChoiceState] = reflect.TypeOf(ChoiceState{})
	stateType[NameWaitState] = reflect.TypeOf(WaitState{})
	stateType[NameParallelState] = reflect.TypeOf(ParallelState{})
	stateType[NameSucceedState] = reflect.TypeOf(SucceedState{})
	stateType[NameFailState] = reflect.TypeOf(FailState{})
}

func BuildState(data []byte) (interface{}, error) {
	var (
		t   reflect.Type
		err error
	)
	if t, err = ParseStateType(data); err != nil {
		return nil, err
	}
	ret := reflect.New(t).Interface()
	err = json.Unmarshal(data, &ret)
	return ret, err
}

func ParseStateType(data []byte) (reflect.Type, error) {
	var (
		t        reflect.Type
		stateMap map[string]*json.RawMessage
		msg      *json.RawMessage
		typeStr  string
		ok       bool
		err      error
	)
	err = json.Unmarshal(data, &stateMap)
	if msg, ok = stateMap["Type"]; !ok {
		logrus.WithFields(logrus.Fields{
			"json":     data,
			"stateMap": stateMap,
		}).Errorln("cannot find type field from json")
		return nil, err
	}
	err = json.Unmarshal(*msg, &typeStr)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"json":     data,
			"stateMap": stateMap,
			"err":      err,
		}).Errorln("cannot unmarshal state from json")
		return nil, err
	}
	if t, ok = stateType[typeStr]; !ok {
		logrus.WithFields(logrus.Fields{
			"stateMap": stateMap,
			"typeStr":  typeStr,
		}).Errorln("unknown state type")
		return nil, errors.New("unknown state type")
	}
	return t, nil
}
