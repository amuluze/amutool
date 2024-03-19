// Package logger
// Date: 2024/03/19 17:57:41
// Author: Amu
// Description:
package logger

import (
	"fmt"
	"log/slog"
	"testing"
)

func TestLogger(t *testing.T) {
	slog.Info("default logger message")
	logx := NewLogger(LevelInfo)
	logx.Debug("this is a debug message")
	logx.Info("this is a info message")
	logx.Error("this is a error message")
	logx.SetErrorLevel()
	fmt.Println()
	logx.Debug("this is a debug message")
	logx.Info("this is a info message")
	logx.Error("this is a error message")
}
