// Package timex
// Date: 2023/4/19 16:50
// Author: Amu
// Description:
package timex

import (
	"fmt"
	"testing"
	"time"
)

func TestInt64ToTime(t *testing.T) {
	timestamp := time.Now().Unix()
	res := Int64ToTime(timestamp)
	fmt.Println(res)
}
