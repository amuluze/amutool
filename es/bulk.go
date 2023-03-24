// Package es
// Date: 2023/3/24 18:00
// Author: Amu
// Description:
package es

import (
	"context"

	"github.com/olivere/elastic/v7"
)

func NewBulkProcess(client *elastic.Client, workNum int, bulkActions int, bulkSize int) (*elastic.BulkProcessor, error) {
	bulkProcess, err := client.BulkProcessor().
		Workers(workNum).
		BulkActions(bulkActions).
		BulkSize(bulkSize).
		Do(context.Background())

	if err != nil {
		return nil, err
	}
	err = bulkProcess.Start(context.Background())
	return bulkProcess, err
}

func Close(bulkProcess *elastic.BulkProcessor) error {
	return bulkProcess.Close()
}
