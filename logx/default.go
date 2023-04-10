// Package logx
// Date: 2023/4/10 17:39
// Author: Amu
// Description:
package logx

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var defaultLogger *Logger

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
				zap.AddCallerSkip(1),
			),
			name:    "default",
			loggers: make(map[string]*Logger),
		}
	})
}

func NewLogger(options ...Option) {
	defaultLogger.NewLogger(options...)
}

func GetLogger(name string) *Logger {
	if _, ok := defaultLogger.loggers[name]; ok {
		return defaultLogger.loggers[name]
	}
	return nil
}

func Debug(args ...interface{}) {
	defaultLogger.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	defaultLogger.Debugf(format, args...)
}

func Info(args ...interface{}) {
	defaultLogger.Info(args...)
}

func Infof(format string, args ...interface{}) {
	defaultLogger.Infof(format, args...)
}

func Warn(args ...interface{}) {
	defaultLogger.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	defaultLogger.Warnf(format, args...)
}

func Error(args ...interface{}) {
	defaultLogger.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	defaultLogger.Errorf(format, args...)
}

func Fatal(args ...interface{}) {
	defaultLogger.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	defaultLogger.Fatalf(format, args...)
}

func Panic(args ...interface{}) {
	defaultLogger.Panic(args...)
}

func Panicf(format string, args ...interface{}) {
	defaultLogger.Panicf(format, args...)
}
