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
			TimeKey:          "timestamp",
			LevelKey:         "level",
			NameKey:          "logger",
			CallerKey:        "caller",
			FunctionKey:      zapcore.OmitKey,
			MessageKey:       "message",
			StacktraceKey:    "stacktrace",
			LineEnding:       zapcore.DefaultLineEnding,
			EncodeLevel:      cEncodeLevel,
			EncodeTime:       cEncodeTime,
			EncodeDuration:   zapcore.NanosDurationEncoder,
			EncodeCaller:     cEncodeCaller,
			ConsoleSeparator: " || ",
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
			LineEnding:    zapcore.DefaultLineEnding,
			EncodeLevel:   zapcore.LowercaseLevelEncoder,
			EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendString(t.Format(TimeFormat))
			},
			EncodeDuration: zapcore.NanosDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
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
