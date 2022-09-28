// Package logger
// Date: 2022/9/29 00:41
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

// Define key
const (
	TraceIDKey  = "trace_id"
	UserIDKey   = "user_id"
	UserNameKey = "user_name"
	TagKey      = "tag"
)

type (
	traceIDKey  struct{}
	userIDKey   struct{}
	userNameKey struct{}
	tagKey      struct{}
)
