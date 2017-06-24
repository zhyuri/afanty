package core

import (
	"encoding/json"
	"github.com/Sirupsen/logrus"
)

func (t *TaskState) Call(data *json.RawMessage) error {
	logrus.Info("")
	return nil
}
