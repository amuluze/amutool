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
						TimeKey:          "timestamp",
						LevelKey:         "level",
						NameKey:          "logger",
						CallerKey:        "caller",
						MessageKey:       "message",
						StacktraceKey:    "stacktrace",
						LineEnding:       " ",
						EncodeLevel:      cEncodeLevel,
						EncodeTime:       cEncodeTime,
						EncodeDuration:   zapcore.SecondsDurationEncoder,
						EncodeCaller:     cEncodeCaller,
						ConsoleSeparator: " || ",
					}),
					zapcore.AddSync(os.Stdout),
					InfoLevel,
				),
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
		Logger:  zap.New(zapcore.NewCore(encoder, writer, level)),
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
	return std.AddInt(key, value)
}

func AddInt32(key string, value int32) *Logger {
	return std.AddInt32(key, value)
}

func AddInt64(key string, value int64) *Logger {
	return std.AddInt64(key, value)
}

func AddFloat32(key string, value float32) *Logger {
	return std.AddFloat32(key, value)
}

func AddFloat64(key string, value float64) *Logger {
	return std.AddFloat64(key, value)
}

func AddString(key, value string) *Logger {
	return std.AddString(key, value)
}

func AddTime(key string, value time.Time) *Logger {
	return std.AddTime(key, value)
}

func AddDuration(key string, value time.Duration) *Logger {
	return std.AddDuration(key, value)
}

func AddBool(key string, value bool) *Logger {
	return std.AddBool(key, value)
}

func AddAny(key string, value interface{}) *Logger {
	return std.AddAny(key, value)
}

func AddError(value error) *Logger {
	return std.AddError(value)
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
