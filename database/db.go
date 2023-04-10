// Package database
// Date: 2023/4/4 13:57
// Author: Amu
// Description:
package database

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func NewDB(cfg *Config) (*DB, error) {
	dial := cfg.Dial()
	db, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 开启调试模式
	if cfg.Debug {
		db = db.Debug()
	}

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(cfg.MaxLifetime))

	return &DB{db}, nil
}

func (db *DB) Close() {
	if db != nil {
		db.Close()
	}
}

func (db *DB) AutoMigrate(models ...interface{}) error {
	return db.AutoMigrate(models...)
}

type Option func(db *DB) *DB

func OptionDB(db *DB, options ...Option) *DB {
	for _, option := range options {
		db = option(db)
	}
	return db
}

func WithInIds(ids ...int64) Option {
	if len(ids) == 0 {
		return nil
	}
	return func(db *DB) *DB {
		db.Where("id IN (?)", ids)
		return db
	}
}

func WithById(id int64) Option {
	if id <= 0 {
		return nil
	}
	return func(db *DB) *DB {
		db.Where("id = ?", id)
		return db
	}
}

func WithOffset(offset int) Option {
	if offset < 0 {
		return nil
	}
	return func(db *DB) *DB {
		db.Offset(offset)
		return db
	}
}

func WithLimit(limit int) Option {
	if limit <= 0 {
		return nil
	}
	return func(db *DB) *DB {
		db.Limit(limit)
		return db
	}
}

func OrderBy(value interface{}) Option {
	return func(db *DB) *DB {
		db.Order(value)
		return db
	}
}
