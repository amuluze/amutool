// Package timex
// Date: 2023/4/10 13:07
// Author: Amu
// Description:
package timex

import (
	"fmt"
	"time"
)

// ReprOfDuration returns the string representation of given duration in ms.
func ReprOfDuration(duration time.Duration) string {
	return fmt.Sprintf("%.1fms", float32(duration)/float32(time.Millisecond))
}
