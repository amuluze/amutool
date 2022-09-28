// Package logger
// Date: 2022/9/29 00:52
// Author: Amu
// Description:
package logger

import (
	"os"
	"path/filepath"

	rotator "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
)

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
