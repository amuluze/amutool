// Package logger
// Date: 2022/9/29 00:52
// Author: Amu
// Description:
package logger

import (
	"time"

	"go.uber.org/zap/zapcore"
)

func getEncoder(config *Config) zapcore.Encoder {
	if config.logFormat == "text" {
		return zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
			TimeKey:        "timestamp",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "message",
			StacktraceKey:  "stacktrace",
			LineEnding:     " ",
			EncodeLevel:    cEncodeLevel,
			EncodeTime:     cEncodeTime,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   cEncodeCaller,
		})
	} else {
		return zapcore.NewJSONEncoder(zapcore.EncoderConfig{
			// 下面以 Key 结尾的参数表示，Json格式日志中的 key
			TimeKey:       "timestamp",
			LevelKey:      "level",
			NameKey:       "logger",
			CallerKey:     "caller",
			FunctionKey:   zapcore.OmitKey,
			MessageKey:    "message",
			StacktraceKey: "stacktrace",
			EncodeLevel:   zapcore.LowercaseLevelEncoder, // 日志级别的以大写还是小写输出
			EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendString(t.Format(TimeFormat))
			}, // timestamp 时间字段的时间字符串格式
			EncodeDuration: zapcore.NanosDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder, // caller 字典展示长路径韩式短路径，可以理解为相对路径和绝对路径
		})
	}
}

func cEncodeLevel(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(level.CapitalString())
}

func cEncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(TimeFormat))
}

func cEncodeCaller(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(caller.TrimmedPath())
}
