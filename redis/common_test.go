// Package redis
// Date: 2023/12/4 14:02
// Author: Amu
// Description:
package redis

import (
	"fmt"
	"testing"
)

func TestKeys(t *testing.T) {
	rc, err := NewClient()
	if err != nil {
		panic(err)
	}
	keys, err := rc.Keys()
	if err != nil {
		panic(err)
	}
	fmt.Printf("keys: %v\n", keys)
}

func TestKeyExists(t *testing.T) {
	key := "hello"
	rc, err := NewClient()
	if err != nil {
		panic(err)
	}
	exists, err := rc.Exists(key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("key %s exists %v\n", key, exists)
}

func TestTTL(t *testing.T) {
	rc, err := NewClient()
	if err != nil {
		panic(err)
	}
	key := "hello"
	ttl, err := rc.TTL(key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("key %s ttl %v\n", key, ttl)
}
