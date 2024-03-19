// Package kafka
// Date: 2023/11/30 15:02
// Author: Amu
// Description:
package kafka

import "github.com/IBM/sarama"

var defaultVersion = sarama.V3_5_1_0

const (
	defaultRetryMax   = 1
	defaultAutoSubmit = true
	OffsetNewest      = sarama.OffsetNewest
	OffsetOldest      = sarama.OffsetOldest
	WaitNone          = sarama.NoResponse
	WaitLeader        = sarama.WaitForLocal
	WaitAll           = sarama.WaitForAll
)

type ConsumerMessage *sarama.ConsumerMessage

var (
	V0_8_2_0  = sarama.V0_8_2_0
	V0_8_2_1  = sarama.V0_8_2_1
	V0_8_2_2  = sarama.V0_8_2_2
	V0_9_0_0  = sarama.V0_9_0_0
	V0_9_0_1  = sarama.V0_9_0_1
	V0_10_0_0 = sarama.V0_10_0_0
	V0_10_0_1 = sarama.V0_10_0_1
	V0_10_1_0 = sarama.V0_10_1_0
	V0_10_1_1 = sarama.V0_10_0_1
	V0_10_2_0 = sarama.V0_10_2_0
	V0_10_2_1 = sarama.V0_10_2_1
	V0_10_2_2 = sarama.V0_10_2_2
	V0_11_0_0 = sarama.V0_11_0_0
	V0_11_0_1 = sarama.V0_11_0_1
	V0_11_0_2 = sarama.V0_11_0_2
	V1_0_0_0  = sarama.V1_0_0_0
	V1_0_1_0  = sarama.V1_0_1_0
	V1_0_2_0  = sarama.V1_0_2_0
	V1_1_0_0  = sarama.V1_1_0_0
	V1_1_1_0  = sarama.V1_1_1_0
	V2_0_0_0  = sarama.V2_0_0_0
	V2_0_1_0  = sarama.V2_0_1_0
	V2_1_0_0  = sarama.V2_1_0_0
	V2_1_1_0  = sarama.V2_1_1_0
	V2_2_0_0  = sarama.V2_2_0_0
	V2_2_1_0  = sarama.V2_2_1_0
	V2_2_2_0  = sarama.V2_2_2_0
	V2_3_0_0  = sarama.V2_3_0_0
	V2_3_1_0  = sarama.V2_3_1_0
	V2_4_0_0  = sarama.V2_4_0_0
	V2_4_1_0  = sarama.V2_4_1_0
	V2_5_0_0  = sarama.V2_5_0_0
	V2_5_1_0  = sarama.V2_5_1_0
	V2_6_0_0  = sarama.V2_6_0_0
	V2_6_1_0  = sarama.V2_6_1_0
	V2_6_2_0  = sarama.V2_6_2_0
	V2_7_0_0  = sarama.V2_7_0_0
	V2_7_1_0  = sarama.V2_7_1_0
	V2_8_0_0  = sarama.V2_8_0_0
	V2_8_1_0  = sarama.V2_8_1_0
	V2_8_2_0  = sarama.V2_8_2_0
	V3_0_0_0  = sarama.V3_0_0_0
	V3_0_1_0  = sarama.V3_0_1_0
	V3_0_2_0  = sarama.V3_0_2_0
	V3_1_0_0  = sarama.V3_1_0_0
	V3_1_1_0  = sarama.V3_1_1_0
	V3_1_2_0  = sarama.V3_1_2_0
	V3_2_0_0  = sarama.V3_2_0_0
	V3_2_1_0  = sarama.V3_2_1_0
	V3_2_2_0  = sarama.V3_2_2_0
	V3_2_3_0  = sarama.V3_2_3_0
	V3_3_0_0  = sarama.V3_3_0_0
	V3_3_1_0  = sarama.V3_3_1_0
	V3_3_2_0  = sarama.V3_3_2_0
	V3_4_0_0  = sarama.V3_4_0_0
	V3_4_1_0  = sarama.V3_4_1_0
	V3_5_0_0  = sarama.V3_5_0_0
	V3_5_1_0  = sarama.V3_5_1_0
	V3_6_0_0  = sarama.V3_6_0_0
)
