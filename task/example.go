// Package task
// Date:   2025/4/17 12:10
// Author: Amu
// Description:
package task

import (
	"fmt"
	"time"
)

type ExampleTask struct{}

func (e *ExampleTask) Name() string {
	return "example"
}

func (e *ExampleTask) Schedule() string {
	return "@every 10s"
}

func (e *ExampleTask) Run() {
	now := time.Now().Unix()
	fmt.Printf("task %v has been handled in %d seconds\n", e.Name(), now)
}

type OneTimeTask struct{}

func (e *OneTimeTask) Name() string {
	return "one-time"
}

func (e *OneTimeTask) Schedule() string {
	return ""
}

func (e *OneTimeTask) Run() {
	fmt.Println("任务执行时间 (AddJob):", time.Now())
}
