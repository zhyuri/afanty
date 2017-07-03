package core

import "context"

type Event struct {
	Name    string
	Ctx     context.Context
	retChan chan interface{}
}
type EventHandler func(Event)

func NewEvent(name string, ctx context.Context) *Event {
	return &Event{
		Name:    name,
		Ctx:     ctx,
		retChan: make(chan interface{}),
	}
}

func (e *Event) WaitResult() (interface{}, error) {
	for {
		select {
		case <-e.Ctx.Done():
			return nil, e.Ctx.Err()
		case r := e.retChan:
			return r, nil
		}
	}
}
