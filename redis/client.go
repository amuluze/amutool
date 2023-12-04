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

type Client struct {
	redis.UniversalClient
}

func NewClient(opts ...Option) (*Client, error) {
	conf := &option{
		Addrs:                 []string{"localhost:6379"},
		Password:              "123456",
		DB:                    0,
		PoolSize:              50,
		MasterName:            "",
		DialConnectionTimeout: "5s",
		DialReadTimeout:       "3s",
		DialWriteTimeout:      "3s",
		IdleTimeout:           "5s",
	}
	for _, opt := range opts {
		opt(conf)
	}

	dailTimeout, err := time.ParseDuration(conf.DialConnectionTimeout)
	if err != nil {
		return nil, err
	}
	readTimeout, err := time.ParseDuration(conf.DialReadTimeout)
	if err != nil {
		return nil, err
	}
	writeTimeout, err := time.ParseDuration(conf.DialWriteTimeout)
	if err != nil {
		return nil, err
	}
	idleTimeout, err := time.ParseDuration(conf.IdleTimeout)
	if err != nil {
		return nil, err
	}
	c := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:        conf.Addrs,
		DB:           conf.DB,
		Password:     conf.Password,
		PoolSize:     conf.PoolSize,
		MasterName:   conf.MasterName,
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

func (rc *Client) Close() {
	rc.Close()
}
