// Package kafka
// Date: 2023/4/12 14:10
// Author: Amu
// Description:
package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/IBM/sarama"
	"github.com/pkg/errors"
)

// IProducer 消息生产者
type IProducer interface {
	SendMessage(key string, data interface{}) (int, error)
	Close() error
}

// SyncProducer 同步生产者
type SyncProducer struct {
	Topic    string
	Producer sarama.SyncProducer
}

func NewSyncProducer(config *ProducerConfig) (*SyncProducer, error) {
	if len(config.Brokers) == 0 {
		return nil, errors.New("未指定brokers")
	}
	if config.Topic == "" {
		return nil, errors.New("未指定topic")
	}
	conf := sarama.NewConfig()
	conf.Producer.RequiredAcks = sarama.RequiredAcks(config.RequiredAcks)
	conf.Producer.Retry.Max = ProducerRetryMax
	conf.Producer.MaxMessageBytes = ProducerMaxMessageBodyBytes
	conf.Producer.Timeout = ProducerTimeout
	conf.Producer.Retry.Backoff = ProducerRetryBackoff
	conf.Producer.Return.Errors = ProducerReturnErrors
	conf.Producer.Compression = ProducerCompression
	conf.Producer.CompressionLevel = ProducerCompressionLevel
	conf.Producer.Return.Successes = ProducerReturnSuccesses
	conf.Producer.Partitioner = ProducerPartitioner
	if config.Username != "" && config.Password != "" {
		conf.Net.SASL.User = config.Username
		conf.Net.SASL.Password = config.Password
		conf.Net.SASL.Enable = true
		conf.Net.SASL.Mechanism = sarama.SASLTypePlaintext
	}

	sp, err := sarama.NewSyncProducer(config.Brokers, conf)
	if err != nil {
		return nil, err
	}
	p := &SyncProducer{
		Producer: sp,
		Topic:    config.Topic,
	}
	return p, nil
}

func (s *SyncProducer) SendMessage(key string, data interface{}) (int, error) {
	ret, err := json.Marshal(data)
	if err != nil {
		return 0, errors.New("data 序列化失败")
	}
	msg := &sarama.ProducerMessage{
		Topic: s.Topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.ByteEncoder(ret),
	}
	_, _, err = s.Producer.SendMessage(msg)
	if err != nil {
		return 0, err
	}
	return msg.Value.Length(), nil
}

func (s *SyncProducer) Close() error {
	err := s.Producer.Close()
	if err != nil {
		return err
	}
	s.Producer = nil
	return nil
}

// AsyncProducer 异步生产者
type AsyncProducer struct {
	Topic    string
	WG       *sync.WaitGroup
	Producer sarama.AsyncProducer
}

func NewAsnycProducer(config *ProducerConfig) (*AsyncProducer, error) {
	if len(config.Brokers) == 0 {
		return nil, errors.New("未指定brokers")
	}
	if config.Topic == "" {
		return nil, errors.New("未指定topic")
	}
	conf := sarama.NewConfig()
	conf.Producer.Compression = sarama.CompressionGZIP
	conf.Producer.RequiredAcks = sarama.RequiredAcks(config.RequiredAcks)
	// SyncProducer本质也是AsyncProducer
	// 因此Producer.Return.Errors和Successes均为true 可以处理错误和成功
	// 需要注意的是一旦设置为true，使用AsyncProducer时就会有成功消息和错误消息需要处理
	conf.Producer.Return.Errors = true
	conf.Producer.Return.Successes = true
	conf.Producer.Retry.Max = ProducerRetryMax
	conf.Producer.MaxMessageBytes = ProducerMaxMessageBodyBytes
	conf.Producer.Timeout = ProducerTimeout
	conf.Producer.Retry.Backoff = ProducerRetryBackoff
	conf.Producer.Return.Errors = ProducerReturnErrors
	conf.Producer.Compression = ProducerCompression
	conf.Producer.CompressionLevel = ProducerCompressionLevel
	conf.Producer.Return.Successes = ProducerReturnSuccesses
	conf.Producer.Partitioner = ProducerPartitioner
	if config.Username != "" && config.Password != "" {
		conf.Net.SASL.User = config.Username
		conf.Net.SASL.Password = config.Password
		conf.Net.SASL.Enable = true
		conf.Net.SASL.Mechanism = sarama.SASLTypePlaintext
	}
	ap, err := sarama.NewAsyncProducer(config.Brokers, conf)
	if err != nil {
		return nil, err
	}

	p := &AsyncProducer{
		Producer: ap,
		WG:       &sync.WaitGroup{},
		Topic:    config.Topic,
	}

	p.WG.Add(1)
	go func() {
		defer p.WG.Done()
		// 由于Producer.Return.Successes为true, 因此必须消费成功消息
		for m := range ap.Successes() {
			fmt.Printf("异步生产消息时，发生成功： %#v\n", m)
		}
		fmt.Println("success exit")
	}()

	p.WG.Add(1)
	go func() {
		defer p.WG.Done()
		//由于Producer.Return.Errors为true
		//因此无论是否有errorHandler都得消费错误信息
		for e := range ap.Errors() {
			fmt.Printf("异步生产消息时，发生错误： %#v\n", e)
		}
		fmt.Println("errors exit")
	}()

	return p, nil
}

func (a *AsyncProducer) SendMessage(key string, data interface{}) (int, error) {
	ret, err := json.Marshal(data)
	if err != nil {
		return 0, errors.New("data 序列化失败")
	}
	msg := &sarama.ProducerMessage{
		Topic: a.Topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.ByteEncoder(ret),
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	select {
	case a.Producer.Input() <- msg:
		return msg.Value.Length(), nil
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}

func (a *AsyncProducer) Close() (err error) {
	a.Producer.AsyncClose()
	a.Producer = nil
	a.WG.Wait()
	return nil
}
