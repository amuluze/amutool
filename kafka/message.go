// Package kafka
// Date: 2023/11/24 14:25:47
// Author: Amu
// Description:
package kafka

import "github.com/IBM/sarama"

// IMessage 表示一个Kafka的数据消息
type IMessage interface {
	// Topic 当前数据所属Topic
	Topic() string
	// Partition 当前数据所属Topic下的分区
	Partition() int32
	// Offset 当前数据的分区偏移量
	Offset() int64
	Key() string
	// Data 实际数据内容
	Data() []byte
}

type Message struct {
	// 实际消息数据
	Msg *sarama.ConsumerMessage
}

func (m *Message) Data() []byte {
	return m.Msg.Value
}

func (m *Message) Key() string {
	return string(m.Msg.Key)
}

func (m *Message) Topic() string {
	return m.Msg.Topic
}

func (m *Message) Partition() int32 {
	return m.Msg.Partition
}

func (m *Message) Offset() int64 {
	return m.Msg.Offset
}
