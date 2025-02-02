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

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Age      int64  `json:"age"`
	Sex      int64  `json:"sex"`
}

func TestNewDB(t *testing.T) {
	db, err := NewDB(
		WithDebug(true),
		WithType("clickhouse"),
		WithHost("localhost"),
		WithPort("9000"),
		WithUsername("root"),
		WithPassword("123456"),
		WithDBName("gorm"),
	)
	if err != nil {
		log.Fatalf("new db error: %v", err)
	}
	defer db.Close()
	fmt.Println(db)
}

func TestQuery(t *testing.T) {
	db, _ := NewDB(
		WithDebug(true),
		WithType("clickhouse"),
		WithHost("localhost"),
		WithPort("9000"),
		WithUsername("root"),
		WithPassword("123456"),
		WithDBName("gorm"),
	)
	defer db.Close()

	query := OptionDB(
		db,
		WithById("123456ddd"),
	)
	var user User
	if err := query.First(&user); err != nil {
		fmt.Printf("get user by %s error: %v\n", "123456ddd", err)
	}
	fmt.Printf("user: %#v\n", user)
}
