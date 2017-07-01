package core

import (
	"errors"
	"time"

	"github.com/Sirupsen/logrus"
)

type AfantyCore struct {
	Name string

	eventBus chan interface{}
	eventMap map[string][]EventHandler
	stopChan chan bool
	log      *logrus.Entry
}

type Option func(*AfantyCore)

func init() {
	logrus.Infoln("init Afanty core")
}

func NewAfantyInstance(name string, options ...Option) *AfantyCore {
	events := make(map[string][]EventHandler)
	entry := logrus.WithFields(logrus.Fields{
		"name":  name,
		"start": time.Now().String(),
	})
	ac := &AfantyCore{
		Name: name,

		eventBus: make(chan interface{}, 0),
		eventMap: events,
		stopChan: make(chan bool),
		log:      entry,
	}
	for _, option := range options {
		option(ac)
	}
	return ac
}

func (c *AfantyCore) Shutdown() {
	c.stopChan <- true
}

func (c *AfantyCore) Pub(e Event) error {
	handlers, ok := c.eventMap[e.Name]
	if !ok {
		return errors.New("event [" + e.Name + "] has not been registered")
	}
	for _, handler := range handlers {
		go handler(e)
	}
	return nil
}

func (c *AfantyCore) Sub(name string, handler EventHandler) error {
	handlers, exist := c.eventMap[name]
	if !exist {
		return errors.New("cannot subscribe to a event not exist, " + name)
	}
	handlers = append(handlers, handler)
	return nil
}
