// Package redis
// Date: 2022/9/27 09:55
// Author: Amu
// Description:
package redis

import (
	"time"
)

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

type Option func(*Config)

func SetAddr(addr []string) Option {
	return func(config *Config) {
		config.Addr = addr
	}
}

func SetPassword(pass string) Option {
	return func(config *Config) {
		config.Password = pass
	}
}

func SetDB(db int) Option {
	return func(config *Config) {
		config.DB = db
	}
}

func SetMasterName(masterName string) Option {
	return func(config *Config) {
		config.MasterName = masterName
	}
}

func SetDialConnectionTimeout(timeout time.Duration) Option {
	return func(config *Config) {
		config.DialConnectionTimeout = timeout
	}
}

func SetReadTimeout(timeout time.Duration) Option {
	return func(config *Config) {
		config.DialReadTimeout = timeout
	}
}

func SetWriteTimeout(timeout time.Duration) Option {
	return func(config *Config) {
		config.DialWriteTimeout = timeout
	}
}

func SetIdleTimeout(timeout time.Duration) Option {
	return func(config *Config) {
		config.IdleTimeout = timeout
	}
}
