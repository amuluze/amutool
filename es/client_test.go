// Package es
// Date: 2023/3/24 14:39
// Author: Amu
// Description:
package es

import (
	"fmt"
	"testing"

	"gitee.com/amuluze/amutool/conf"
)

func getClient() *Client {
	var cfg = new(Config)
	conf.MustLoad(cfg, "./config.toml")
	fmt.Printf("cfg: %#v\n", cfg)
	client, err := NewEsClient(cfg)
	if err != nil {
		fmt.Println(err)
	}
	return client
}

func TestIsRunning(t *testing.T) {
	var esClient = getClient()
	isRunning := esClient.IsRunning()
	fmt.Println(isRunning)
}
