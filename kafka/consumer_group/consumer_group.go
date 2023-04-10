// Package consumer_group
// Date: 2023/3/31 16:04
// Author: Amu
// Description:
package consumer_group

import (
	"context"

	"gitee.com/amuluze/amutool/kafka"
	"github.com/Shopify/sarama"
)

type Config struct {
	Addrs   []string
	GroupID string
	Topics  []string
}

type Group struct {
	group   sarama.ConsumerGroup
	config  *sarama.Config
	groupID string
	topics  []string
}

func NewConsumerGroup(config *Config) (*Group, error) {
	g := Group{}
	g.config = sarama.NewConfig()
	g.config.Version = kafka.Version
	g.config.Consumer.Return.Errors = kafka.ConsumerReturnErrors
	g.config.Consumer.Offsets.Initial = kafka.ConsumerOffset
	g.config.Consumer.Fetch.Min = kafka.ConsumerFetchMin
	g.config.Consumer.Fetch.Max = kafka.ConsumerFetchMax
	g.config.Consumer.Fetch.Default = kafka.ConsumerFetchDefault
	g.config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{kafka.ConsumerRebalanceStrategy}

	group, err := sarama.NewConsumerGroup(config.Addrs, config.GroupID, g.config)
	if err != nil {
		return nil, err
	}
	g.group = group
	g.groupID = config.GroupID
	g.topics = config.Topics
	return &g, nil
}

func (g *Group) RegisterHandlerAndConsumer(handler sarama.ConsumerGroupHandler) {
	ctx := context.Background()
	for {
		err := g.group.Consume(ctx, g.topics, handler)
		if err != nil {
			panic(err)
		}
	}
}
