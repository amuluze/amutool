// Package logger
// Date: 2022/12/12 22:48:56
// Author: Amu
// Description:
package logger

import (
	"errors"
	"testing"
)

func TestInfo(t *testing.T) {
	std.Info("hello status: ", 200)
	err := errors.New("request failed")
	std.Errorf("this is a error message: %s", err)
}
