// Package clickhousex
// Date: 2023/12/7 11:42
// Author: Amu
// Description:
package clickhousex

import (
	"fmt"
	"testing"
)

func TestConn(t *testing.T) {
	opt := &option{
		Debug:           false,
		Addr:            "localhost:9000",
		Database:        "gorm",
		Username:        "root",
		Password:        "123456",
		MaxIdleConns:    50,
		MaxOpenConns:    100,
		DialTimeout:     10,
		ConnMaxLifeTime: 3600,
	}

	conn := getConn(opt)
	fmt.Printf("conn: %#v\n", conn)
}
