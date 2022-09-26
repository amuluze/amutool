// Package redis
// Date: 2022/9/23 00:51
// Author: Amu
// Description:
package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type Config struct {
	Addr                  []string      `toml:"addr"`               // redis 地址，兼容单机和集群
	Password              string        `toml:"password"`           // 密码，没有则为空
	DB                    int           `toml:"db"`                 // 使用数据库
	MasterName            string        `toml:"master_name"`        // 有值，则为哨兵模式
	DialConnectionTimeout time.Duration `toml:"connection_timeout"` // 连接超时，默认 5s
	DialReadTimeout       time.Duration `toml:"read_timeout"`       // 读取超时，默认 3s，-1 表示取消读超时
	DialWriteTimeout      time.Duration `toml:"write_timeout"`      // 写入超时，默认 3s， -1 表示取消写超时
	IdleTimeout           time.Duration `toml:"idle_time"`          // 空闲连接超时，默认 5m，-1 表示取消闲置超时
}

func NewClient(config *Config) redis.UniversalClient {

	c := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:        config.Addr,
		DB:           config.DB,
		Password:     config.Password,
		MasterName:   config.MasterName,
		DialTimeout:  config.DialConnectionTimeout,
		ReadTimeout:  config.DialReadTimeout,
		WriteTimeout: config.DialWriteTimeout,
		IdleTimeout:  config.IdleTimeout,
	})
	return c
}
