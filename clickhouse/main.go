// Package main
// Date: 2024/10/21 16:17:18
// Author: Amu
// Description:
package main

import (
	"database/sql"
	"gorm.io/driver/clickhouse"
	"log/slog"
)

func main() {
	conn, err := sql.Open("clickhouse", "http://127.0.0.1:8123/default")
	if err != nil {
		slog.Error("clickhouse conn failed", "error", err)
	}
	if err := conn.Ping(); err != nil {
		slog.Error("clickhouse ping failed", "error", err)
	}
	clickhouse.Open("http://127.0.0.1:8123/default")
}
