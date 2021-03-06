package core

import (
	"encoding/json"
	"errors"
	"github.com/Sirupsen/logrus"
	"reflect"
)

func BuildState(data []byte) (interface{}, error) {
	t, err := ParseStateType(data)
	if err != nil {
		return nil, err
	}
	ret := reflect.New(t).Interface()
	err = json.Unmarshal(data, &ret)
	if err != nil {
		logrus.WithField("err", err).Errorln("json unmarshal failed")
		return nil, err
	}
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
	logEntry := logrus.WithFields(logrus.Fields{
		"json":     data,
		"stateMap": stateMap,
	})
	if msg, ok = stateMap["Type"]; !ok {
		logEntry.Errorln("cannot find type field from json")
		return nil, err
	}
	err = json.Unmarshal(*msg, &typeStr)
	if err != nil {
		logEntry.WithField("err", err).Errorln("cannot unmarshal state from json")
		return nil, err
	}
	if t, ok = stateType[typeStr]; !ok {
		logEntry.WithField("typeStr", typeStr).Errorln("unknown state type")
		return nil, errors.New("unknown state type")
	}
	return t, nil
}
