// Package task
// Date: 2023/7/19 15:30
// Author: Amu
// Description:
package task

import (
	"fmt"
	"sync"

	"github.com/robfig/cron/v3"
)

var (
	taskManager *Manager
)

type CronTask interface {
	Name() string
	Schedule() string
	Run()
}

type Manager struct {
	c           *cron.Cron
	quit        chan struct{}
	taskMap     sync.Map
	scheduleMap sync.Map
}

func InitTaskManager() {
	taskManager = &Manager{
		quit:        make(chan struct{}),
		c:           cron.New(cron.WithSeconds()),
		taskMap:     sync.Map{},
		scheduleMap: sync.Map{},
	}
}

func GetCommonTask() *Manager {
	return taskManager
}

func (t *Manager) Start() {
	t.c.Start()
}

func (t *Manager) Stop() {
	t.c.Stop()
}

func (t *Manager) AddTask(task CronTask) {
	// 如果已经存在，则先移除，再重新添加
	if taskID, ok := t.taskMap.Load(task.Name()); ok {
		t.c.Remove(taskID.(cron.EntryID))
		t.taskMap.Delete(task.Name())
		t.scheduleMap.Delete(task.Name())
	}
	fmt.Printf("task schedule %s \n", task.Schedule())
	taskID, err := t.c.AddJob(task.Schedule(), task)
	if err != nil {
		fmt.Printf("add job err:%v\n", err)
		return
	}
	t.taskMap.Store(task.Name(), taskID)
	t.scheduleMap.Store(task.Name(), task.Schedule())
}

// AddOnceTask 添加只执行一次且立即执行的任务
func (t *Manager) AddOnceTask(task CronTask) {
	go task.Run()
}

func (t *Manager) RemoveTask(task CronTask) {
	if taskID, ok := t.taskMap.Load(task.Name()); ok {
		t.c.Remove(taskID.(cron.EntryID))
		t.taskMap.Delete(task.Name())
		t.scheduleMap.Delete(task.Name())
	}
}

type Detail struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Schedule string `json:"schedule"`
	Prev     int64  `json:"prev"`
	Next     int64  `json:"next"`
}

func (t *Manager) GetTaskInfo() []Detail {
	details := make([]Detail, 0)
	t.taskMap.Range(func(key, value interface{}) bool {
		taskName := key.(string)
		taskID := value.(cron.EntryID)
		entry := t.c.Entry(taskID)
		schedule := ""
		if taskSchedule, ok := t.scheduleMap.Load(taskName); ok {
			schedule = taskSchedule.(string)
		}

		details = append(details, Detail{
			ID:       int(taskID),
			Name:     taskName,
			Schedule: schedule,
			Prev: func() int64 {
				if entry.Prev.Unix() > 0 {
					return entry.Prev.Unix()
				}
				return 0
			}(),
			Next: entry.Next.Unix(),
		})
		return true
	})
	return details
}
