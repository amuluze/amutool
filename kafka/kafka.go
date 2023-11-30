// Package kafka
// Date: 2023/11/24 17:35
// Author: Amu
// Description:
package kafka

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/IBM/sarama"
)

type Manager struct {
	Producer         sarama.AsyncProducer
	ConsumerGroup    sarama.ConsumerGroup
	ConsumerMessages chan *sarama.ConsumerMessage
	Consumers        chan Consumer
	option           *option
}

func New(opts ...Option) (*Manager, error) {
	opt := &option{
		version:                 defaultVersion,
		producerRetryMax:        defaultRetryMax,
		producerRequiredAcks:    WaitLeader,
		consumerOffsetsRetryMax: defaultRetryMax,
		autoSubmit:              defaultAutoSubmit,
	}

	for _, f := range opts {
		f(opt)
	}

	manager := &Manager{option: opt}

	if len(opt.producerBrokers) == 0 && len(opt.consumerBrokers) == 0 {
		return manager, errors.New("producer Brokers和consumer Brokers不能同时为空")
	}

	conf := sarama.NewConfig()
	conf.Version = opt.version                            // kafka 版本
	conf.ClientID = opt.clientID                          // 发送请求时传递给服务端的ClientID. 用来追溯请求源
	conf.Producer.Return.Successes = true                 // 接收服务器反馈的成功响应
	conf.Producer.Return.Errors = true                    // 接收服务器反馈的失败响应
	conf.Producer.Compression = sarama.CompressionGZIP    // 压缩格式
	conf.Producer.Retry.Max = opt.producerRetryMax        // 消息发送失败重试次数
	conf.Producer.RequiredAcks = opt.producerRequiredAcks // 消息发送确认
	conf.Producer.Partitioner = sarama.NewHashPartitioner // 通过msg中的key生成hash值,选择分区

	conf.Consumer.Offsets.Initial = OffsetNewest                  // 消费模式
	conf.Consumer.Offsets.AutoCommit.Enable = true                // 自动提交 offset
	conf.Consumer.Offsets.AutoCommit.Interval = 3 * time.Second   // 自动提交 offset 间隔
	conf.Consumer.Offsets.Retry.Max = opt.consumerOffsetsRetryMax // 消费失败重试次数
	conf.Consumer.Return.Errors = true

	if opt.username != "" && opt.password != "" {
		conf.Net.SASL.User = "admin"
		conf.Net.SASL.Password = "12345678"
		conf.Net.SASL.Enable = true
		conf.Net.SASL.Mechanism = sarama.SASLTypePlaintext
	}

	// 初始化一个异步的Producer
	if len(opt.producerBrokers) > 0 {
		producer, err := sarama.NewAsyncProducer(opt.producerBrokers, conf)
		if err != nil {
			return manager, err
		}

		go func(p sarama.AsyncProducer) {
			for {
				select {
				case err = <-p.Errors():
					if err != nil {
						fmt.Println("send msg to kafka failed:")
					}
				case <-p.Successes():
				}
			}
		}(producer)

		manager.Producer = producer
	}

	// 初始化指定消费组
	if len(opt.consumerBrokers) > 0 {
		group, err := sarama.NewConsumerGroup(opt.consumerBrokers, opt.consumerGroup, conf)
		if err != nil {
			return manager, err
		}

		fmt.Println("Consumer Group init success")

		if opt.autoSubmit {
			manager.ConsumerMessages = make(chan *sarama.ConsumerMessage, 256)
		} else {
			manager.Consumers = make(chan Consumer, 256)
		}

		go manager.consumerLoop(opt.consumerTopics)

		manager.ConsumerGroup = group
	}

	return manager, nil
}

// SendMessage 发送一条<key,value>到kafka指定topic中
func (m *Manager) SendMessage(topic string, key, value sarama.Encoder) {
	fmt.Println("Received a message")

	m.Producer.Input() <- &sarama.ProducerMessage{
		Topic: topic,
		Key:   key,
		Value: value,
	}
}

// consumerLoop 循环从kafka获取数据
func (m *Manager) consumerLoop(topics []string) {
	ctx := context.Background()
	fmt.Println("Consumer...")
	for {
		if err := m.ConsumerGroup.Consume(ctx, topics, m); err != nil {
			fmt.Println("Consume Error")
			return
		}
	}
}

// ConsumeClaim push message
func (m *Manager) ConsumeClaim(s sarama.ConsumerGroupSession, c sarama.ConsumerGroupClaim) error {
	for msg := range c.Messages() {
		// push message
		if msg != nil {
			// 自动提交开启时，将所有的消息推送到 m.ConsumerMessages，不需要推送ConsumerGroupSession
			if m.option.autoSubmit {
				s.MarkMessage(msg, "done")
				m.ConsumerMessages <- msg
				return nil
			}

			// 手动提交逻辑
			m.Consumers <- &consumer{
				message: msg,
				session: s,
			}
		}
	}

	return nil
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (m *Manager) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (m *Manager) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (m *Manager) Close() {
	var err error
	if m.Producer != nil {
		err = m.Producer.Close()
		if err != nil {
			fmt.Println("Close Producer failed")
		}
	}

	if m.ConsumerGroup != nil {
		err = m.ConsumerGroup.Close()
		if err != nil {
			fmt.Println("Close ConsumerGroup failed")
		}
	}
}
