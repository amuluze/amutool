// Package redis
// Date: 2022/9/26 17:56
// Author: Amu
// Description:
package redis

import (
	"fmt"
	"testing"

	"gitee.com/amuluze/amutool/conf"
)

func getClient() *Client {
	var cfg = new(Config)
	conf.MustLoad(cfg, "./config.toml")
	fmt.Printf("cfg: %#v\n", cfg)
	client := NewClient(cfg)
	return client
}

var client = getClient()

func TestKeys(t *testing.T) {
	keys, err := client.Keys()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(keys)
}

func TestGet(t *testing.T) {
	result, err := client.Get("hello")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
