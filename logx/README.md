### log
Golang log 工具封装

toml 配置实例
```yaml
[log]
Name = "test"                       # logger name
LogFile = "test.log"                # 日志输出文件名
LogLevel = "info"                   # 日志等级
LogFormat = "json"                  # 日志格式化方式，json text
LogFileRotationTime = 1             # 日志切割间隔，单位：D
LogFileMaxAge = 7                   # 日志切割间隔，单位：D
LogOutput = "file"                  # 日志输出位置，console file kafka
LogFileSuffix = ".%Y%m%d"           # 归档日志后缀，.20230411
```