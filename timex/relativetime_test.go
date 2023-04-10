// Package timex
// Date: 2023/4/10 13:08
// Author: Amu
// Description:
package timex

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRelativeTime(t *testing.T) {
	fmt.Printf("now: %v\n", time.Now())
	fmt.Printf("init time: %v\n", initTime)
	fmt.Printf("time: %v\n", Time())
	time.Sleep(time.Millisecond)
	now := Now()
	assert.True(t, now > 0)
	time.Sleep(time.Millisecond)
	assert.True(t, Since(now) > 0)
}

func BenchmarkTimeSince(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = time.Since(time.Now())
	}
}

func BenchmarkTimexSince(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = Since(Now())
	}
}
