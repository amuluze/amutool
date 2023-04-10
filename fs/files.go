// Package fs
// Date: 2023/4/10 15:42
// Author: Amu
// Description:
package fs

import (
	"os"
	"syscall"
)

func CloseOnExec(file *os.File) {
	if file != nil {
		syscall.CloseOnExec(int(file.Fd()))
	}
}
