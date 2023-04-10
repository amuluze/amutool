// Package consumer
// Date: 2023/3/31 11:37
// Author: Amu
// Description:
package consumer

import (
	"sync"

	"github.com/Shopify/sarama"
)

type Config struct {
	Addrs         []string
	Topic         string
	PartitionList []int32
}

type Consumer struct {
	addrs         []string
	consumer      sarama.Consumer
	topic         string
	partitionList []int32
	wg            sync.WaitGroup
}

func NewConsumer(config *Config) (*Consumer, error) {
	c := Consumer{}
	c.topic = config.Topic
	c.addrs = config.Addrs

	consumer, err := sarama.NewConsumer(c.addrs, nil)
	if err != nil {
		return nil, err
	}
	c.consumer = consumer

	partitionList, err := consumer.Partitions(config.Topic)
	if err != nil {
		return nil, err
	}
	c.partitionList = partitionList
	return &c, nil
}
