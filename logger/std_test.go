// Package logger
// Date: 2022/9/29 01:10
// Author: Amu
// Description:
package logger

import (
	"errors"
	"testing"
)

type User struct {
	Name string
}

func TestInfo(t *testing.T) {
	err := errors.New("bad request")

	AddError(err).AddInt("status", 200).Info("hello status")
}

func TestInitLogger(t *testing.T) {
	InitLogger(
		SetLogLevel("info"),
		SetLogFormat("json"),
	)
	err := errors.New("bad request")

	AddError(err).AddInt("status", 200).Info("hello")
	AddError(err).AddInt("status", 500).Error("hello")

	Info("good")
}
