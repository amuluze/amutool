// Package logger
// Date: 2022/9/28 12:48
// Author: Amu
// Description:
package logger

import (
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var once sync.Once

type Logger struct {
	*zap.Logger
	name    string
	lock    sync.Mutex
	fields  []zap.Field
	loggers map[string]*Logger
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
		Logger: zap.New(
			zapcore.NewCore(encoder, writer, level),
			zap.AddCaller(),
			zap.AddCallerSkip(1),
		),
		name:    config.name,
		loggers: make(map[string]*Logger),
	}
	l.loggers[config.name] = newLogger
}

func (l *Logger) AddInt(key string, value int) *Logger {
	l.fields = append(l.fields, zap.Int(key, value))
	return l
}

func (l *Logger) AddInt32(key string, value int32) *Logger {
	l.fields = append(l.fields, zap.Int32(key, value))
	return l
}

func (l *Logger) AddInt64(key string, value int64) *Logger {
	l.fields = append(l.fields, zap.Int64(key, value))
	return l
}

func (l *Logger) AddFloat32(key string, value float32) *Logger {
	l.fields = append(l.fields, zap.Float32(key, value))
	return l
}

func (l *Logger) AddFloat64(key string, value float64) *Logger {
	l.fields = append(l.fields, zap.Float64(key, value))
	return l
}

func (l *Logger) AddString(key, value string) *Logger {
	l.fields = append(l.fields, zap.String(key, value))
	return l
}

func (l *Logger) AddTime(key string, value time.Time) *Logger {
	l.fields = append(l.fields, zap.Time(key, value))
	return l
}

func (l *Logger) AddDuration(key string, value time.Duration) *Logger {
	l.fields = append(l.fields, zap.Duration(key, value))
	return l
}

func (l *Logger) AddBool(key string, value bool) *Logger {
	l.fields = append(l.fields, zap.Bool(key, value))
	return l
}

func (l *Logger) AddAny(key string, any interface{}) *Logger {
	l.fields = append(l.fields, zap.Any(key, any))
	return l
}

func (l *Logger) AddError(value error) *Logger {
	l.fields = append(l.fields, zap.Error(value))
	return l
}

func (l *Logger) Debug(message string) {
	l.Logger.Debug(message, l.fields...)
	l.fields = l.fields[0:0]
}

func (l *Logger) Info(message string) {
	l.Logger.Info(message, l.fields...)
	l.fields = l.fields[0:0]
}

func (l *Logger) Warn(message string) {
	l.Logger.Warn(message, l.fields...)
	l.fields = l.fields[0:0]
}

func (l *Logger) Error(message string) {
	l.Logger.Error(message, l.fields...)
	l.fields = l.fields[0:0]
}

func (l *Logger) Fatal(message string) {
	l.Logger.Fatal(message, l.fields...)
	l.fields = l.fields[0:0]
}

func (l *Logger) Panic(message string) {
	l.Logger.Panic(message, l.fields...)
	l.fields = l.fields[0:0]
}
