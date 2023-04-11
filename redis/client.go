// Package redis
// Date: 2022/9/23 00:51
// Author: Amu
// Description:
package redis

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type Config struct {
	Redis Redis
}

type Redis struct {
	Addrs                 []string // redis 地址，兼容单机和集群
	Password              string   // 密码，没有则为空
	DB                    int      // 使用数据库
	PoolSize              int      // 连接池大小
	MasterName            string   // 有值，则为哨兵模式
	DialConnectionTimeout string   // 连接超时，默认 5s
	DialReadTimeout       string   // 读取超时，默认 3s，-1 表示取消读超时
	DialWriteTimeout      string   // 写入超时，默认 3s， -1 表示取消写超时
	IdleTimeout           string   // 空闲连接超时，默认 5m，-1 表示取消闲置超时
}

type Client struct {
	redis.UniversalClient
}

func NewClient(config *Config) (*Client, error) {
	cfg := config.Redis
	dailTimeout, err := time.ParseDuration(cfg.DialConnectionTimeout)
	if err != nil {
		return nil, err
	}
	readTimeout, err := time.ParseDuration(cfg.DialReadTimeout)
	if err != nil {
		return nil, err
	}
	writeTimeout, err := time.ParseDuration(cfg.DialWriteTimeout)
	if err != nil {
		return nil, err
	}
	idleTimeout, err := time.ParseDuration(cfg.IdleTimeout)
	if err != nil {
		return nil, err
	}
	c := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:        cfg.Addrs,
		DB:           cfg.DB,
		Password:     cfg.Password,
		PoolSize:     cfg.PoolSize,
		MasterName:   cfg.MasterName,
		DialTimeout:  dailTimeout,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		IdleTimeout:  idleTimeout,
	})
	_, err = c.Ping(ctx).Result()
	if err != nil {
		log.Fatalln("Failed to create elastic client")
		return nil, err
	}
	return &Client{c}, nil
}
