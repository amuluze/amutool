// Package logx
// Date: 2023/4/10 17:18
// Author: Amu
// Description:
package logx

import (
	"go.uber.org/zap/zapcore"
)

type Config struct {
	Name                string        `load:"std"`         // 【默认】Logger 名称
	LogFile             string        `load:"scanner.log"` // 【默认】日志文件名称
	LogLevel            zapcore.Level `load:"info"`        // 【默认】日志打印级别
	LogFormat           string        `load:"text"`        // 【默认】日志打印样式，支持 text 和 json
	LogFileRotationTime int           `load:"1"`           // 【默认】日志文件切割间隔，单位 D
	LogFileMaxAge       int           `load:"7"`           // 【默认】日志文件保留时间，单位 D
	LogOutput           string        `load:"stdout"`      // 【默认】日志输出位置，只会 stdout iohelper
	LogFileSuffix       string        `load:".%Y%m%d"`     // 【默认】归档日志后缀
}

type Option func(*Config)

func SetName(name string) Option {
	return func(config *Config) {
		config.Name = name
	}
}

func SetLogFile(logFile string) Option {
	return func(config *Config) {
		config.LogFile = logFile
	}
}

func SetLogLevel(level string) Option {
	return func(config *Config) {
		switch level {
		case "debug":
			config.LogLevel = zapcore.DebugLevel
		case "info":
			config.LogLevel = zapcore.InfoLevel
		case "warn":
			config.LogLevel = zapcore.WarnLevel
		case "error":
			config.LogLevel = zapcore.ErrorLevel
		case "fatal":
			config.LogLevel = zapcore.FatalLevel
		case "panic":
			config.LogLevel = zapcore.PanicLevel
		}
	}
}

func SetLogFormat(format string) Option {
	return func(config *Config) {
		config.LogFormat = format
	}
}

func SetLogOutput(output string) Option {
	return func(config *Config) {
		config.LogOutput = output
	}
}

func SetLogFileRotationTime(duration int) Option {
	return func(config *Config) {
		config.LogFileRotationTime = duration
	}
}

func SetLogFileMaxAge(duration int) Option {
	return func(config *Config) {
		config.LogFileMaxAge = duration
	}
}

func SetLogFileSuffix(suffix string) Option {
	return func(config *Config) {
		config.LogFileSuffix = suffix
	}
}
