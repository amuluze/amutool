// Package trigger
// Date: 2023/7/19 16:37
// Author: Amu
// Description: 事件触发器
package trigger

import (
	"sync"

	"github.com/pkg/errors"
)

type Trigger struct {
	eventMap sync.Map
}

func NewTrigger() *Trigger {
	return &Trigger{}
}

func (t *Trigger) RegisterEvent(eventName string, callback func(params interface{}) error) error {
	if _, ok := t.eventMap.Load(eventName); ok {
		return errors.New("event already defined")
	}
	t.eventMap.Store(eventName, callback)
	return nil
}

func (t *Trigger) DeleteEvent(eventName string) {
	t.eventMap.Delete(eventName)
}

func (t *Trigger) HasEvent(eventName string) bool {
	if _, ok := t.eventMap.Load(eventName); ok {
		return true
	}
	return false
}

func (t *Trigger) ClearEvent() {
	t.eventMap.Range(func(key, _ interface{}) bool {
		t.eventMap.Delete(key)
		return true
	})
}

func (t *Trigger) Events() (eventNames []string) {
	t.eventMap.Range(func(key, value interface{}) bool {
		eventNames = append(eventNames, key.(string))
		return true
	})
	return
}

func (t *Trigger) CallEvent(eventName string, param interface{}) error {
	event, ok := t.eventMap.Load(eventName)
	if !ok {
		return errors.New("event notfound")
	}
	return event.(func(params interface{}) error)(param)
}
