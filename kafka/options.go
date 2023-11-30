// Package kafka
// Date: 2023/11/30 15:01:29
// Author: Amu
// Description:
package kafka

import (
	"os"

	"github.com/IBM/sarama"
)

type Option func(*option)

type option struct {
	clientID                string
	version                 sarama.KafkaVersion // 版本
	username                string              // 用户名
	password                string              // 密码
	producerBrokers         []string            // 生产者连接地址
	producerRetryMax        int                 // 生产者最大重试次数
	producerRequiredAcks    sarama.RequiredAcks // 生产者确认模式
	consumerBrokers         []string            // 消费者连接地址
	consumerTopics          []string            // 消费主题
	consumerGroup           string              // 消费者组
	consumerOffsetInitial   int64               // 消费模式
	consumerOffsetsRetryMax int                 // 消费最大重试次数
	autoSubmit              bool                // 消费时自动提交
}

func WithClientID(CID string) Option {
	return func(o *option) {
		hostname, _ := os.Hostname()
		o.clientID = CID + "-" + hostname
	}
}

func WithVersion(version sarama.KafkaVersion) Option {
	return func(o *option) {
		o.version = version
	}
}

func WithUsername(username string) Option {
	return func(o *option) {
		o.username = username
	}
}

func WithPassword(password string) Option {
	return func(o *option) {
		o.password = password
	}
}

func WithProducerBrokers(brokers []string) Option {
	return func(o *option) {
		o.producerBrokers = brokers
	}
}

func WithProducerRetryMax(retryMax int) Option {
	return func(o *option) {
		o.producerRetryMax = retryMax
	}
}

func WithProducerRequiredAcks(acks sarama.RequiredAcks) Option {
	return func(o *option) {
		o.producerRequiredAcks = acks
	}
}

func WithConsumerBrokers(brokers []string) Option {
	return func(o *option) {
		o.consumerBrokers = brokers
	}
}

func WithConsumerTopics(topics []string) Option {
	return func(o *option) {
		o.consumerTopics = topics
	}
}

func WithConsumerGroup(group string) Option {
	return func(o *option) {
		o.consumerGroup = group
	}
}

func WithConsumerOffsetInitial(offset int64) Option {
	return func(o *option) {
		o.consumerOffsetInitial = offset
	}
}

func WithConsumerOffsetsRetryMax(retryMax int) Option {
	return func(o *option) {
		o.consumerOffsetsRetryMax = retryMax
	}
}

func WithAutoSubmit(auto bool) Option {
	return func(o *option) {
		o.autoSubmit = auto
	}
}
