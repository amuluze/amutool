// Package clickhousex
// Date: 2023/12/7 10:33
// Author: Amu
// Description:
package clickhousex

type BatchOption func(*batchOption)

type batchOption struct {
	db            *DB
	batchSize     int
	batchInterval int
	model         interface{}
}

func WithDB(db *DB) BatchOption {
	return func(b *batchOption) {
		b.db = db
	}
}

func WithBatchSize(size int) BatchOption {
	return func(b *batchOption) {
		b.batchSize = size
	}
}

func WithBatchInterval(interval int) BatchOption {
	return func(b *batchOption) {
		b.batchInterval = interval
	}
}

func WithModel(model interface{}) BatchOption {
	return func(b *batchOption) {
		b.model = model
	}
}
