// Package clickhousex
// Date: 2023/12/6 18:08
// Author: Amu
// Description:
package clickhousex

import (
	"fmt"

	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func NewDB(opts ...Option) (*DB, error) {
	opt := &option{
		Debug:           false,
		MaxIdleConns:    50,
		MaxOpenConns:    100,
		DialTimeout:     10,
		ConnMaxLifeTime: 3600,
	}
	for _, o := range opts {
		o(opt)
	}
	conn := getConn(opt)
	fmt.Printf("conn: %#v\n", conn)

	db, err := gorm.Open(clickhouse.New(clickhouse.Config{
		Conn:                         conn, // initialize with existing database conn
		DisableDatetimePrecision:     true,
		DontSupportRenameColumn:      true,
		DontSupportEmptyDefaultValue: false,
		SkipInitializeWithVersion:    false,
		DefaultGranularity:           3,
		DefaultCompression:           "LZ4",
		DefaultIndexType:             "minmax",
		DefaultTableEngineOpts:       "ENGINE=MergeTree() ORDER BY tuple()",
	}), &gorm.Config{})
	return &DB{db}, err
}

func (db *DB) Close() {
	if db != nil {
		conn, err := db.DB.DB()
		if err != nil {
			return
		}
		err = conn.Close()
		if err != nil {
			return
		}
	}
}

func (db *DB) RunInTransaction(fn func(tx *gorm.DB) error) error {
	return db.Transaction(fn)
}

func (db *DB) AutoMigrate(models ...interface{}) error {
	return db.DB.AutoMigrate(models...)
}
