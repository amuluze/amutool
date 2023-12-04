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

var temp = `{
    "index_patterns": ["users-*"],
    "priority": 1,
    "template": {
    "settings": {
        "number_of_shards": 1,
        "number_of_replicas": 1,
        "index.lifecycle.name": "users",
        "index.lifecycle.rollover_alias": "users",
        "index.lifecycle.parse_origination_date": true
    },
    "mappings": {
        "_source": { "enabled": true },
        "properties": {
        "username": {
            "type": "text",
            "fields": {
            "keyword": {
                "type": "keyword",
                "ignore_above": 256
            }
            },
            "norms": false
        },
        "age": {
            "type": "long"
        },
        "sex": {
            "type": "long"
        }
        }
    }
    }
}`

func TestTemplateCreate(t *testing.T) {
	client, _ := NewEsClient(
		WithAddr("http://localhost:9200"),
		WithUsername("elastic"),
		WithPassword("123456"),
		WithSniff(true),
		WithDebug(true),
		WithHealthcheck(true),
	)
	template, err := client.PutIndexTemplate(context.TODO(), indexName, temp)
	if err != nil {
		panic(err)
	}
	fmt.Printf("template %s create: %v\n", indexName, template)
}

func TestTemplateExists(t *testing.T) {
	client, _ := NewEsClient(
		WithAddr("http://localhost:9200"),
		WithUsername("elastic"),
		WithPassword("123456"),
		WithSniff(true),
		WithDebug(true),
		WithHealthcheck(true),
	)
	exists, err := client.TemplateExists(context.TODO(), indexName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("template %s exists: %v\n", indexName, exists)
}

func TestTemplateGet(t *testing.T) {
	client, _ := NewEsClient(
		WithAddr("http://localhost:9200"),
		WithUsername("elastic"),
		WithPassword("123456"),
		WithSniff(true),
		WithDebug(true),
		WithHealthcheck(true),
	)
	template, err := client.GetIndexTemplate(context.TODO(), indexName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("template %s settings: %v\n", indexName, template.Settings)
	fmt.Printf("template %s mappings: %v\n", indexName, template.Mappings)
	fmt.Printf("template %s aliases: %v\n", indexName, template.Aliases)
}

func TestTemplateDelete(t *testing.T) {
	client, _ := NewEsClient(
		WithAddr("http://localhost:9200"),
		WithUsername("elastic"),
		WithPassword("123456"),
		WithSniff(true),
		WithDebug(true),
		WithHealthcheck(true),
	)
	template, err := client.DeleteIndexTemplate(context.TODO(), indexName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("template %s delete: %v", indexName, template)
}
