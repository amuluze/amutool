// Package gormx
// Date: 2022/9/27 17:11
// Author: Amu
// Description:
package gormx

type Config struct {
	Debug        bool
	DBType       string
	DSN          string
	TablePrefix  string
	MaxLifetime  int
	MaxOpenConns int
	MaxIdleConns int
}

type Option func(*Config)

func SetDebug(debug bool) Option {
	return func(config *Config) {
		config.Debug = debug
	}
}

func SetDBType(dbType string) Option {
	return func(config *Config) {
		config.DBType = dbType
	}
}

func SetDSN(dsn string) Option {
	return func(config *Config) {
		config.DSN = dsn
	}
}

func SetTablePrefix(prefix string) Option {
	return func(config *Config) {
		config.TablePrefix = prefix
	}
}

func SetMaxLifetime(lifetime int) Option {
	return func(config *Config) {
		config.MaxLifetime = lifetime
	}
}

func SetMaxOpenConns(openConns int) Option {
	return func(config *Config) {
		config.MaxOpenConns = openConns
	}
}

func SetMaxIdleConns(idleConns int) Option {
	return func(config *Config) {
		config.MaxIdleConns = idleConns
	}
}
