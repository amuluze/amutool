// Package es
// Date: 2023/12/1 16:05
// Author: Amu
// Description:
package es

import (
	"fmt"
	"testing"
)

func TestClient(t *testing.T) {
	client, err := NewEsClient(
		WithAddr("http://localhost:9200"),
		WithUsername("elastic"),
		WithPassword("123456"),
		WithSniff(true),
		WithDebug(true),
		WithHealthcheck(true),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(client)
}
