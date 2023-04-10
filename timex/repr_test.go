// Package timex
// Date: 2023/4/10 13:22
// Author: Amu
// Description:
package timex

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestReprOfDuration(t *testing.T) {
	assert.Equal(t, "1000.0ms", ReprOfDuration(time.Second))
	assert.Equal(t, "1111.6ms", ReprOfDuration(
		time.Second+time.Millisecond*111+time.Microsecond*555))
}
