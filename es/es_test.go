// Package es
// Date: 2023/3/24 14:39
// Author: Amu
// Description:
package es

import (
	"fmt"
	"testing"

	"github.com/olivere/elastic/v7"

	"gitee.com/amuluze/amutool/conf"
)

func getClient() *elastic.Client {
	var cfg = new(Config)
	conf.MustLoad(cfg, "./config.yaml")
	fmt.Printf("cfg: %#v\n", cfg)
	client := NewEsClient(cfg)
	return client
}

func TestNewEsClient(t *testing.T) {
	c := getClient()
	fmt.Println(c)
}

func TestIsRunning(t *testing.T) {
	client := getClient()
	isRunning := client.IsRunning()
	fmt.Println(isRunning)
}
