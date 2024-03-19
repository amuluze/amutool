// Package logger
// Date: 2024/03/19 17:57:41
// Author: Amu
// Description:
package logger

import (
	"fmt"
	"testing"
)

func TestTextLogger(t *testing.T) {
	logx := NewTextLogger()
	logx.SetDebugLevel()
	logx.Debug("this is a debug message")
	logx.Info("this is a info message")
	logx.Error("this is a error message")
	fmt.Println("--------------------------")
	logx.SetErrorLevel()
	logx.Debug("this is a debug message")
	logx.Info("this is a info message")
	logx.Error("this is a error message")
}

func TestJsonLogger(t *testing.T) {
	logx := NewJsonLogger()
	logx.SetDebugLevel()
	logx.Debug("this is a debug message")
	logx.Info("this is a info message")
	logx.Error("this is a error message")
	fmt.Println("--------------------------")
	logx.SetErrorLevel()
	logx.Debug("this is a debug message")
	logx.Info("this is a info message")
	logx.Error("this is a error message")
}

func TestJsonFileLogger(t *testing.T) {
	logx := NewJsonFileLogger(
		SetName("test"),
		SetLogFile("test.log"),
		SetLogLevel("info"),
		SetLogFileRotationTime(1),
		SetLogFileMaxAge(7),
		SetLogFileSuffix(".%Y%m%d"),
	)
	logx.SetDebugLevel()
	logx.Debug("this is a debug message")
	logx.Info("this is a info message")
	logx.Error("this is a error message")
	fmt.Println("--------------------------")
	logx.SetErrorLevel()
	logx.Debug("this is a debug message")
	logx.Info("this is a info message")
	logx.Error("this is a error message")
}
