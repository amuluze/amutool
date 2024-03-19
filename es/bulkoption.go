// Package es
// Date: 2023/12/1 15:43
// Author: Amu
// Description:
package es

type BulkOption func(*bulkOption)

type bulkOption struct {
	BulkWorkers       int
	BulkActions       int
	BulkSize          int
	BulkFlushInterval string
}

func WithBulkWorkers(workers int) BulkOption {
	return func(option *bulkOption) {
		option.BulkWorkers = workers
	}
}

func WithBulkActions(actions int) BulkOption {
	return func(option *bulkOption) {
		option.BulkActions = actions
	}
}

func WithBulkSize(size int) BulkOption {
	return func(option *bulkOption) {
		option.BulkSize = size
	}
}

func WithBulkFlushInterval(interval string) BulkOption {
	return func(option *bulkOption) {
		option.BulkFlushInterval = interval
	}
}
