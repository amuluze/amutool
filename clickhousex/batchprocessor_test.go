// Package clickhousex
// Date: 2023/12/7 11:20
// Author: Amu
// Description:
package clickhousex

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestBatchProcessor(t *testing.T) {
	db, err := NewDB(
		WithDebug(true),
		WithAddr("localhost:9000"),
		WithDatabase("gorm"),
		WithUsername("root"),
		WithPassword("123456"),
	)
	if err != nil {
		panic(err)
	}
	batchProcessor := NewBatchProcessor(
		WithDB(db),
		WithBatchInterval(15),
		WithBatchSize(100),
		WithModel(&User{}),
	)
	go batchProcessor.Start()

	for i := 1; i < 2231; i++ {
		u := map[string]interface{}{
			"id":       "id-" + strconv.Itoa(i),
			"username": "name-" + strconv.Itoa(i),
			"age":      int64(i),
			"sex":      1,
		}
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		err := batchProcessor.Add(u)
		if err != nil {
			fmt.Printf("add error: %#v\n", err)
			continue
		}
	}
	time.Sleep(3 * time.Second)
	batchProcessor.Stop()
}
