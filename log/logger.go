// Package log
// Date: 2022/8/26 15:28
// Author: Amu
// Description:
package log

import (
	"context"
	"os"
	"path/filepath"
	"sync"
	"time"

	rotator "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var once sync.Once

type Logger struct {
	*zap.SugaredLogger
	name    string
	lock    sync.Mutex
	loggers map[string]*Logger
}

func init() {
	once.Do(func() {
		std = &Logger{
			SugaredLogger: zap.New(zapcore.NewCore(zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()), zapcore.AddSync(os.Stdout), InfoLevel), zap.AddCaller(), zap.AddCallerSkip(2)).Sugar(),
			name:          "std",
			loggers:       make(map[string]*Logger),
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
		SugaredLogger: zap.New(zapcore.NewCore(encoder, writer, level), zap.AddCaller(), zap.AddCallerSkip(1)).Sugar(),
		name:          config.name,
		loggers:       make(map[string]*Logger),
	}
}

func getEncoder(config *Config) zapcore.Encoder {
	if config.logFormat == "text" {
		return zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	}

	var baseConfig = zapcore.EncoderConfig{
		// 下面以 Key 结尾的参数表示，Json格式日志中的 key
		TimeKey:       "timestamp",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "message",
		StacktraceKey: "stacktrace",
		EncodeLevel:   zapcore.LowercaseLevelEncoder, // 日志级别的以大写还是小写输出
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		}, // timestamp 时间字段的时间字符串格式
		EncodeDuration: zapcore.NanosDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, // caller 字典展示长路径韩式短路径，可以理解为相对路径和绝对路径
	}
	return zapcore.NewJSONEncoder(baseConfig)
}

func getWriter(config *Config) zapcore.WriteSyncer {
	if config.logOutput == "stdout" {
		return zapcore.AddSync(os.Stdout)
	}
	logFilePath := config.logFile
	if !filepath.IsAbs(config.logFile) {
		abspath, _ := filepath.Abs(filepath.Join(filepath.Dir(os.Args[0]), config.logFile))
		logFilePath = abspath
	}

	_log, _ := rotator.New(
		filepath.Join(logFilePath+config.logFileSuffix),
		// 生成软连接，指向最新的日志文件
		rotator.WithLinkName(logFilePath),
		// 保留文件期限
		rotator.WithMaxAge(config.logFileMaxAge),
		// 日志文件的切割间隔
		rotator.WithRotationTime(config.logFileRotationTime),
	)
	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(_log), zapcore.AddSync(os.Stdout))
}

func (l *Logger) CreateLogger(options ...Option) {
	l.lock.Lock()
	defer l.lock.Unlock()
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

	if _, ok := l.loggers[config.name]; ok {
		return
	}
	encoder := getEncoder(config)
	writer := getWriter(config)
	level := config.logLevel

	newLogger := &Logger{
		SugaredLogger: zap.New(zapcore.NewCore(encoder, writer, level), zap.AddCaller(), zap.AddCallerSkip(1)).Sugar(),
		name:          config.name,
		loggers:       make(map[string]*Logger),
	}
	l.loggers[config.name] = newLogger
}

func (l *Logger) NewTagContext(ctx context.Context, tag string) context.Context {
	return context.WithValue(ctx, tagKey{}, tag)
}

func (l *Logger) FromTagContext(ctx context.Context) string {
	v := ctx.Value(tagKey{})
	if v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

func (l *Logger) NewTraceIDContext(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, traceIDKey{}, traceID)
}

func (l *Logger) FromTraceIDContext(ctx context.Context) string {
	v := ctx.Value(traceIDKey{})
	if v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

func (l *Logger) NewUserIDContext(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIDKey{}, userID)
}

func (l *Logger) FromUserIDContext(ctx context.Context) string {
	v := ctx.Value(userIDKey{})
	if v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

func (l *Logger) NewUserNameContext(ctx context.Context, userName string) context.Context {
	return context.WithValue(ctx, userNameKey{}, userName)
}

func (l *Logger) FromUserNameContext(ctx context.Context) string {
	v := ctx.Value(userNameKey{})
	if v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

func (l *Logger) WithContext(ctx context.Context) *zap.SugaredLogger {
	var fields []interface{}

	if v := l.FromTraceIDContext(ctx); v != "" {
		fields = append(fields, TraceIDKey, v)
	}

	if v := l.FromUserIDContext(ctx); v != "" {
		fields = append(fields, UserIDKey, v)
	}

	if v := l.FromUserNameContext(ctx); v != "" {
		fields = append(fields, UserNameKey, v)
	}

	if v := l.FromTagContext(ctx); v != "" {
		fields = append(fields, TagKey, v)
	}

	return l.SugaredLogger.With(fields...)
}

func (l *Logger) Debug(args ...interface{}) {
	l.SugaredLogger.Debug(args...)
}

func (l *Logger) Debugf(ctx context.Context, args ...interface{}) {
	l.WithContext(ctx).Debug(args...)
}

func (l *Logger) Info(args ...interface{}) {
	l.SugaredLogger.Info(args...)
}

func (l *Logger) Infof(ctx context.Context, args ...interface{}) {
	l.WithContext(ctx).Info(args...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.SugaredLogger.Warn(args...)
}

func (l *Logger) Warnf(ctx context.Context, args ...interface{}) {
	l.WithContext(ctx).Warn(args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.SugaredLogger.Error(args...)
}

func (l *Logger) Errorf(ctx context.Context, args ...interface{}) {
	l.WithContext(ctx).Error(args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.SugaredLogger.Fatal(args...)
}

func (l *Logger) Fatalf(ctx context.Context, args ...interface{}) {
	l.WithContext(ctx).Fatal(args...)
}

func (l *Logger) Panic(args ...interface{}) {
	l.SugaredLogger.Panic(args...)
}

func (l *Logger) Panicf(ctx context.Context, args ...interface{}) {
	l.WithContext(ctx).Panic(args...)
}
