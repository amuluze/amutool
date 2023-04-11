// Package syncx
// Date: 2023/4/11 15:04
// Author: Amu
// Description:
package syncx

import "sync"

var (
	ReallyCrash = false
)

// WaitGroupWrapper 对sync.WaitGroup做了一层封装
// 可兼容sync.WaitGroup的方法，调用sync.WaitGroup.Wait()表示阻塞等待所有goroutine执行结束
type WaitGroupWrapper struct {
	sync.WaitGroup
}

// Wrap 开启一个goroutine, goroutine执行结束之后调用sync.WaitGroup.Done()
func (w *WaitGroupWrapper) Wrap(cb func()) {
	w.Add(1)
	go func() {
		cb()
		w.Done()
	}()
}

// WrapHandleCrash cb执行函数内部必须执行HandleCrash, 通过reallyCrash判断是否需要执行panic
func (w *WaitGroupWrapper) WrapHandleCrash(cb func(), reallyCrash bool) {
	if reallyCrash != ReallyCrash {
		ReallyCrash = reallyCrash
	}
	w.Add(1)
	go func() {
		// cb内部的panic无法被捕捉到
		// 必须在cb的内执行HandleCrash
		cb()
		w.Done()
	}()
}
