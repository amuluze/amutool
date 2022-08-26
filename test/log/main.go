// Package main
// Date: 2022/8/26 15:36
// Author: Amu
// Description:
package main

import (
	"time"

	"gitee.com/amuluze/amutool/log"
)

func InitLogger() {
	log.InitLogger(
		log.SetLogOutput("file"),
		log.SetLogFile("./logs/std.log"),
		log.SetLogLevel("info"),
		log.SetLogFormat("json"),
	)

	log.CreateLogger(
		log.SetName("nlog"),
		log.SetLogFile("./logs/nlog.log"),
		log.SetLogLevel("info"),
		log.SetLogOutput("file"),
		log.SetLogFormat("json"),
		log.SetLogFileRotationTime(time.Hour),
		log.SetLogFileMaxAge(time.Hour*24*7),
		log.SetLogFileSuffix(".%Y%m%d%H"),
	)

	log.CreateLogger(
		log.SetName("mlog"),
		log.SetLogFile("./logs/mlog.log"),
		log.SetLogLevel("info"),
		log.SetLogOutput("file"),
		log.SetLogFormat("text"),
		log.SetLogFileRotationTime(time.Hour),
		log.SetLogFileMaxAge(time.Hour*24*7),
		log.SetLogFileSuffix(".%Y%m%d%H"),
	)
}

func main() {
	InitLogger()
	log.Info("hello amutool log")

	nlog := log.GetLoggerByName("nlog")
	nlog.Error("test error log")
}
