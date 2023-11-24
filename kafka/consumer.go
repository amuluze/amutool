// Package kafka
// Date: 2023/4/12 16:15
// Author: Amu
// Description: Consumer 采用自动提交机制，保证将向 topic 中的消息写入下游即可
package kafka

import (
	"context"
	"fmt"
	"sync"

	"github.com/IBM/sarama"
	"github.com/pkg/errors"
)

// IConsumer 消息消费者
type IConsumer interface {
	//Run 开始进行阻塞消费
	Run(ctx context.Context) error
	//Close 关闭消费者
	Close() error
	//Confirm 确认单笔消息及该消息之前的所有消息
	Confirm(msg IMessage) error
	//ConfirmBatch 批量确认不同Topic与Partition的最后一笔消息
	ConfirmBatch(msgs []IMessage) error
}

type Consumer struct {
	Topics         []string                    // 待消费的 topic 列表
	MessageHandler OnMessageReceived           // 消息处理函数
	ErrorHandler   OnErrorOccurred             // 消息消费错误处理函数
	Session        sarama.ConsumerGroupSession // 消费者组会话
	Group          sarama.ConsumerGroup        // 消费者组对象
	ManualCommit   bool                        // 手动提交
	IsStarted      bool                        // 标记当前消费者是否开启消费
	IsStartedLock  sync.RWMutex
	SessionLock    sync.RWMutex
}

type Option func(c *Consumer)

func WithMessageHandler(m OnMessageReceived) Option {
	return func(c *Consumer) {
		c.MessageHandler = m
	}
}

func WithErrorHandler(e OnErrorOccurred) Option {
	return func(c *Consumer) {
		c.ErrorHandler = e
	}
}

func NewConsumer(config *ConsumerConfig, options ...Option) (*Consumer, error) {
	conf := sarama.NewConfig()
	conf.Consumer.Offsets.Initial = config.InitialOffset

	conf.Consumer.Retry.Backoff = ConsumerRetryBackoff
	conf.Consumer.Return.Errors = ConsumerReturnErrors
	conf.Consumer.Fetch.Min = ConsumerFetchMin
	conf.Consumer.Fetch.Default = ConsumerFetchDefault
	conf.Consumer.MaxWaitTime = ConsumerMaxWaitTime
	conf.Consumer.MaxProcessingTime = ConsumerMaxProcessingTime
	if !config.ManualCommit { // 自动提交
		conf.Consumer.Offsets.AutoCommit.Interval = ConsumerOffsetsAutoCommitInterval
		conf.Consumer.Offsets.AutoCommit.Enable = ConsumerOffsetsAutoCommitEnable
	} else {
		conf.Consumer.Offsets.AutoCommit.Enable = config.ManualCommit // 手动提交
	}
	if config.Username != "" && config.Password != "" {
		conf.Net.SASL.User = config.Username
		conf.Net.SASL.Password = config.Password
		conf.Net.SASL.Enable = true
		conf.Net.SASL.Mechanism = sarama.SASLTypePlaintext
	}
	group, err := sarama.NewConsumerGroup(config.Brokers, config.GroupID, conf)
	if err != nil {
		return nil, err
	}

	c := &Consumer{
		Topics: config.Topics,
		Group:  group,
	}

	for _, option := range options {
		option(c)
	}

	go func() {
		for e := range c.Group.Errors() {
			c.ErrorHandler(&ConsumeError{Err: e})
		}
	}()
	return c, nil
}

func (c *Consumer) Run(ctx context.Context) error {
	for {
		handler := defaultConsumerGroupHandler{c: c}
		errChan := make(chan error, 1)
		go func() {
			err := c.Group.Consume(ctx, c.Topics, handler)
			errChan <- err
		}()
		select {
		case <-ctx.Done():
			break
		case <-errChan:
			break
		}
	}
}

func (c *Consumer) Close() error {
	if c.Group != nil {
		_ = c.Group.Close()
	}
	return nil
}

func (c *Consumer) Confirm(msg IMessage) error {
	if err := c.checkManualCommit(); err != nil {
		return err
	}

	if err := c.checkConsumerStatus(); err != nil {
		return err
	}

	c.Session.MarkOffset(msg.Topic(), msg.Partition(), msg.Offset()+1, "")
	c.Session.Commit()
	return nil
}

func (c *Consumer) BatchConfirm(msgs []IMessage) error {
	if len(msgs) == 0 {
		return errors.New("无任何待确认数据")
	}
	if err := c.checkManualCommit(); err != nil {
		return err
	}
	if err := c.checkConsumerStatus(); err != nil {
		return err
	}
	if err := c.checkSession(); err != nil {
		return err
	}
	//对于Kafka而言，相同topic下的相同partition下的多个不同的offset，以最后一个确认的为准
	//比如offset的顺序是 1 2 3 4 5 6 7 8 4 ，那么提交的offset不是8而是4
	//所以需要按照topic和partition进行分组，以组中最大的offset+1作为提交依据
	//因为kafka是按照提交的offset作为下次的消费的依据，提交10则下次就从10开始消费，所以+1以便从11开始消费
	groups := make(map[string]IMessage)
	for _, v := range msgs {
		if v == nil {
			continue
		}

		key := fmt.Sprintf("%v-%v", v.Topic(), v.Partition())
		m, ok := groups[key]
		if !ok {
			groups[key] = v
			continue
		}
		if v.Offset() > m.Offset() {
			groups[key] = v
		}
	}
	for _, v := range groups {
		c.Session.MarkOffset(v.Topic(), v.Partition(), v.Offset()+1, "")
	}
	c.Session.Commit()
	return nil
}

func (c *Consumer) setStarted(started bool) {
	c.IsStartedLock.RLock()
	c.IsStarted = started
	c.IsStartedLock.RUnlock()
}

func (c *Consumer) setSession(sess sarama.ConsumerGroupSession) {
	c.SessionLock.Lock()
	c.Session = sess
	c.SessionLock.Unlock()
}

func (c *Consumer) checkSession() error {
	c.SessionLock.RLock()
	defer c.SessionLock.RUnlock()

	if c.Session == nil {
		return errors.New("无有效 Session")
	}
	return nil
}

func (c *Consumer) checkManualCommit() error {
	if !c.ManualCommit {
		return errors.New("自动提交模式不能手动确认")
	}
	return nil
}

func (c *Consumer) checkConsumerStatus() error {
	c.IsStartedLock.RLock()
	defer c.IsStartedLock.RUnlock()
	if !c.IsStarted {
		return errors.New("未开始消费数据")
	}
	return nil
}

type defaultConsumerGroupHandler struct {
	c *Consumer
}

// Setup 在新的会话开始前执行 早于ConsumeClaim
func (defaultConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error {
	fmt.Println("default consumer handler setup")
	return nil
}

// Cleanup 在会话结束时执行
func (defaultConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	fmt.Println("default consumer handler cleanup")
	return nil
}

func (h defaultConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// 当与broker断开连接，重连并且消费到新消息时会重新进入ConsumeClaim
	// 断开过程中收到的消息无法进行确认（会失败，但API不会返回错误） 只能待恢复后用新会话进行提交
	h.c.setSession(sess)
	h.c.setStarted(true)
	for msg := range claim.Messages() {
		if msg == nil || msg.Value == nil {
			break
		}

		message := &Message{
			Msg: msg,
		}
		if !h.c.ManualCommit {
			//自动提交模式下每收到一条消息就进行标记，保证自动Commit时可以确认收到的消息
			sess.MarkMessage(msg, "")
		}
		h.c.MessageHandler(message)
	}
	h.c.setStarted(false)
	return nil
}
