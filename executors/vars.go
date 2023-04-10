// Package executors
// Date: 2023/4/10 11:34
// Author: Amu
// Description:
package executors

import "time"

const defaultFlushInterval = time.Second

// Execute defines the method to execute tasks.
type Execute func(task []interface{})
