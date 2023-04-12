// Package kafka
// Date: 2023/4/12 14:10
// Author: Amu
// Description:
package kafka

import (
	"encoding/json"

	"github.com/Shopify/sarama"
)

type ProducerConfig struct {
	Addrs        []string
	Topic        string
	RequiredAcks int16
}

type Producer struct {
	sarama.SyncProducer
	topic  string
	config *sarama.Config
}

func NewProducer(config *ProducerConfig) (*Producer, error) {
	p := Producer{}
	p.topic = config.Topic

	p.config = sarama.NewConfig()
	p.config.Producer.RequiredAcks = sarama.RequiredAcks(config.RequiredAcks)
	p.config.Producer.Retry.Max = ProducerRetryMax
	p.config.Producer.MaxMessageBytes = ProducerMaxMessageBodyBytes
	p.config.Producer.Timeout = ProducerTimeout
	p.config.Producer.Retry.Backoff = ProducerRetryBackoff
	p.config.Producer.Return.Errors = ProducerReturnErrors
	p.config.Producer.Compression = ProducerCompression
	p.config.Producer.CompressionLevel = ProducerCompressionLevel
	p.config.Producer.Return.Successes = ProducerReturnSuccesses
	p.config.Producer.Partitioner = ProducerPartitioner

	producer, err := sarama.NewSyncProducer(config.Addrs, p.config)
	if err != nil {
		return nil, err
	}
	p.SyncProducer = producer
	p.topic = config.Topic
	return &p, nil
}

func (p *Producer) Publish(messages []interface{}, key string) error {
	var msgs []*sarama.ProducerMessage
	for _, message := range messages {
		msg := &sarama.ProducerMessage{}
		msg.Topic = p.topic
		msg.Key = sarama.StringEncoder(key)

		bMsg, err := json.Marshal(message)
		if err != nil {
			return err
		}
		msg.Value = sarama.ByteEncoder(bMsg)
		msgs = append(msgs, msg)
	}

	return p.SendMessages(msgs)
}

func (p *Producer) Send(message interface{}) error {
	msg := &sarama.ProducerMessage{}
	msg.Topic = p.topic
	bMsg, err := json.Marshal(message)
	if err != nil {
		return err
	}
	msg.Value = sarama.ByteEncoder(bMsg)
	_, _, err = p.SendMessage(msg)
	return err
}
