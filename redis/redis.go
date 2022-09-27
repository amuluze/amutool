// Package redis
// Date: 2022/9/23 00:51
// Author: Amu
// Description:
package redis

import (
	"context"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

var once sync.Once
var ctx = context.Background()

type RedisClient struct {
	redis.UniversalClient
}

func init() {
	once.Do(func() {
		rc = &RedisClient{
			redis.NewUniversalClient(&redis.UniversalOptions{
				Addrs:    []string{"127.0.0.1:6379"},
				DB:       0,
				Password: "Be1s.Az3",
			}),
		}
	})
}

func InitRedis(options ...Option) {
	config := &Config{
		Addr:                  []string{"127.0.0.1:6379"},
		Password:              "Be1s.Az3",
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
	rc = &RedisClient{
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
