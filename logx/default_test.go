// Package logx
// Date: 2023/4/10 17:52
// Author: Amu
// Description:
package logx

import "testing"

func TestInfo(t *testing.T) {
	WithField(Field("trace_id", 123456))
	Info("hello golang")
}

func TestNewLogger(t *testing.T) {
	NewLogger(
		SetName("test"),
		SetLogFormat("json"),
		//SetLogFile("./logs/test.log"),
		//SetLogOutput("file"),
		SetLogLevel("debug"),
	)
	testLogger := GetLogger("test")
	testLogger.WithField(Field("trace_id", "1234456"))
	testLogger.Error("test logger error message")
}
