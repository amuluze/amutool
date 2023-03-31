// Package es
// Date: 2023/3/24 18:00
// Author: Amu
// Description:
package es

import (
	"context"
	"time"

	"github.com/olivere/elastic/v7"
)

type BulkProcessor struct {
	*elastic.BulkProcessor
}

type BulkRequest struct {
	Doc    interface{}
	Index  string
	Action string
}

type BulkStats struct {
	Flushed   int64 // number of times the flush interval has been invoked
	Committed int64 // # of times workers committed bulk requests
	Indexed   int64 // # of requests indexed
	Created   int64 // # of requests that ES reported as creates (201)
	Updated   int64 // # of requests that ES reported as updates
	Deleted   int64 // # of requests that ES reported as deletes
	Succeeded int64 // # of requests that ES reported as successful
	Failed    int64 // # of requests that ES reported as failed
}

func NewBulkProcess(client *Client, config *Config) (*BulkProcessor, error) {
	cfg := config.Elastic
	service := elastic.NewBulkProcessorService(client.Client).
		FlushInterval(time.Duration(cfg.BulkFlushInterval) * time.Second).
		Workers(cfg.BulkWorkers).
		BulkActions(cfg.BulkActions).
		BulkSize(cfg.BulkSize * 1024 * 1024).
		Stats(true)

	processor, err := service.Do(context.Background())
	if err != nil {
		return nil, err
	}
	return &BulkProcessor{processor}, err
}

func (p *BulkProcessor) Close() error {
	return p.BulkProcessor.Close()
}

func (p *BulkProcessor) Start(ctx context.Context) error {
	return p.BulkProcessor.Start(ctx)
}

func (p *BulkProcessor) Stop() error {
	return p.BulkProcessor.Stop()
}

func (p *BulkProcessor) Stats() BulkStats {
	temp := p.BulkProcessor.Stats()
	return BulkStats{
		Flushed:   temp.Flushed,
		Committed: temp.Committed,
		Indexed:   temp.Indexed,
		Created:   temp.Created,
		Updated:   temp.Updated,
		Deleted:   temp.Deleted,
		Succeeded: temp.Succeeded,
		Failed:    temp.Failed,
	}
}

func (p *BulkProcessor) Add(request BulkRequest) error {
	req, err := p.buildRequest(request)
	if err != nil {
		return err
	}
	p.BulkProcessor.Add(req)
	return nil
}

func (p *BulkProcessor) Flush() error {
	return p.BulkProcessor.Flush()
}

func (p *BulkProcessor) buildRequest(request BulkRequest) (elastic.BulkableRequest, error) {
	var req elastic.BulkableRequest
	switch request.Action {
	case "create":
		req = elastic.NewBulkCreateRequest().UseEasyJSON(true).Index(request.Index).Doc(request.Doc)
	case "update":
		req = elastic.NewBulkUpdateRequest().UseEasyJSON(true).Index(request.Index).Doc(request.Doc)
	case "delete":
		req = elastic.NewBulkDeleteRequest().UseEasyJSON(true).Index(request.Index)
	case "index":
		req = elastic.NewBulkIndexRequest().UseEasyJSON(true).Index(request.Index).OpType(request.Action).Doc(request.Doc)
	}
	return req, nil
}
