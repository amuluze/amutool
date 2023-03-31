// Package kafka
// Date: 2023/3/31 11:01
// Author: Amu
// Description:
package kafka

type Config struct {
	Kafka Kafka
}

type Kafka struct {
	Producer Producer
	Consumer Consumer
}

type Producer struct {
	Topics []string
	Retry  int
}

type Consumer struct {
	Topics            []string
	Group             string
	ChannelBufferSize int
	BulkSize          int
	BulkTimeout       string
	Retry             int
	Process           []string
}
