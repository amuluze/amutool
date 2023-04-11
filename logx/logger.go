// Package logx
// Date: 2023/4/10 17:14
// Author: Amu
// Description:
package logx

import (
	"fmt"
	"os"
	"sync"

	"go.uber.org/zap/zapcore"

	"go.uber.org/zap"
)

var once sync.Once

type Logger struct {
	*zap.Logger
	name    string
	lock    sync.Mutex
	loggers map[string]*Logger
}

func init() {
	once.Do(func() {
		defaultLogger = &Logger{
			Logger: zap.New(
				zapcore.NewCore(
					zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
						TimeKey:          "time",
						LevelKey:         "level",
						NameKey:          "logger",
						CallerKey:        "caller",
						MessageKey:       "message",
						StacktraceKey:    "stacktrace",
						LineEnding:       zapcore.DefaultLineEnding,
						EncodeLevel:      cEncodeLevel,
						EncodeTime:       cEncodeTime,
						EncodeDuration:   zapcore.SecondsDurationEncoder,
						EncodeCaller:     cEncodeCaller,
						ConsoleSeparator: " || ",
					}),
					zapcore.AddSync(os.Stdout),
					zap.InfoLevel,
				),
				zap.AddCaller(),
				zap.AddCallerSkip(2),
			),
			name:    "default",
			loggers: make(map[string]*Logger),
		}
	})
}

func (l *Logger) NewLogger(options ...Option) {
	l.lock.Lock()
	defer l.lock.Unlock()
	config := &Config{
		Name:                "default",
		LogFile:             "default.log",
		LogLevel:            zap.InfoLevel,
		LogFormat:           "text",
		LogFileRotationTime: 1,
		LogFileMaxAge:       7,
		LogOutput:           "stdout",
		LogFileSuffix:       ".%Y%m%d",
	}
	for _, option := range options {
		option(config)
	}

	if _, ok := l.loggers[config.Name]; ok {
		return
	}
	encoder := getEncoder(config)
	writer := getWriter(config)
	level := config.LogLevel

	newLogger := &Logger{
		Logger: zap.New(
			zapcore.NewCore(encoder, writer, level),
			zap.AddCaller(),
			zap.AddCallerSkip(1),
		),
		name:    config.Name,
		loggers: make(map[string]*Logger),
	}
	l.loggers[config.Name] = newLogger
}

func (l *Logger) WithField(fields ...zap.Field) {
	l.Logger = l.Logger.With(fields...)
}

func (l *Logger) Debug(args ...interface{}) {
	l.Logger.Debug(fmt.Sprint(args...))
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.Logger.Debug(fmt.Sprintf(format, v...))
}

func (l *Logger) Info(args ...interface{}) {
	l.Logger.Info(fmt.Sprint(args...))
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.Logger.Info(fmt.Sprintf(format, v...))
}

func (l *Logger) Warn(args ...interface{}) {
	l.Logger.Warn(fmt.Sprint(args...))
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.Logger.Warn(fmt.Sprintf(format, v...))
}

func (l *Logger) Error(args ...interface{}) {
	l.Logger.Error(fmt.Sprint(args...))
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.Logger.Error(fmt.Sprintf(format, v...))
}

func (l *Logger) Fatal(args ...interface{}) {
	l.Logger.Fatal(fmt.Sprint(args...))
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.Logger.Fatal(fmt.Sprintf(format, v...))
}

func (l *Logger) Panic(args ...interface{}) {
	l.Logger.Panic(fmt.Sprint(args...))
}

func (l *Logger) Panicf(format string, v ...interface{}) {
	l.Logger.Panic(fmt.Sprintf(format, v...))
}
