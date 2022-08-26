// Package log
// Date: 2022/8/26 15:28
// Author: Amu
// Description:
package log

import (
	"context"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var std = &Logger{
	SugaredLogger: zap.New(zapcore.NewCore(zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()), zapcore.AddSync(os.Stdout), InfoLevel)).Sugar(),
	name:          "std",
	loggers:       make(map[string]*Logger),
}

func CreateLogger(options ...Option) {
	std.CreateLogger(options...)
}

func GetLoggerByName(name string) *Logger {
	if _, ok := std.loggers[name]; ok {
		return std.loggers[name]
	}
	return nil
}

func NewTagContext(ctx context.Context, tag string) context.Context {
	return std.NewTagContext(ctx, tag)
}

func FromTagContext(ctx context.Context) string {
	return std.FromTagContext(ctx)
}

func NewTraceIDContext(ctx context.Context, traceID string) context.Context {
	return std.NewTraceIDContext(ctx, traceID)
}

func FromTraceIDContext(ctx context.Context) string {
	return std.FromTraceIDContext(ctx)
}

func NewUserIDContext(ctx context.Context, userID string) context.Context {
	return std.NewUserIDContext(ctx, userID)
}

func FromUserIDContext(ctx context.Context) string {
	return std.FromUserIDContext(ctx)
}

func NewUserNameContext(ctx context.Context, userName string) context.Context {
	return std.NewUserNameContext(ctx, userName)
}

func FromUserNameContext(ctx context.Context) string {
	return std.FromUserNameContext(ctx)
}

func WithContext(ctx context.Context) {
	std.SugaredLogger = std.WithContext(ctx)
}

func Debug(args ...interface{}) {
	std.Debug(args...)
}

func Debugf(ctx context.Context, args ...interface{}) {
	WithContext(ctx)
	std.Info(args...)
}

func Info(args ...interface{}) {
	std.Info(args...)
}

func Infof(ctx context.Context, args ...interface{}) {
	WithContext(ctx)
	std.Info(args...)
}

func Warn(args ...interface{}) {
	std.Warn(args...)
}

func Warnf(ctx context.Context, args ...interface{}) {
	WithContext(ctx)
	std.Warn(args...)
}

func Error(args ...interface{}) {
	std.Error(args...)
}

func Errorf(ctx context.Context, args ...interface{}) {
	WithContext(ctx)
	std.Error(args...)
}

func Fatal(args ...interface{}) {
	std.Fatal(args...)
}

func Fatalf(ctx context.Context, args ...interface{}) {
	WithContext(ctx)
	std.Fatal(args...)
}

func Panic(args ...interface{}) {
	std.Panic(args...)
}

func Panicf(ctx context.Context, args ...interface{}) {
	WithContext(ctx)
	std.Panic(args...)
}
