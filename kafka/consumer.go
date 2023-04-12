// Package kafka
// Date: 2023/4/12 16:15
// Author: Amu
// Description:
package kafka

import (
	"log"
	"os"
	"os/signal"

	"github.com/Shopify/sarama"
)

type ConsumerConfig struct {
	Addrs     []string
	Partition int32
	Topic     string
}

type Consumer struct {
	sarama.Consumer
	topic     string
	partition int32
}

func NewConsumer(config ConsumerConfig) (*Consumer, error) {
	consumer := Consumer{}
	consumer.topic = config.Topic
	consumer.partition = config.Partition
	c, err := sarama.NewConsumer(config.Addrs, sarama.NewConfig())
	if err != nil {
		return nil, err
	}
	consumer.Consumer = c
	return &consumer, err
}

func (c *Consumer) Run() error {
	partitionConsumer, err := c.ConsumePartition(c.topic, c.partition, sarama.OffsetNewest)
	if err != nil {
		return err
	}

	// Trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

Exit:
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			log.Printf("msg: %v", msg)
		case <-signals:
			break Exit
		}
	}
	return nil
}
