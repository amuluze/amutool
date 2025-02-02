// Package stringx
// Date: 2023/4/10 14:38
// Author: Amu
// Description:
package stringx

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRandn(t *testing.T) {
	Seed(time.Now().UnixNano())
	assert.True(t, len(Rand()) > 0)
	assert.True(t, len(RandId()) > 0)

	const size = 10
	assert.True(t, len(Randn(size)) == size)
}

func BenchmarkRandString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Randn(10)
	}
}
