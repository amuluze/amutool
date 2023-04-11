// Package clickhouse
// Date: 2023/4/7 18:09
// Author: Amu
// Description:
package clickhouse

import (
	"fmt"
	"time"

	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/jmoiron/sqlx"
)

type Client struct {
	*sqlx.DB
}

func NewClient(config *Config) (*Client, error) {
	source := fmt.Sprintf("tcp://%s?debug=true&username=%s&password=%s&database=%s", config.Addr, config.Username, config.Password, config.Database)

	dialTimeout, err := time.ParseDuration(config.DialTimeout)
	if err != nil {
		return nil, err
	}
	maxLifeTime, err := time.ParseDuration(config.ConnMaxLifeTime)
	if err != nil {
		return nil, err
	}

	pool, err := sqlx.Open("clickhouse", source)
	if err != nil {
		return nil, err
	}
	pool.SetConnMaxLifetime(maxLifeTime)
	pool.SetConnMaxIdleTime(dialTimeout)
	pool.SetMaxIdleConns(config.MaxIdleConns)
	pool.SetMaxOpenConns(config.MaxOpenConns)

	return &Client{pool}, nil
}

func (c *Client) Close() {
	c.Close()
}

func (c *Client) GetTx() (*Tx, error) {
	tx, err := c.Beginx()
	if err != nil {
		return nil, err
	}
	return &Tx{tx}, err
}
