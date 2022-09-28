// Package logger
// Date: 2022/9/29 01:10
// Author: Amu
// Description:
package logger

import (
	"testing"
)

func TestInfo(t *testing.T) {
	Info("hello.", "status", 200)
}
