// Package logger
// Date: 2022/12/12 22:33:33
// Author: Amu
// Description:
package logger

import (
	"fmt"
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
	*zap.Logger
	name    string
	lock    sync.Mutex
	loggers map[string]*Logger
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
		Logger: zap.New(
			zapcore.NewCore(encoder, writer, level),
			zap.AddCaller(),
			zap.AddCallerSkip(1),
		),
		name:    config.name,
		loggers: make(map[string]*Logger),
	}
	l.loggers[config.name] = newLogger
}

func (l *Logger) Info(args ...interface{}) {
	l.Logger.Info(fmt.Sprintln(args...))
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

func encodeLevel(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(level.CapitalString())
}

func encodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(TimeFormat))
}

func encodeCaller(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(caller.TrimmedPath())
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
