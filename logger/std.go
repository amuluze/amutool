// Package logger
// Date: 2022/9/29 00:28
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
						EncodeLevel:      cEncodeLevel,
						EncodeTime:       cEncodeTime,
						EncodeDuration:   zapcore.SecondsDurationEncoder,
						EncodeCaller:     cEncodeCaller,
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

func GetLoggerByName(name string) *Logger {
	if _, ok := std.loggers[name]; ok {
		return std.loggers[name]
	}
	return nil
}

func AddInt(key string, value int) *Logger {
	std.AddInt(key, value)
	return std
}

func AddInt32(key string, value int32) *Logger {
	std.AddInt32(key, value)
	return std
}

func AddInt64(key string, value int64) *Logger {
	std.AddInt64(key, value)
	return std
}

func AddFloat32(key string, value float32) *Logger {
	std.AddFloat32(key, value)
	return std
}

func AddFloat64(key string, value float64) *Logger {
	std.AddFloat64(key, value)
	return std
}

func AddString(key, value string) *Logger {
	std.AddString(key, value)
	return std
}

func AddTime(key string, value time.Time) *Logger {
	std.AddTime(key, value)
	return std
}

func AddDuration(key string, value time.Duration) *Logger {
	std.AddDuration(key, value)
	return std
}

func AddBool(key string, value bool) *Logger {
	std.AddBool(key, value)
	return std
}

func AddAny(key string, value interface{}) *Logger {
	std.AddAny(key, value)
	return std
}

func AddError(value error) *Logger {
	std.AddError(value)
	return std
}

func Debug(message string) {
	std.Debug(message)
}

func Info(message string) {
	std.Info(message)
}

func Warn(message string) {
	std.Warn(message)
}

func Error(message string) {
	std.Error(message)
}

func Fatal(message string) {
	std.Fatal(message)
}

func Panic(message string) {
	std.Panic(message)
}
