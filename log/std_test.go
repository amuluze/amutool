// Package log
// Date: 2022/9/29 16:36
// Author: Amu
// Description:
package log

import (
	"testing"
)

func TestInfo(t *testing.T) {
	Info("hello")
}

func TestInitLogger(t *testing.T) {
	InitLogger(
		SetLogLevel("info"),
		SetLogFormat("text"),
	)

	std.Info("hello")
	std.Error("hello")

	std.Info("good")
}

func TestCreateLogger(t *testing.T) {
	CreateLogger(
		SetName("logger"),
		SetLogLevel("info"),
		SetLogFormat("text"),
	)

	log := GetLoggerByName("logger")
	log.Info("hello logger")

	log.Error("test create")
}
