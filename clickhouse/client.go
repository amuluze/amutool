// Package clickhouse
// Date: 2023/4/7 18:09
// Author: Amu
// Description:
package clickhouse

import (
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
)

func DefaultConn() (clickhouse.Conn, error) {
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{"localhost:9000"},
		Auth: clickhouse.Auth{
			Database: "test",
			Username: "root",
			Password: "123456",
		},
		DialTimeout:     30 * time.Second,
		MaxOpenConns:    16,
		MaxIdleConns:    8,
		ConnMaxLifetime: time.Hour,
	})
	return conn, err
}
