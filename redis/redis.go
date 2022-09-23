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

type Config struct {
	Addr                  []string      `toml:"addr"`               // redis 地址，兼容单机和集群
	Password              string        `toml:"password"`           // 密码，没有则为空
	DB                    int           `toml:"db"`                 // 使用数据库
	MasterName            string        `toml:"master_name"`        // 有值，则为哨兵模式
	DialConnectionTimeout time.Duration `toml:"connection_timeout"` // 连接超时
	DialReadTimeout       time.Duration `toml:"read_timeout"`       // 读取超时
	DialWriteTimeout      time.Duration `toml:"write_timeout"`      // 写入超时
	IdleTimeout           time.Duration `toml:"idle_time"`          // 空闲连接超时
}

func MustBootUp(config Config) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

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

	_, err := c.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}
