// Package logx
// Date: 2023/4/10 17:24
// Author: Amu
// Description:
package logx

import (
	"os"
	"path/filepath"

	rotator "github.com/lestrrat-go/file-rotatelogs"

	"go.uber.org/zap/zapcore"
)

func getConsoleWriter() zapcore.WriteSyncer {
	return zapcore.AddSync(os.Stdout)
}

func getFileWriter(config *Config) zapcore.WriteSyncer {
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
	return zapcore.AddSync(_log)
}

func getWriter(config *Config) zapcore.WriteSyncer {
	if config.logOutput == "stdout" {
		return getConsoleWriter()
	} else {
		return zapcore.NewMultiWriteSyncer(getConsoleWriter(), getFileWriter(config))
	}
}
