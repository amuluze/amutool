// Package es
// Date: 2023/3/24 18:00
// Author: Amu
// Description:
package es

import (
	"context"
	"time"

	"gitee.com/amuluze/amutool/uuidx"

	"gitee.com/amuluze/amutool/logx"
	"github.com/olivere/elastic/v7"
)

// BulkProcessor 增删改批量操作
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

func NewBulkProcessor(client *Client, opts ...BulkOption) (*BulkProcessor, error) {
	opt := &bulkOption{}
	for _, o := range opts {
		o(opt)
	}
	interval, _ := time.ParseDuration(opt.BulkFlushInterval)
	service := elastic.NewBulkProcessorService(client.Client).
		FlushInterval(interval).
		Workers(opt.BulkWorkers).
		BulkActions(opt.BulkActions).
		BulkSize(opt.BulkSize * 1024 * 1024).
		Stats(true)

	processor, err := service.After(after).Do(context.Background())
	if err != nil {
		return nil, err
	}
	return &BulkProcessor{processor}, err
}

func after(executionID int64, requests []elastic.BulkableRequest, response *elastic.BulkResponse, err error) {
	if err != nil {
		logx.Errorf("bulk commit failed, err: %v\n", err)
	}
	// do what ever you want in case bulk commit success
	logx.Infof("commit successfully, len(requests): %d, execution id: %d, response: %v\n", len(requests), executionID, response)
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

func (p *BulkProcessor) Flush() error {
	return p.BulkProcessor.Flush()
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

func (p *BulkProcessor) Create(indexName string, doc interface{}) {
	cr := elastic.NewBulkCreateRequest().Index(indexName).Id(uuidx.MustString()).Doc(doc)
	p.BulkProcessor.Add(cr)
}

func (p *BulkProcessor) Update(indexName string, docID string, doc interface{}) {
	ur := elastic.NewBulkUpdateRequest().Index(indexName).Id(docID).Doc(doc)
	p.BulkProcessor.Add(ur)
}

func (p *BulkProcessor) Delete(indexName string, docIDs []string) {
	for _, docID := range docIDs {
		dr := elastic.NewBulkDeleteRequest().Index(indexName).Id(docID)
		p.BulkProcessor.Add(dr)
	}
}
