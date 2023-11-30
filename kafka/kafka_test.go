// Package kafka
// Date: 2023/11/24 17:44
// Author: Amu
// Description:
package kafka

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/IBM/sarama"
)

func TestKafkaProducer(t *testing.T) {
	kafkaClient, err := New(
		WithClientID("test-producer"),
		WithProducerBrokers([]string{"localhost:9092"}),
	)
	if err != nil {
		fmt.Println("new kafka client error")
	}
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		strI := strconv.Itoa(i)
		kafkaClient.SendMessage("kafka-test", sarama.StringEncoder("testKey"+strI), sarama.StringEncoder("testValue"+strI))
	}
}

func TestKafkaConsumer(t *testing.T) {
	kafkaClient, err := New(
		WithClientID("consumer"),
		WithConsumerTopics([]string{"test"}),
		WithConsumerGroup("kafka-test"),
		WithConsumerBrokers([]string{"localhost:9092"}),
	)

	if err != nil {
		t.Fatal("kafka consumer init error:" + err.Error())
	}

	for {
		time.Sleep(time.Second)
		msg := <-kafkaClient.ConsumerMessages
		fmt.Printf("A message was received: %s\n", msg.Value)
	}
}
