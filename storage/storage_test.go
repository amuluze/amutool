// Package storage
// Date: 2024/3/26 12:20
// Author: Amu
// Description:
package storage

import "testing"

func TestString(t *testing.T) {
	storage := NewStorage("/Users/amu/Desktop/storage.db")
	err := storage.PutString("hello", "world")
	if err != nil {
		t.Log("storage put string error", err)
	}
	res, err := storage.GetString("hello")
	t.Log(res, err)
}

type User struct {
	Name string
	Age  int
	Sex  int
}

func TestJson(t *testing.T) {
	storage := NewStorage("/Users/amu/Desktop/storage.db")
	user := User{Name: "john", Age: 12, Sex: 1}
	err := storage.PutJson("json", user)
	if err != nil {
		t.Log("storage put json error", err)
	}

	var u User
	err = storage.GetJson("json", &u)
	t.Log(u.Name, err)
}
