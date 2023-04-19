// Package clickhousex
// Date: 2023/4/7 18:09
// Author: Amu
// Description:
package clickhousex

import (
	"fmt"
	"time"

	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/jmoiron/sqlx"
)

type Config struct {
	Addr            string // 服务地址
	Username        string // 用户名
	Password        string // 密码，没有则为空
	Database        string // 使用数据库
	DialTimeout     string // 读取超时，默认 3s，-1 表示取消读超时
	MaxOpenConns    int    // 最大连接数
	MaxIdleConns    int    // 最大空闲连接数
	ConnMaxLifeTime string // 连接存活时间
}

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

	pool, err := sqlx.Open("clickhousex", source)
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
