// Package redis
// Date: 2023/12/4 12:12
// Author: Amu
// Description:
package redis

type Option func(*option)

type option struct {
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

func WithAddrs(addrs []string) Option {
	return func(o *option) {
		o.Addrs = addrs
	}
}

func WithPassword(password string) Option {
	return func(o *option) {
		o.Password = password
	}
}

func WithDB(db int) Option {
	return func(o *option) {
		o.DB = db
	}
}

func WithPoolSize(poolSize int) Option {
	return func(o *option) {
		o.PoolSize = poolSize
	}
}

func WithMasterName(masterName string) Option {
	return func(o *option) {
		o.MasterName = masterName
	}
}

func WithConnectionTimeout(connectionTimeout string) Option {
	return func(o *option) {
		o.DialConnectionTimeout = connectionTimeout
	}
}

func WithReadTimeout(readTimeout string) Option {
	return func(o *option) {
		o.DialReadTimeout = readTimeout
	}
}

func WithWriteTimeout(writeTimeout string) Option {
	return func(o *option) {
		o.DialWriteTimeout = writeTimeout
	}
}

func WithIdleTimeout(idleTimeout string) Option {
	return func(o *option) {
		o.IdleTimeout = idleTimeout
	}
}
