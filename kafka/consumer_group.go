// Package kafka
// Date: 2023/4/12 14:10
// Author: Amu
// Description:
package kafka

import (
	"context"
	"gitee.com/amuluze/amutool/logx"
	"github.com/Shopify/sarama"
)

type GroupConfig struct {
	Addrs   []string
	GroupID string
	Topics  []string
}

type ConsumerGroup struct {
	sarama.ConsumerGroup
	groupID string
	topics  []string
	config  *sarama.Config
	handler sarama.ConsumerGroupHandler
}

func NewConsumerGroup(gc *GroupConfig, handler sarama.ConsumerGroupHandler) (*ConsumerGroup, error) {
	g := ConsumerGroup{}
	g.config = sarama.NewConfig()
	g.config.Version = Version
	g.config.Consumer.Return.Errors = ConsumerReturnErrors
	g.config.Consumer.Offsets.Initial = ConsumerOffset
	g.config.Consumer.Fetch.Min = ConsumerFetchMin
	g.config.Consumer.Fetch.Max = ConsumerFetchMax
	g.config.Consumer.Fetch.Default = ConsumerFetchDefault
	g.config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{ConsumerRebalanceStrategy}

	group, err := sarama.NewConsumerGroup(gc.Addrs, gc.GroupID, g.config)
	if err != nil {
		return nil, err
	}
	g.ConsumerGroup = group
	g.groupID = gc.GroupID
	g.topics = gc.Topics
	g.handler = handler
	return &g, nil
}

func (group *ConsumerGroup) Close() error {
	return group.Close()
}

func (group *ConsumerGroup) Run() error {
	for {
		err := group.Consume(context.TODO(), group.topics, group.handler)
		if err != nil {
			logx.Errorf("consumer error: %v", err)
		}
	}
}
