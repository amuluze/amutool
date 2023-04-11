// Package logx
// Date: 2023/4/10 17:23
// Author: Amu
// Description:
package logx

import (
	"time"

	"go.uber.org/zap/zapcore"
)

func getEncoder(config *Config) zapcore.Encoder {
	if config.LogFormat == "text" {
		return zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
			TimeKey:          "time",
			LevelKey:         "level",
			NameKey:          "logger",
			CallerKey:        "caller",
			FunctionKey:      zapcore.OmitKey,
			MessageKey:       "message",
			StacktraceKey:    "stacktrace",
			LineEnding:       zapcore.DefaultLineEnding,
			EncodeLevel:      cEncodeLevel,
			EncodeTime:       cEncodeTime,
			EncodeDuration:   zapcore.SecondsDurationEncoder,
			EncodeCaller:     cEncodeCaller,
			ConsoleSeparator: " || ",
		})
	} else {
		return zapcore.NewJSONEncoder(zapcore.EncoderConfig{
			// 下面以 Key 结尾的参数表示，Json格式日志中的 key
			TimeKey:       "time",
			LevelKey:      "level",
			NameKey:       "logger",
			CallerKey:     "caller",
			FunctionKey:   zapcore.OmitKey,
			MessageKey:    "message",
			StacktraceKey: "stacktrace",
			LineEnding:    zapcore.DefaultLineEnding,
			EncodeLevel:   zapcore.LowercaseLevelEncoder,
			EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendString(t.Format(timeFormat))
			},
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		})
	}
}

func cEncodeLevel(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(level.CapitalString())
}

func cEncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(timeFormat))
}

func cEncodeCaller(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(caller.TrimmedPath())
}
