// Package task
// Date:   2025/4/17 12:10
// Author: Amu
// Description:
package task

import (
	"fmt"
	"testing"
	"time"
)

func TestGetCommonTask(t *testing.T) {
	InitTaskManager()
	tManager := GetCommonTask()
	tManager.Start()

	example := ExampleTask{}
	tManager.AddTask(&example)
	time.Sleep(20 * time.Second)
	t.Logf("task info: %#v", tManager.GetTaskInfo())

	time.Sleep(100 * time.Second)
	tManager.Stop()
}

func TestOneTimeTask(t *testing.T) {
	InitTaskManager()
	tManager := GetCommonTask()

	tManager.Start()
	one := OneTimeTask{}
	fmt.Println("one time task start")
	time.Sleep(5 * time.Second)
	tManager.AddOnceTask(&one)
	time.Sleep(10 * time.Second)
	fmt.Println("one time task stop")
	tManager.Stop()
}
