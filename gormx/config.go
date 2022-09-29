// Package gormx
// Date: 2022/9/27 17:11
// Author: Amu
// Description:
package gormx

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"gorm.io/driver/mysql"
)

type Config struct {
	Debug        bool
	Type         string
	Host         string
	Port         string
	UserName     string
	Password     string
	Name         string
	TablePrefix  string
	MaxLifetime  int
	MaxOpenConns int
	MaxIdleConns int
}

func (c *Config) Dial() gorm.Dialector {
	var dsn string
	var dialector gorm.Dialector
	if c.Type == "mysql" {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			c.UserName,
			c.Password,
			c.Host,
			c.Port,
			c.Name,
		)
		dialector = mysql.New(mysql.Config{
			DSN:                       dsn,
			DefaultStringSize:         256,   // default size for string fields
			DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
			DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
			DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
			SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
		})
	} else if c.Type == "postgres" {
		dsn = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable TimeZone=Asia/Shanghai",
			c.Host,
			c.Port,
			c.UserName,
			c.Name,
			c.Password,
		)
		dialector = postgres.New(postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true,
		})
	} else {
		dsn = fmt.Sprintf("%s.db", c.Name)
		dialector = sqlite.Open(dsn)
	}
	return dialector
}

type Option func(*Config)

func SetDebug(debug bool) Option {
	return func(config *Config) {
		config.Debug = debug
	}
}

func SetDBType(dbType string) Option {
	return func(config *Config) {
		config.Type = dbType
	}
}

func SetHost(host string) Option {
	return func(config *Config) {
		config.Host = host
	}
}

func SetPort(port string) Option {
	return func(config *Config) {
		config.Port = port
	}
}

func SetUserName(username string) Option {
	return func(config *Config) {
		config.UserName = username
	}
}

func SetPassword(pass string) Option {
	return func(config *Config) {
		config.Password = pass
	}
}

func SetName(name string) Option {
	return func(config *Config) {
		config.Name = name
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
