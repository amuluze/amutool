// Package clickhousex
// Date: 2023/12/6 18:08
// Author: Amu
// Description:
package clickhousex

type Option func(*option)

type option struct {
	Debug           bool
	Addr            string // 服务地址
	Username        string // 用户名
	Password        string // 密码，没有则为空
	Database        string // 使用数据库
	DialTimeout     int    // 读取超时，单位 s
	ConnMaxLifeTime int    // 连接存活时间 单位 s
	MaxOpenConns    int    // 最大连接数
	MaxIdleConns    int    // 最大空闲连接数
}

func WithDebug(debug bool) Option {
	return func(o *option) {
		o.Debug = debug
	}
}

func WithAddr(addr string) Option {
	return func(o *option) {
		o.Addr = addr
	}
}

func WithUsername(username string) Option {
	return func(o *option) {
		o.Username = username
	}
}

func WithPassword(password string) Option {
	return func(o *option) {
		o.Password = password
	}
}

func WithDatabase(dbName string) Option {
	return func(o *option) {
		o.Database = dbName
	}
}

func WithDialTimeout(dialTimeout int) Option {
	return func(o *option) {
		o.DialTimeout = dialTimeout
	}
}

func WithMaxOpenConns(maxOpen int) Option {
	return func(o *option) {
		o.MaxOpenConns = maxOpen
	}
}

func WithMaxIdleConns(maxIdle int) Option {
	return func(o *option) {
		o.MaxIdleConns = maxIdle
	}
}

func WithConnMaxLifeTime(lifeTime int) Option {
	return func(o *option) {
		o.ConnMaxLifeTime = lifeTime
	}
}
