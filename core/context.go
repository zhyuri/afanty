package core

import "time"

type Context struct {
	Values    map[string]interface{}
	startTime time.Time
}
