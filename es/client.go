// Package es
// Date: 2023/3/24 14:30
// Author: Amu
// Description:
package es

import (
	"context"
	"log"
	"os"

	"github.com/olivere/elastic/v7"
)

// ================================================= client =================================================

type Client struct {
	*elastic.Client
}

func NewEsClient(opts ...ClientOption) (*Client, error) {
	opt := &clientOption{}
	for _, o := range opts {
		o(opt)
	}

	client, err := elastic.NewClient(
		//elastic 服务地址
		elastic.SetURL(opt.Addr),
		elastic.SetSniff(opt.Sniff),
		elastic.SetHealthcheck(opt.Healthcheck),
		elastic.SetBasicAuth(opt.Username, opt.Password),
		// 设置错误日志输出
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		// 设置info日志输出
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
	)
	info, code, err := client.Ping(opt.Addr).Do(context.Background())
	if err != nil {
		log.Fatalln("Failed to create elastic client")
		return nil, err
	}
	log.Printf("Elasticsearch returned with code: %d and version: %s\n", code, info.Version.Number)

	// 创建 policy index
	cli := &Client{client}
	return cli, nil
}
