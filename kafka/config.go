// Package kafka
// Date: 2023/11/24 14:24:57
// Author: Amu
// Description:
package kafka

type ConsumeError struct {
	Err error
}

// OnMessageReceived 消费消息处理函数
type OnMessageReceived func(msg IMessage)

// OnErrorOccurred 表示发生一个消费错误
type OnErrorOccurred func(err *ConsumeError)

type ConsumerConfig struct {
	Brokers       []string // kafka 的 broker 列表
	Topics        []string // 待消费的 topic 列表
	GroupID       string   // 消费者组ID
	Username      string
	Password      string
	InitialOffset int64 // 消费者 auto.offset.rest 模式
	ManualCommit  bool  // 手动提交
}

type ProducerConfig struct {
	Brokers      []string
	Username     string
	Password     string
	Topic        string
	Type         string // 生产方式，同步 or 异步
	RequiredAcks int64
}
