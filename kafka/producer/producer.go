// Package producer
// Date: 2023/3/31 11:37
// Author: Amu
// Description:
package producer

import (
	"encoding/json"

	"gitee.com/amuluze/amutool/kafka"
	"github.com/Shopify/sarama"
)

type Config struct {
	Addrs        []string
	Topic        string
	RequiredAcks int16
}

type Producer struct {
	topic    string
	config   *sarama.Config
	producer sarama.SyncProducer
}

func NewProducer(config *Config) *Producer {
	p := Producer{}
	p.topic = config.Topic

	p.config = sarama.NewConfig()
	p.config.Producer.RequiredAcks = sarama.RequiredAcks(config.RequiredAcks)
	p.config.Producer.Retry.Max = kafka.ProducerRetryMax
	p.config.Producer.MaxMessageBytes = kafka.ProducerMaxMessageBodyBytes
	p.config.Producer.Timeout = kafka.ProducerTimeout
	p.config.Producer.Retry.Backoff = kafka.ProducerRetryBackoff
	p.config.Producer.Return.Errors = kafka.ProducerReturnErrors
	p.config.Producer.Compression = kafka.ProducerCompression
	p.config.Producer.CompressionLevel = kafka.ProducerCompressionLevel
	p.config.Producer.Return.Successes = kafka.ProducerReturnSuccesses
	p.config.Producer.Partitioner = kafka.ProducerPartitioner

	producer, err := sarama.NewSyncProducer(config.Addrs, p.config)
	if err != nil {
		panic(err)
	}
	p.producer = producer
	return &p
}

func (p *Producer) SendMessage(message interface{}, key string) (int32, int64, error) {
	msg := &sarama.ProducerMessage{}
	msg.Topic = p.topic
	msg.Key = sarama.StringEncoder(key)

	bMsg, err := json.Marshal(message)
	if err != nil {
		return -1, -1, err
	}
	msg.Value = sarama.ByteEncoder(bMsg)
	return p.producer.SendMessage(msg)
}
