// Package es
// Date: 2023/3/30 12:07
// Author: Amu
// Description: client crud
package es

import (
	"context"

	"github.com/olivere/elastic/v7"
)

// ================================================= 查找 =================================================

func (c *Client) Find(ctx context.Context, bulkModel Model, options ...QueryOption) (*elastic.SearchResult, error) {
	return c.GetQueryService(options...).Index(bulkModel.GetIndexName()).Do(ctx)
}

func (c *Client) FindByIndexName(ctx context.Context, indexNames []string, options ...QueryOption) (*elastic.SearchResult, error) {
	return c.GetQueryService(options...).Index(indexNames...).Do(ctx)
}

func (c *Client) FindScroll(ctx context.Context, bulkModel Model, scrollID string, size int, options ...ScrollQueryOption) (*elastic.SearchResult, error) {
	if scrollID == "" {
		return c.GetScrollQueryService(options...).Scroll("5m").Size(size).Index(bulkModel.GetIndexName()).Do(ctx)
	}
	return c.GetScrollQueryService(options...).ScrollId(scrollID).Scroll("5m").Size(size).Index(bulkModel.GetIndexName()).Do(ctx)
}

func (c *Client) FindByIndexNameScroll(ctx context.Context, indexNames []string, scrollID string, size int, options ...ScrollQueryOption) (*elastic.SearchResult, error) {
	if scrollID == "" {
		return c.GetScrollQueryService(options...).Scroll("5m").Size(size).Index(indexNames...).Do(ctx)
	}
	return c.GetScrollQueryService(options...).ScrollId(scrollID).Scroll("5m").Size(size).Index(indexNames...).Do(ctx)
}

func (c *Client) CountByIndices(ctx context.Context, indexNames []string, options ...CountOption) (int64, error) {
	return c.GetCountService(options...).Index(indexNames...).Do(ctx)
}
