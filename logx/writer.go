// Package logx
// Date: 2023/4/10 17:24
// Author: Amu
// Description:
package logx

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/Shopify/sarama"

	rotator "github.com/lestrrat-go/file-rotatelogs"

	"go.uber.org/zap/zapcore"
)

var destinations []zapcore.WriteSyncer = []zapcore.WriteSyncer{
	zapcore.AddSync(os.Stdout),
}

func getFileWriter(config *Config) zapcore.WriteSyncer {
	if config.LogOutput != "file" {
		return nil
	}
	logFilePath := config.LogFile
	if !filepath.IsAbs(config.LogFile) {
		abspath, _ := filepath.Abs(filepath.Join(filepath.Dir(os.Args[0]), config.LogFile))
		logFilePath = abspath
	}

	_log, _ := rotator.New(
		filepath.Join(logFilePath+config.LogFileSuffix),
		// 生成软连接，指向最新的日志文件
		rotator.WithLinkName(logFilePath),
		// 保留文件期限，默认 7 天
		rotator.WithMaxAge(time.Duration(config.LogFileMaxAge)*time.Hour*24*7),
		// 日志文件的切割间隔，默认 1 天分割一个文件
		rotator.WithRotationTime(time.Duration(config.LogFileRotationTime)*time.Hour*24),
	)
	return zapcore.AddSync(_log)
}

func getKafkaWriter(config *Config) zapcore.WriteSyncer {
	var kl LogKafka
	var err error
	kl.Topic = "test_topic"
	cfg := sarama.NewConfig()
	cfg.Producer.RequiredAcks = sarama.WaitForAll
	cfg.Producer.Partitioner = sarama.NewHashPartitioner
	cfg.Producer.Return.Successes = true
	cfg.Producer.Return.Errors = true

	kl.Producer, err = sarama.NewSyncProducer([]string{"localhost:9000"}, cfg)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	return zapcore.AddSync(&kl)
}

func getWriter(config *Config) zapcore.WriteSyncer {
	if config.LogOutput == "file" {
		if fileWriter := getFileWriter(config); fileWriter != nil {
			destinations = append(destinations, fileWriter)
		}
	}
	if config.LogOutput == "kafka" {
		destinations = append(destinations, getKafkaWriter(config))
	}
	return zapcore.NewMultiWriteSyncer(destinations...)
}

type LogKafka struct {
	Producer sarama.SyncProducer
	Topic    string
}

func (lk *LogKafka) Write(p []byte) (n int, err error) {
	msg := &sarama.ProducerMessage{}
	msg.Topic = lk.Topic
	msg.Value = sarama.ByteEncoder(p)
	_, _, err = lk.Producer.SendMessage(msg)
	if err != nil {
		return
	}
	return
}
