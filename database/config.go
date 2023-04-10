// Package database
// Date: 2023/4/4 14:04
// Author: Amu
// Description:
package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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
