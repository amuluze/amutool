// Package clickhousex
// Date: 2023/12/6 18:28
// Author: Amu
// Description:
package clickhousex

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
)

func getConn(opt *option) *sql.DB {
	conn := clickhouse.OpenDB(&clickhouse.Options{
		Addr: []string{opt.Addr},
		Auth: clickhouse.Auth{
			Database: opt.Database,
			Username: opt.Username,
			Password: opt.Password,
		},
		Debug: opt.Debug,
		Debugf: func(format string, v ...any) {
			fmt.Printf(format, v)
		},
		Settings: clickhouse.Settings{
			"max_execution_time": 60,
		},
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		},
		DialTimeout: time.Duration(opt.DialTimeout) * time.Second,
		//ConnOpenStrategy: clickhouse.ConnOpenInOrder,
		//BlockBufferSize:  10,
	})
	conn.SetMaxOpenConns(opt.MaxOpenConns)
	conn.SetConnMaxLifetime(time.Duration(opt.ConnMaxLifeTime) * time.Second)
	conn.SetMaxIdleConns(opt.MaxIdleConns)

	return conn
}
