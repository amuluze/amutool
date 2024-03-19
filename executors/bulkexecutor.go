// Package executors
// Date: 2023/4/10 11:36
// Author: Amu
// Description:
package executors

import "time"

type bulkWriteOptions struct {
	cachedTask    int
	flushInterval time.Duration
}

type BulkOption func(options *bulkWriteOptions)

type BulkExecutor struct {
}
