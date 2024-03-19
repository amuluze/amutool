// Package envx
// Date: 2022/9/7 10:30
// Author: Amu
// Description:
package envx

import (
	"fmt"
	"testing"
)

func TestGetEnv(t *testing.T) {
	key := "GOROOT"
	res := GetEnv(key)
	fmt.Println(res)
}
