// Package task
// Date: 2023/4/4 16:51
// Author: Amu
// Description:
package task

import (
	"context"
	"sync"
	"sync/atomic"
	"time"
)

type Task struct {
	wg         sync.WaitGroup
	name       string
	plan       *Plan
	status     uint32
	handlerCtx context.Context
	handler    HandlerStruct
	ctx        context.Context
	cancel     context.CancelFunc
	timer      *time.Timer
}

type HandlerStruct func(ctx context.Context)

func NewTask(name string, handler HandlerStruct, plan *Plan, handlerCtx context.Context) *Task {
	return &Task{
		name:       name,
		handlerCtx: handlerCtx,
		handler:    handler,
		plan:       plan,
		timer:      time.NewTimer(plan.GetFirstDuration()),
	}
}
func (task *Task) run() {
	task.ctx, task.cancel = context.WithCancel(context.Background())
breakLoop:
	for {
		select {
		case <-task.timer.C:
			task.wg.Add(1)
			go func() {
				defer task.wg.Done()
				task.runHandler()
			}()

			duration := task.getNextDuration()
			if duration > 0 {
				task.timer.Reset(duration)
			} else {
				break breakLoop
			}
		case <-task.ctx.Done():
			break breakLoop
		}
	}

	task.close()
}

func (task *Task) close() {
	if !task.timer.Stop() {
		<-task.timer.C
	}
}

func (task *Task) stop() {
	if task.cancel != nil {
		task.cancel()
		task.wg.Wait()
	}
}

func (task *Task) getNextDuration() time.Duration {
	return task.plan.getDuration()
}

func (task *Task) runHandler() {
	if !atomic.CompareAndSwapUint32(&task.status, 0, 1) {
		return
	}

	task.handler(task.handlerCtx)
	task.status = 0
}
