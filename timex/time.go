// Package timex
// Date: 2023/4/19 16:49
// Author: Amu
// Description:
package timex

import "time"

// Int64ToTime 秒值时间戳转 time.Time
func Int64ToTime(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}
