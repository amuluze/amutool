// Package redis
// Date: 2023/12/4 14:02
// Author: Amu
// Description:
package redis

import (
	"fmt"
	"testing"
)

func TestClient(t *testing.T) {
	rc, err := NewClient()
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println("redis client: ", rc)
}
