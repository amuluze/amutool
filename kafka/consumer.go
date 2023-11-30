// Package kafka
// Date: 2023/11/30 15:03
// Author: Amu
// Description:
package kafka

import "github.com/IBM/sarama"

type Consumer interface {
	GetMsg() *sarama.ConsumerMessage
	Submit()
}

type consumer struct {
	message *sarama.ConsumerMessage
	session sarama.ConsumerGroupSession
}

func (c *consumer) GetMsg() *sarama.ConsumerMessage {
	return c.message
}

func (c *consumer) Submit() {
	c.session.MarkMessage(c.message, "done")
}
