// Package gormx
// Date: 2022/9/30 17:34
// Author: Amu
// Description:
package gormx

import (
	"fmt"
	"strings"
)

type OrderDirection int

const (
	OrderByASC OrderDirection = iota + 1
	OrderByDESC
)

type OrderField struct {
	Key       string
	Direction OrderDirection
}
type OrderFieldFunc func(string) string

func ParseOrder(items []*OrderField, handle ...OrderFieldFunc) string {
	orders := make([]string, len(items))

	for i, item := range items {
		key := item.Key
		if len(handle) > 0 {
			key = handle[0](key)
		}

		direction := "ASC"
		if item.Direction == OrderByDESC {
			direction = "DESC"
		}

		orders[i] = fmt.Sprintf("%s %s", key, direction)
	}
	return strings.Join(orders, ",")
}
