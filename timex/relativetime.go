// Package timex
// Date: 2023/4/10 13:07
// Author: Amu
// Description:
package timex

import "time"

var initTime = time.Now().AddDate(-1, -1, -1)

func Now() time.Duration {
	return time.Since(initTime)
}

func Since(d time.Duration) time.Duration {
	return time.Since(initTime) - d
}

// Time current time
func Time() time.Time {
	return initTime.Add(Now())
}
