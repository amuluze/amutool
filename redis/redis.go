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

type Client struct {
	redis.UniversalClient
}

var rc redis.UniversalClient
var ctx = context.Background()

func InitRedis(options ...Option) {
	config := &Config{
		Addr:                  []string{"127.0.0.1:6379"},
		Password:              "123456",
		DB:                    0,
		MasterName:            "",
		DialConnectionTimeout: 5 * time.Second,
		DialReadTimeout:       3 * time.Second,
		DialWriteTimeout:      3 * time.Second,
		IdleTimeout:           5 * 60 * time.Second,
	}

	for _, option := range options {
		option(config)
	}
	rc = &Client{
		redis.NewUniversalClient(&redis.UniversalOptions{
			Addrs:        config.Addr,
			DB:           config.DB,
			Password:     config.Password,
			MasterName:   config.MasterName,
			DialTimeout:  config.DialConnectionTimeout,
			ReadTimeout:  config.DialReadTimeout,
			WriteTimeout: config.DialWriteTimeout,
			IdleTimeout:  config.IdleTimeout,
		}),
	}
}
