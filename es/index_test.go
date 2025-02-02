// Package es
// Date: 2023/12/1 16:05
// Author: Amu
// Description:
package es

import (
	"context"
	"fmt"
	"testing"
	"time"
)

var currentIndexName = indexName + "-" + time.Now().Format("2006.01.02")

func TestCreateIndex(t *testing.T) {
	client, _ := NewEsClient(
		WithAddr("http://localhost:9200"),
		WithUsername("elastic"),
		WithPassword("123456"),
		WithSniff(true),
		WithDebug(true),
		WithHealthcheck(true),
	)
	fmt.Printf("new index name: %s\n", currentIndexName)
	created, err := client.CreateIndex(context.TODO(), currentIndexName, "")
	if err != nil {
		panic(err)
	}
	fmt.Printf("index %s created: %v\n", currentIndexName, created)
}

func TestIndexExists(t *testing.T) {
	client, _ := NewEsClient(
		WithAddr("http://localhost:9200"),
		WithUsername("elastic"),
		WithPassword("123456"),
		WithSniff(true),
		WithDebug(true),
		WithHealthcheck(true),
	)
	fmt.Printf("new index name: %s\n", currentIndexName)
	exists, err := client.IndexExists(context.TODO(), currentIndexName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("index %s exists: %v\n", currentIndexName, exists)
}

func TestIndexStatus(t *testing.T) {
	client, _ := NewEsClient(
		WithAddr("http://localhost:9200"),
		WithUsername("elastic"),
		WithPassword("123456"),
		WithSniff(true),
		WithDebug(true),
		WithHealthcheck(true),
	)
	status, err := client.IndexStatus(context.TODO(), currentIndexName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("index %s status: %#v\n", currentIndexName, status[currentIndexName])
}

func TestDeleteIndex(t *testing.T) {
	client, _ := NewEsClient(
		WithAddr("http://localhost:9200"),
		WithUsername("elastic"),
		WithPassword("123456"),
		WithSniff(true),
		WithDebug(true),
		WithHealthcheck(true),
	)
	err := client.DeleteIndex(context.TODO(), []string{currentIndexName})
	if err != nil {
		panic(err)
	}
	fmt.Printf("index %s deleted", currentIndexName)
}

func TestGetMappings(t *testing.T) {
	client, _ := NewEsClient(
		WithAddr("http://localhost:9200"),
		WithUsername("elastic"),
		WithPassword("123456"),
		WithSniff(true),
		WithDebug(true),
		WithHealthcheck(true),
	)
	mappings, err := client.GetMappings(context.TODO(), currentIndexName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("index %s mappings: %v\n", currentIndexName, mappings)
}

func TestGetSettings(t *testing.T) {
	client, _ := NewEsClient(
		WithAddr("http://localhost:9200"),
		WithUsername("elastic"),
		WithPassword("123456"),
		WithSniff(true),
		WithDebug(true),
		WithHealthcheck(true),
	)
	settings, err := client.GetSettings(context.TODO(), currentIndexName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("index %s settings: %v\n", currentIndexName, settings)
}
