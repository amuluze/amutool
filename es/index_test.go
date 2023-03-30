// Package es
// Date: 2023/3/28 15:05
// Author: Amu
// Description:
package es

import (
	"context"
	"fmt"
	"testing"
)

func TestIndexExists(t *testing.T) {
	var esClient = getClient()
	exists, err := esClient.IndexExists(context.Background(), "hello")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Println(exists)
}

func TestGetSettings(t *testing.T) {
	var esClient = getClient()
	indexName := "leefs_logs-2023.03.29-00001"
	res, err := esClient.GetSettings(context.Background(), indexName)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

func TestGetMappings(t *testing.T) {
	var esClient = getClient()
	indexName := "leefs_logs-2023.03.29-00001"
	res, err := esClient.GetMappings(context.Background(), indexName)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("-------", res)
}
