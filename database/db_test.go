// Package database
// Date: 2023/11/23 13:54
// Author: Amu
// Description:
package database

import (
	"fmt"
	"log"
	"testing"
)

var config = &Config{
	Debug:        true,
	Type:         "mysql",
	Host:         "localhost",
	Port:         "3306",
	UserName:     "root",
	Password:     "amcation",
	DBName:       "amcation",
	TablePrefix:  "s_",
	MaxLifetime:  7200,
	MaxOpenConns: 100,
	MaxIdleConns: 50,
}

func TestNewDB(t *testing.T) {
	db, err := NewDB(config)
	if err != nil {
		log.Fatalf("new db error: %v", err)
	}
	defer db.Close()
	fmt.Println(db)
}
