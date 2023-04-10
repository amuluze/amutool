// Package kafka
// Date: 2023/4/3 11:38
// Author: Amu
// Description:
package kafka

import (
	"time"

	"github.com/Shopify/sarama"
)

var (
	Version                   = sarama.V2_6_0_0
	ConsumerRebalanceStrategy = sarama.BalanceStrategyRange
	ProducerPartitioner       = sarama.NewHashPartitioner
)

const (
	ConsumerReturnErrors = true
	ConsumerFetchMin     = 1
	ConsumerFetchMax     = 0
	ConsumerOffset       = sarama.OffsetNewest
	ConsumerFetchDefault = 1 << 20
)

const (
	ProducerRetryMax            = 3
	ProducerReturnErrors        = true
	ProducerMaxMessageBodyBytes = 1000000
	ProducerReturnSuccesses     = true
	ProducerCompressionLevel    = sarama.CompressionLevelDefault
	ProducerCompression         = sarama.CompressionNone
	ProducerRetryBackoff        = 100 * time.Millisecond
	ProducerTimeout             = 10 * time.Second
)
