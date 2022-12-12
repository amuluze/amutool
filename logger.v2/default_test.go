// Package logger
// Date: 2022/12/12 22:48:56
// Author: Amu
// Description:
package logger

import (
	"testing"
)

func TestInfo(t *testing.T) {
	std.Info("hello status: %s", 200)
}
