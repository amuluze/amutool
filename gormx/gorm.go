// Package gormx
// Date: 2022/9/23 14:07
// Author: Amu
// Description:
package gormx

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB(options ...Option) error {
	cfg := &Config{
		Debug:        true,
		Type:         "mysql",
		Host:         "127.0.0.1",
		Port:         "3306",
		UserName:     "root",
		Password:     "123456",
		Name:         "test",
		TablePrefix:  "",
		MaxLifetime:  7200,
		MaxOpenConns: 10,
		MaxIdleConns: 5,
	}

	for _, option := range options {
		option(cfg)
	}

	dial := cfg.Dial()
	conn, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	sqlDB, err := conn.DB()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(cfg.MaxLifetime))
	db = conn
	return nil
}

func GetDB() *gorm.DB {
	return db
}
