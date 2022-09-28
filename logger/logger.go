// Package logger
// Date: 2022/9/28 12:48
// Author: Amu
// Description:
package main

import (
	"sync"

	"go.uber.org/zap"
)

type Logger struct {
	*zap.Logger
	name    string
	lock    sync.Mutex
	loggers map[string]*Logger
}

//func main() {
//	logger, _ := zap.NewProduction()
//	defer logger.Sync() // flushes buffer, if any
//	sugar := logger.Sugar()
//	sugar.Infow("failed to fetch URL",
//		// Structured context as loosely typed key-value pairs.
//		"url", "http://www.baidu.com",
//		"attempt", 3,
//		"backoff", time.Second,
//	)
//	sugar.Infof("Failed to fetch URL: %s", "http://www.baidu.com")

//logger, _ := zap.NewProduction()
//defer logger.Sync()
//logger.Info("failed to fetch URL",
//	// Structured context as strongly typed Field values.
//	zap.String("url", "http://www.baidu.com"),
//	zap.Int("attempt", 3),
//	zap.Duration("backoff", time.Second),
//)
//}
