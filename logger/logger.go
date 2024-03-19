// Package logger
// Date: 2024/03/19 17:09:20
// Author: Amu
// Description:
package logger

import (
	"fmt"
	"log/slog"
	"os"
	"time"
)

type Logger struct {
	*slog.Logger
	level slog.LevelVar
}

func NewLogger(level slog.Level) *Logger {
	var lvl slog.LevelVar
	lvl.Set(level)
	return &Logger{
		Logger: slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true,
			Level:     &lvl,
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
				if a.Key == slog.TimeKey {
					if t, ok := a.Value.Any().(time.Time); ok {
						a.Value = slog.StringValue(t.Format(time.DateTime))
					}
				}
				return a
			},
		})),
		level: lvl,
	}
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.Debug(fmt.Sprintf(format, args...))
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.Info(fmt.Sprintf(format, args...))
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.Warn(fmt.Sprintf(format, args...))
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.Error(fmt.Sprintf(format, args...))
}

func (l *Logger) SetDebugLevel() {
	l.level.Set(slog.LevelDebug)
}

func (l *Logger) SetInfoLevel() {
	l.level.Set(slog.LevelInfo)
}

func (l *Logger) SetWarnLevel() {
	l.level.Set(slog.LevelWarn)
}

func (l *Logger) SetErrorLevel() {
	l.level.Set(slog.LevelError)
}
