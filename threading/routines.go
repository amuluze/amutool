// Package threading
// Date: 2023/4/10 14:14
// Author: Amu
// Description:
package threading

func GoSafe(fn func()) {
	go RunSafe(fn)
}

func RunSafe(fn func()) {
	defer rescue.Recover()

	fn()
}
