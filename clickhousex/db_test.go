// Package clickhousex
// Date: 2023/12/7 11:46
// Author: Amu
// Description:
package clickhousex

import (
	"fmt"
	"testing"
)

func TestNewDB(t *testing.T) {
	db, err := NewDB(
		WithDebug(true),
		WithAddr("localhost:9000"),
		WithDatabase("gorm"),
		WithUsername("root"),
		WithPassword("123456"),
	)
	if err != nil {
		fmt.Printf("new db error: %#v\n", err)
	}
	fmt.Printf("db: %#v\n", db)
}
