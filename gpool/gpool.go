// Package gpool
// Date: 2023/12/6 10:53
// Author: Amu
// Description:
package gpool

import (
	"errors"
	"sync"
	"time"
)

type TaskFunc func(args ...interface{})

type Task struct {
	f    TaskFunc
	args []interface{}
}

func (t *Task) Execute() {
	t.f(t.args...)
}

type GoroutinePool struct {
	taskQueue     chan *Task
	stopCh        chan struct{}
	maxWorkerSize int
	active        bool
	wg            *sync.WaitGroup
}

func NewGPool(workerSize, taskSize int) *GoroutinePool {
	return &GoroutinePool{
		maxWorkerSize: workerSize,
		taskQueue:     make(chan *Task, taskSize),
		stopCh:        make(chan struct{}, workerSize),
		active:        true,
		wg:            &sync.WaitGroup{},
	}
}

func (gp *GoroutinePool) Start() *GoroutinePool {
	for i := 0; i < gp.maxWorkerSize; i++ {
		go gp.run()
	}
	return gp
}

func (gp *GoroutinePool) Stop() {
	// 关闭接受新的任务
	gp.active = false
	// task.Execute() 执行完毕
	gp.wg.Wait()
	// 确保每个 task.run() 都结束
	for i := 0; i < gp.maxWorkerSize; i++ {
		gp.stopCh <- struct{}{}
	}
	time.Sleep(2 * time.Second)
	close(gp.taskQueue)
}

func (gp *GoroutinePool) Submit(f TaskFunc, args ...interface{}) error {
	if !gp.isActive() {
		return errors.New("task pool is not active")
	}

	gp.taskQueue <- &Task{
		f:    f,
		args: args,
	}
	return nil
}

func (gp *GoroutinePool) isActive() bool {
	return gp.active
}

func (gp *GoroutinePool) run() {
	for {
		select {
		case task := <-gp.taskQueue:
			gp.wg.Add(1)
			task.Execute()
			gp.wg.Done()
		case <-gp.stopCh:
			return
		}
	}
}
