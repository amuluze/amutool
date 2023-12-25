// Package gpool
// Date: 2023/12/6 10:58
// Author: Amu
// Description:
package gpool

import (
	"strconv"
	"testing"
	"time"
)

func TestNewGPool(t *testing.T) {
	gp := NewGPool(3, 5)
	gp.Start()
	for i := 0; i < 10; i++ {
		err := gp.Submit(func(args ...interface{}) {
			time.Sleep(1 * time.Second)
			name := args[0].(string)
			t.Log("hello " + name + ", now is " + time.Now().Format("2006-01-02 04:05:00"))
		}, "jack-"+strconv.Itoa(i))
		if err != nil {
			break
		}
	}
	time.Sleep(10 * time.Second)
	gp.Stop()
}
