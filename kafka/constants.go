// Package kafka
// Date: 2023/4/3 11:38
// Author: Amu
// Description:
package kafka

import (
	"time"

	"github.com/IBM/sarama"
)

var (
	Version                   = sarama.V2_6_0_0
	ConsumerRebalanceStrategy = sarama.BalanceStrategyRange
	ProducerPartitioner       = sarama.NewHashPartitioner
)

const (
	ConsumerFetchMin                  = 1
	ConsumerFetchDefault              = 1 << 20
	ConsumerRetryBackoff              = 2 * time.Second
	ConsumerMaxWaitTime               = 250 * time.Millisecond
	ConsumerMaxProcessingTime         = 100 * time.Millisecond
	ConsumerReturnErrors              = true
	ConsumerOffsetsAutoCommitEnable   = true
	ConsumerOffsetsAutoCommitInterval = 1 * time.Second
	ConsumerOfsettsInitial            = sarama.OffsetNewest
)

const (
	ProducerTimeout             = 10 * time.Second
	ProducerRetryMax            = 3
	ProducerRetryBackoff        = 100 * time.Millisecond
	ProducerReturnErrors        = true
	ProducerReturnSuccesses     = true
	ProducerMaxMessageBodyBytes = 1000000
	ProducerCompression         = sarama.CompressionNone
	ProducerCompressionLevel    = sarama.CompressionLevelDefault
)
