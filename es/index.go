// Package es
// Date: 2023/3/27 15:50
// Author: Amu
// Description:
package es

import (
	"context"
	"fmt"
)

const (
	CreateIndexRetry = 50
)

func (c *Client) IndexExists(ctx context.Context, indexName string) (bool, error) {
	res, err := c.IndexGet(indexName).Human(true).Do(ctx)
	if err != nil {
		return false, err
	}
	fmt.Println(res)
	return true, nil
}

func (c *Client) CreateIndex(ctx context.Context, indexName string, indexBody string) (bool, error) {
	res, err := c.Client.CreateIndex(indexName).BodyString(indexBody).Do(ctx)
	if err != nil {
		return false, err
	}
	fmt.Println(res)
	return true, nil
}

func (c *Client) DeleteIndex(ctx context.Context, indexName string) error {
	res, err := c.Client.DeleteIndex(indexName).Do(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("delete index response: %v\n", res)
	return nil
}

func (c *Client) GetSettings(ctx context.Context, indexName string) (map[string]interface{}, error) {
	res, err := c.Client.IndexGetSettings(indexName).Human(true).Pretty(true).Do(ctx)
	if err != nil {
		return nil, err
	}
	return res[indexName].Settings, nil
}

func (c *Client) GetMappings(ctx context.Context, indexName string) (map[string]interface{}, error) {
	res, err := c.Client.GetMapping().Index(indexName).Human(true).Pretty(true).Do(ctx)
	if err != nil {
		return nil, err
	}
	//fmt.Printf("index mappings response: %#v", res[indexName])
	return (res[indexName].(map[string]interface{}))["mappings"].(map[string]interface{}), nil
}
