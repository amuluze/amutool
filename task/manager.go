// Package task
// Date: 2023/7/19 15:30
// Author: Amu
// Description:
package task

import (
	"context"
	"sync"

	"github.com/pkg/errors"
)

type Manager struct {
	tasks     sync.Map
	taskChan  chan *Task
	Ctx       context.Context
	cancel    context.CancelFunc
	runStatus *AtomicStatus
}

func NewManager(ctx context.Context, m *Manager) {
	m.tasks = sync.Map{}
	m.runStatus = &AtomicStatus{}
	m.Ctx, m.cancel = context.WithCancel(ctx)
}

func (m *Manager) Register(task *Task) error {
	if _, ok := m.tasks.Load(task.name); ok {
		return errors.New("has same register task name")
	}

	m.tasks.Store(task.name, task)
	if m.IsRun() {
		m.taskChan <- task
	}
	return nil
}

func (m *Manager) Run() {
	m.tasks.Range(func(key, value interface{}) bool {
		go value.(*Task).run()

		return true
	})

	m.taskChan = make(chan *Task)
	m.runStatus.SetRunStatus()
breakLoop:
	for {
		select {
		case task := <-m.taskChan:
			go task.run()
		case <-m.Ctx.Done():
			break breakLoop
		}
	}

	<-m.Ctx.Done()
	m.close()
}

func (m *Manager) IsRun() bool {
	return m.runStatus.IsRun()
}

func (m *Manager) Cancel(taskName string) {
	task, ok := m.tasks.Load(taskName)
	if !ok {
		return
	}

	m.tasks.Delete(taskName)
	task.(*Task).stop()
}

func (m *Manager) close() {
	m.tasks.Range(func(key, value interface{}) bool {
		go value.(*Task).stop()
		return true
	})

	if m.taskChan != nil {
		for {
			select {
			case <-m.taskChan:
			default:
				m.runStatus.SetStopStatus()
				close(m.taskChan)
				return
			}
		}
	}
}

func (m *Manager) Stop() {
	m.close()
}
