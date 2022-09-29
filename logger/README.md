### log
Golang log v2

#### example
- 不进行初始化直接使用时，只会在终端打印 log 信息
```go
package main

import (
	"github.com/amuluze/golog"
)

func main() {
	golog.Info("hello world.")
}
```

- 初始化，将 log 信息写入文件
```go
package main

import (
	"github.com/amuluze/amutool/logger"
	"time"
)

func main() {
	logger.InitLogger(
		logger.SetLogOutput("file"), // 定义日志输出方式为 file
		logger.SetLogFile("./logs/std.log"),  // 定义日志写入文件路径
		logger.SetLogLevel("error"), // 定义日志级别，默认 info
		logger.SetLogFileRotationTime(time.Hour),  // 定义日志切割间隔，默认为 1 天
		logger.SetLogFileMaxAge(time.Hour*24*7), // 定义日志保留时长，默认保留近 7 天的日志
		logger.SetLogFileSuffix(".%Y%m%d%H"),  // 定义日志归档文件后缀，这里注意要与切割间隔保持一直
	)

	logger.Info("hello", "good")
	logger.Error("this is a error message")
}
```

- 获取多个 log 实例，将不同的日志信息写入不同的文件
```go
package main

import (
	"context"
	"github.com/amuluze/amutool/logger"
	"time"
)

func InitLog() {
	logger.CreateLogger(
		logger.SetName("nlog"),
		logger.SetLogFile("./logs/nlog.log"),
		logger.SetLogLevel("info"),
		logger.SetLogOutput("file"),
		logger.SetLogFormat("json"),
		logger.SetLogFileRotationTime(time.Hour),
		logger.SetLogFileMaxAge(time.Hour*24*7),
		logger.SetLogFileSuffix(".%Y%m%d%H"),
	)

	logger.CreateLogger(
		logger.SetName("mlog"),
		logger.SetLogFile("./logs/mlog.log"),
		logger.SetLogLevel("info"),
		logger.SetLogOutput("file"),
		logger.SetLogFormat("text"),
		logger.SetLogFileRotationTime(time.Hour),
		logger.SetLogFileMaxAge(time.Hour*24*7),
		logger.SetLogFileSuffix(".%Y%m%d%H"),
	)
}

func main() {
	InitLog()
	nLogger := logger.GetLoggerByName("nlog")
	nLogger.Info("test info level log")

	ctx := nLogger.NewTraceIDContext(context.Background(), "123456")
	ctx = nLogger.NewTagContext(ctx, "__main__")
	nLogger.Infof(ctx, "test log with context")

	mLogger := logger.GetLoggerByName("mlog")
	mLogger.Error("test error level log")
}
```




#### 依赖
```shell
go.uber.org/zap
https://www.jianshu.com/p/fc90ea603ef2
https://github.com/arthurkiller/rollingwriter
```
