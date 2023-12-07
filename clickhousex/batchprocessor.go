// Package clickhousex
// Date: 2023/12/7 10:30
// Author: Amu
// Description:
package clickhousex

import (
	"fmt"
	"time"

	"gitee.com/amuluze/amutool/errors"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Age      int64  `json:"age"`
	Sex      int64  `json:"sex"`
}

type BatchProcessor struct {
	db          *DB
	dataCh      chan map[string]interface{}
	stopCh      chan interface{}
	data        []map[string]interface{}
	model       interface{}
	batchSize   int
	active      bool
	batchTicker *time.Ticker
}

func NewBatchProcessor(opts ...BatchOption) *BatchProcessor {
	opt := &batchOption{}
	for _, o := range opts {
		o(opt)
	}
	tk := time.NewTicker(time.Duration(opt.batchInterval) * time.Second)
	return &BatchProcessor{
		db:          opt.db,
		dataCh:      make(chan map[string]interface{}),
		stopCh:      make(chan interface{}),
		data:        make([]map[string]interface{}, 0),
		batchSize:   opt.batchSize,
		model:       opt.model,
		active:      true,
		batchTicker: tk,
	}
}

func (b *BatchProcessor) Add(item map[string]interface{}) error {
	if !b.isActive() {
		return errors.New("add failed, batch processor is closed.")
	}
	b.dataCh <- item
	return nil
}

func (b *BatchProcessor) Start() {
	for {
		select {
		case item := <-b.dataCh:
			b.data = append(b.data, item)
			if len(b.data) >= b.batchSize {
				b.db.Model(b.model).CreateInBatches(b.data, b.batchSize)
				b.data = b.data[:0]
			}
		case <-b.batchTicker.C:
			b.db.Model(b.model).CreateInBatches(b.data, len(b.data))
			b.data = b.data[:0]
		case <-b.stopCh:
			fmt.Println("stop ...")
			return
		}
	}
}

func (b *BatchProcessor) Stop() {
	b.active = false
	for {
		if len(b.data) == 0 {
			break
		}
		time.Sleep(2 * time.Second)
	}
	close(b.stopCh)
	close(b.dataCh)
}

func (b *BatchProcessor) isActive() bool {
	return b.active
}
