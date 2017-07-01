package core

import "context"

type Event struct {
	Name    string
	Context context.Context
	retChan chan context.Context
}
type EventHandler func(Event)
