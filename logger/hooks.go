// Package logger
// Date: 2022/9/29 17:46
// Author: Amu
// Description:
package logger

import (
	"fmt"

	"go.uber.org/atomic"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Hooks 创建时作为参数添加
func Hooks() zap.Option {
	count := &atomic.Int64{}
	return zap.Hooks(func(entry zapcore.Entry) error {
		fmt.Println("count:", count.Inc(), "msg:", entry.Message, "level:", entry.Level)
		return nil
	})
}
