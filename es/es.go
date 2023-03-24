// Package es
// Date: 2023/3/24 14:30
// Author: Amu
// Description:
package es

import (
	"log"
	"os"

	"github.com/olivere/elastic/v7"
)

type Config struct {
	Addr        string
	Username    string
	Password    string
	Sniff       bool
	Debug       bool
	Healthcheck bool
}

func NewEsClient(cfg *Config) *elastic.Client {
	client, err := elastic.NewClient(
		//elastic 服务地址
		elastic.SetURL(cfg.Addr),
		elastic.SetSniff(cfg.Sniff),
		elastic.SetHealthcheck(cfg.Healthcheck),
		elastic.SetBasicAuth(cfg.Username, cfg.Password),
		// 设置错误日志输出
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		// 设置info日志输出
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)))
	if err != nil {
		log.Fatalln("Failed to create elastic client")
	}
	return client
}
