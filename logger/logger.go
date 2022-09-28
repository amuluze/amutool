// Package logger
// Date: 2022/9/28 12:48
// Author: Amu
// Description:
package logger

import (
	"context"
	"sync"
	"time"

	"go.uber.org/zap/zapcore"

	"go.uber.org/zap"
)

var once sync.Once

type Logger struct {
	*zap.Logger
	name    string
	lock    sync.Mutex
	ctx     context.Context
	loggers map[string]*Logger
}

func (l *Logger) With(kv ...interface{}) {
	l.Logger.With()
}

func (l *Logger) CreateLogger(options ...Option) {
	l.lock.Lock()
	defer l.lock.Unlock()
	config := &Config{
		name:                "std",
		logFile:             "default.log",
		logLevel:            InfoLevel,
		logFormat:           "text",
		logFileRotationTime: time.Hour * 24,
		logFileMaxAge:       time.Hour * 24 * 7,
		logOutput:           "stdout",
		logFileSuffix:       ".%Y%m%d",
	}
	for _, option := range options {
		option(config)
	}

	if _, ok := l.loggers[config.name]; ok {
		return
	}
	encoder := getEncoder(config)
	writer := getWriter(config)
	level := config.logLevel

	newLogger := &Logger{
		Logger:  zap.New(zapcore.NewCore(encoder, writer, level), zap.AddCaller(), zap.AddCallerSkip(1)),
		name:    config.name,
		loggers: make(map[string]*Logger),
	}
	l.loggers[config.name] = newLogger
}
