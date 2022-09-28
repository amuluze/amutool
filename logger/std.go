// Package logger
// Date: 2022/9/29 00:28
// Author: Amu
// Description:
package logger

import (
	"context"
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
					zapcore.NewConsoleEncoder(commonConfig),
					zapcore.AddSync(os.Stdout),
					InfoLevel,
				),
			),
			name:    "std",
			ctx:     context.Background(),
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
		ctx:     context.Background(),
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

func Info(args ...interface{}) {
	std.Sugar().Info()
}
