// Package logger
// Date: 2022/12/12 22:37:14
// Author: Amu
// Description:
package logger

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var std *Logger

func CreateLogger(options ...Option) {
	std.CreateLogger(options...)
}

func Debug(args ...interface{}) {
	std.Debug(args...)
}

func Debugf(args ...interface{}) {
	std.Debug(args...)
}

func Info(args ...interface{}) {
	std.Info(args...)
}

func Infof(args ...interface{}) {
	std.Infof(args...)
}

func Warn(args ...interface{}) {
	std.Warn(args...)
}

func Warnf(args ...interface{}) {
	std.Warnf(args...)
}

func Error(args ...interface{}) {
	std.Error(args...)
}

func Errorf(args ...interface{}) {
	std.Errorf(args...)
}

func Fatal(args ...interface{}) {
	std.Fatal(args...)
}

func Fatalf(args ...interface{}) {
	std.Fatalf(args...)
}

func Panic(args ...interface{}) {
	std.Panic(args...)
}

func Panicf(args ...interface{}) {
	std.Panicf(args...)
}
