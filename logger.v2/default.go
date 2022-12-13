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

func init() {
	once.Do(func() {
		std = &Logger{
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
						EncodeLevel:      encodeLevel,
						EncodeTime:       encodeTime,
						EncodeDuration:   zapcore.SecondsDurationEncoder,
						EncodeCaller:     encodeCaller,
						ConsoleSeparator: " || ",
					}),
					zapcore.AddSync(os.Stdout),
					InfoLevel,
				),
				zap.AddCaller(),
				zap.AddCallerSkip(1),
			),
			name:    "std",
			loggers: make(map[string]*Logger),
		}
	})
}

func InitLogger(options ...Option) {
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

	encoder := getEncoder(config)
	writer := getWriter(config)
	level := config.logLevel

	std = &Logger{
		Logger: zap.New(
			zapcore.NewCore(encoder, writer, level),
			zap.AddCaller(),
			zap.AddCallerSkip(1),
		),
		name:    config.name,
		loggers: make(map[string]*Logger),
	}
}

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
