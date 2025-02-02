// Package es
// Date: 2023/12/1 16:06
// Author: Amu
// Description:
package es

import (
	"context"
	"fmt"
	"testing"
)

var indexName = "users"
var ilm = `{
	"policy": {
		"phases": {
			"hot": {
				"min_age": "0ms",
				"actions": {
					"rollover": {
						"max_age": "1d"
					},
					"set_priority": {
						"priority": 100
					}
				}
			},
			"warm": {
				"min_age": "1d",
				"actions": {
					"set_priority": {
						"priority": 50
					}
				}
			},
			"cold": {
				"min_age": "30d",
				"actions": {
					"set_priority": {
						"priority": 0
					},
					"freeze": {}
				}
			},
			"delete": {
				"min_age": "90d",
				"actions": {
					"delete": {}
				}
			}
		}
	}
}`

func TestILMPolicyCreate(t *testing.T) {
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
	created, err := client.PutILMPolicy(context.TODO(), indexName, ilm)
	if err != nil {
		fmt.Printf("create ilm policy error: %v\n", err)
	}
	fmt.Printf("create policy %s: %v\n", indexName, created)
}

func TestILMPolicyExists(t *testing.T) {
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

	exists, err := client.ILMPolicyExists(context.TODO(), indexName)
	if err != nil {
		fmt.Printf("search iml policy exists error: %v\n", err)
	}
	fmt.Printf("%s policy exists %v\n", indexName, exists)
}

func TestILMPolicyGet(t *testing.T) {
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
	ilm, err := client.GetILMPolicy(context.TODO(), indexName)
	if err != nil {
		fmt.Printf("get ilm policy error: %v\n", err)
	}
	fmt.Println(ilm)
}

func TestILMPolicyDelete(t *testing.T) {
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
	res, err := client.DeleteILMPolicy(context.TODO(), indexName)
	if err != nil {
		fmt.Printf("delete ilm policy error: %v\n", err)
	}
	fmt.Printf("delete %s ilm policy: %v\n", indexName, res)
}
