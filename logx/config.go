// Package logx
// Date: 2023/4/10 17:18
// Author: Amu
// Description:
package logx

import (
	"time"

	"go.uber.org/zap/zapcore"
)

type Config struct {
	name                string        `default:"std"`         // 【默认】Logger 名称
	logFile             string        `default:"scanner.log"` // 【默认】日志文件名称
	logLevel            zapcore.Level `default:"info"`        // 【默认】日志打印级别
	logFormat           string        `default:"text"`        // 【默认】日志打印样式，支持 text 和 json
	logFileRotationTime time.Duration `default:"1d"`          // 【默认】日志文件切割间隔
	logFileMaxAge       time.Duration `default:"7d"`          // 【默认】日志文件保留时间
	logOutput           string        `default:"stdout"`      // 【默认】日志输出位置，只会 stdout iohelper
	logFileSuffix       string        `default:".%Y%m%d"`     // 【默认】归档日志后缀
}

type Option func(*Config)

func SetName(name string) Option {
	return func(config *Config) {
		config.name = name
	}
}

func SetLogFile(logFile string) Option {
	return func(config *Config) {
		config.logFile = logFile
	}
}

func SetLogLevel(level string) Option {
	return func(config *Config) {
		switch level {
		case "debug":
			config.logLevel = zapcore.DebugLevel
		case "info":
			config.logLevel = zapcore.InfoLevel
		case "warn":
			config.logLevel = zapcore.WarnLevel
		case "error":
			config.logLevel = zapcore.ErrorLevel
		case "fatal":
			config.logLevel = zapcore.FatalLevel
		case "panic":
			config.logLevel = zapcore.PanicLevel
		}
	}
}

func SetLogFormat(format string) Option {
	return func(config *Config) {
		config.logFormat = format
	}
}

func SetLogOutput(output string) Option {
	return func(config *Config) {
		config.logOutput = output
	}
}

func SetLogFileRotationTime(duration time.Duration) Option {
	return func(config *Config) {
		config.logFileRotationTime = duration
	}
}

func SetLogFileMaxAge(duration time.Duration) Option {
	return func(config *Config) {
		config.logFileMaxAge = duration
	}
}

func SetLogFileSuffix(suffix string) Option {
	return func(config *Config) {
		config.logFileSuffix = suffix
	}
}