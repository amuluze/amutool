// Package es
// Date: 2023/3/24 18:01
// Author: Amu
// Description:
package es

import "github.com/olivere/elastic/v7"

type Model interface {
	GetIndexName() string
	GetIndexConfig() interface{}
	GetId() string
	ExistIndex(string, *elastic.Client) (bool, error)
	SetNowIndex(string)
	GetSortField() string
}
