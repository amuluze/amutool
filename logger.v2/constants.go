// Package logger
// Date: 2022/12/12 22:32:33
// Author: Amu
// Description:
package logger

import "go.uber.org/zap/zapcore"

const (
	PanicLevel = zapcore.PanicLevel
	FatalLevel = zapcore.FatalLevel
	ErrorLevel = zapcore.ErrorLevel
	WarnLevel  = zapcore.WarnLevel
	InfoLevel  = zapcore.InfoLevel
	DebugLevel = zapcore.DebugLevel
)

const (
	TimeFormat = "2006-01-02 15:04:05"
)
