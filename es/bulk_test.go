// Package es
// Date: 2023/12/1 16:05
// Author: Amu
// Description:
package es

import (
	"fmt"
	"strconv"
	"testing"
)

func TestNewBulkProcessor(t *testing.T) {
	client, _ := NewEsClient(
		WithAddr("http://localhost:9200"),
		WithUsername("elastic"),
		WithPassword("123456"),
		WithSniff(true),
		WithDebug(true),
		WithHealthcheck(true),
	)
	processor, _ := NewBulkProcessor(
		client,
		WithBulkActions(100),        // worker 的队列容量
		WithBulkSize(50),            // 刷新大小限制，与 FlushInterval 共同作用
		WithBulkWorkers(1),          // worker 数量
		WithBulkFlushInterval("2s"), // 刷新间隔
	)
	fmt.Println("processor: ", processor)
	for i := 1000; i < 10000; i++ {
		user := &User{
			Username: "amu-" + strconv.Itoa(i),
			Age:      i,
			Sex:      1,
		}
		processor.Create(currentIndexName, user)
	}
	processor.Flush()
	processor.Close()
}
