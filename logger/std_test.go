// Package logger
// Date: 2022/9/29 01:10
// Author: Amu
// Description:
package logger

import (
	"errors"
	"fmt"
	"testing"
)

type User struct {
	Name string
}

func TestInfo(t *testing.T) {
	std.Info(fmt.Sprintf("hello status: %d", 200))
}

func TestInitLogger(t *testing.T) {
	InitLogger(
		SetLogLevel("info"),
		SetLogFormat("json"),
	)
	err := errors.New("bad request")

	std.AddError(err).AddInt("status", 200).Info("hello")
	std.AddError(err).AddInt("status", 500).Error("hello")

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

	err := errors.New("bad error")
	log.AddString("ni", "hao").AddInt("status", 200).AddError(err).Info("test craete")
}
