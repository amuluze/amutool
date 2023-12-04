// Package es
// Date: 2023/12/1 17:02
// Author: Amu
// Description:
package es

import (
	"context"
	"fmt"
	"testing"
)

func TestGetQueryService(t *testing.T) {
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
	qService := client.GetQueryService()
	do, err := qService.Index(indexName).Do(context.TODO())
	if err != nil {
		return
	}
	fmt.Printf("query service: %#v", do)
}
