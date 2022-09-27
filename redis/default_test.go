// Package redis
// Date: 2022/9/26 17:56
// Author: Amu
// Description:
package redis

import (
	"fmt"
	"testing"
)

func TestKeys(t *testing.T) {
	fmt.Println(&rc)
	keys, err := Keys()
	if err != nil {
		panic(err)
	}
	fmt.Println(keys)
}

func TestType(t *testing.T) {
	fmt.Println(&rc)
}
