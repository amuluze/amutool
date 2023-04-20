// Package es
// Date: 2023/3/24 14:30
// Author: Amu
// Description:
package es

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"gitee.com/amuluze/amutool/iohelper"

	"github.com/olivere/elastic/v7"
)

// ================================================= config =================================================

type Config struct {
	Addr              string
	Username          string
	Password          string
	ConfigPath        string
	BulkFlushInterval string
	Sniff             bool
	Debug             bool
	Healthcheck       bool
	BulkWorkers       int
	BulkActions       int
	BulkSize          int
	IndexNames        []string
}

// ================================================= client =================================================

type Client struct {
	*elastic.Client
}

func NewEsClient(cfg *Config) (*Client, error) {
	client, err := elastic.NewClient(
		//elastic 服务地址
		elastic.SetURL(cfg.Addr),
		elastic.SetSniff(cfg.Sniff),
		elastic.SetHealthcheck(cfg.Healthcheck),
		elastic.SetBasicAuth(cfg.Username, cfg.Password),
		// 设置错误日志输出
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		// 设置info日志输出
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
	)
	info, code, err := client.Ping(cfg.Addr).Do(context.Background())
	if err != nil {
		log.Fatalln("Failed to create elastic client")
		return nil, err
	}
	log.Printf("Elasticsearch returned with code: %d and version: %s\n", code, info.Version.Number)
	// 创建 policy index
	cli := &Client{client}
	if len(cfg.IndexNames) > 0 {
		res, err := checkPolicy(cfg, cli)
		fmt.Printf("res: %v, err: %v\n", res, err)
		res, err = checkTemplate(cfg, cli)
		fmt.Printf("res: %v, err: %v\n", res, err)
		res, err = checkIndex(cfg, cli)
		fmt.Printf("res: %v, err: %v\n", res, err)
	}
	return cli, nil
}

// checkPolicy 检查及更新索引声明周期策略
func checkPolicy(cfg *Config, cli *Client) (bool, error) {
	try := 0
	for _, indexName := range cfg.IndexNames {
		policyFileName := PolicyFilePrefix + indexName + PolicyFileSuffix
		policyCheckTextMd5 := iohelper.NewCheckTextMd5(policyFileName, "", cfg.ConfigPath, PolicyFilePrefix+indexName)
		policyChanged := policyCheckTextMd5.Change()

		for {
			exists, err := cli.ILMPolicyExists(context.Background(), policyFileName)
			fmt.Printf("ex: %v, err: %v\n", exists, err)
			try++
			if try > CreatePolicyRetry {
				panic("try to create policy over 50 times")
			}
			//if err != nil {
			//	continue
			//}

			// 更新 policy
			if !exists || policyChanged {
				// 存在，但是需要更新，所以先删除旧的
				if exists {
					err := cli.DeleteILMPolicy(context.TODO(), policyFileName)
					if err != nil {
						continue
					}
				}

				err := cli.PutILMPolicy(context.TODO(), policyFileName, cfg.ConfigPath)
				fmt.Printf("put err: %v\n", err)
				if err != nil {
					continue
				}

				err = policyCheckTextMd5.Write()
				exists = true
			}
			if exists {
				break
			}
			time.Sleep(2 * time.Second)
		}
	}
	return true, nil
}

// checkTemplate 检查及更新索引模版
func checkTemplate(cfg *Config, cli *Client) (bool, error) {
	try := 0
	for _, indexName := range cfg.IndexNames {
		templateFIleName := TemplateFilePrefix + indexName + TemplateFileSuffix
		templateCheckTextMd5 := iohelper.NewCheckTextMd5(templateFIleName, "", cfg.ConfigPath, TemplateFilePrefix+indexName)
		templateChanged := templateCheckTextMd5.Change()

		for {
			exists, err := cli.TemplateExists(context.Background(), templateFIleName)
			try++
			if try > CreateTemplateRetry {
				panic("try to create index template over 50 times")
			}

			if err != nil {
				fmt.Printf("template exists error: %v, exists: %v\n", err, exists)
			}
			fmt.Printf("template exists error: %v, exists: %v\n", err, exists)
			if !exists || templateChanged {
				// 存在，但是需要更新，所以先删除旧的
				if exists {
					err := cli.DeleteIndexTemplate(context.TODO(), templateFIleName)
					if err != nil {
						continue
					}
				}

				err := cli.PutIndexTemplate(context.TODO(), templateFIleName, cfg.ConfigPath)
				if err != nil {
					fmt.Printf("put index template error: %v\n", err)
					continue
				}

				err = templateCheckTextMd5.Write()
			}

			if exists {
				break
			}
			time.Sleep(2 * time.Second)
		}
	}
	return true, nil
}

// checkIndex 检查索引是否存在如果不存在则创建
func checkIndex(cfg *Config, cli *Client) (bool, error) {
	try := 0
	for _, indexName := range cfg.IndexNames {
		for {
			exists, err := cli.IndexExists(context.Background(), indexName)
			try++
			if try > CreateIndexRetry {
				panic("try to create index over 50 times")
			}
			if err != nil {
				fmt.Printf("index exists error: %v", err)
			}
			if !exists {
				newIndexName := fmt.Sprintf("<%s-{now/d}-00001>", indexName)
				indexBody := `{
					"aliases": {
						"%s": {"is_write_index": true}
					}
				}`

				indexBody = fmt.Sprintf(indexBody, indexName)

				res, err := cli.CreateIndex(context.TODO(), newIndexName, indexBody)
				if err != nil || res == false {
					fmt.Printf("create index failure: %v\n", err)
				}
				exists = true
			}
			if exists {
				break
			}
			time.Sleep(2 * time.Second)
		}
	}
	return true, nil
}

// ================================================= option =================================================

type QueryOption interface {
	Action(service *elastic.SearchService) *elastic.SearchService
}

type ScrollQueryOption interface {
	ActionScroll(service *elastic.ScrollService) *elastic.ScrollService
}

type CountOption interface {
	Action(service *elastic.CountService) *elastic.CountService
}

type DeleteOption interface {
	Action(service *elastic.DeleteService) *elastic.DeleteService
}

// 自定义查询结果

type Attr struct {
	Attr    []string // 需要包含或剔除的字段
	Include bool     // 包含还是剔除
}

func (a *Attr) Action(service *elastic.SearchService) *elastic.SearchService {
	source := elastic.NewFetchSourceContext(true)
	if a.Include {
		source = source.Include(a.Attr...)
	} else {
		source = source.Exclude(a.Attr...)
	}
	return service.FetchSourceContext(source)
}

func (a *Attr) ActionScroll(service *elastic.SearchService) *elastic.SearchService {
	source := elastic.NewFetchSourceContext(true)
	if a.Include {
		source = source.Include(a.Attr...)
	} else {
		source = source.Exclude(a.Attr...)
	}
	return service.FetchSourceContext(source)
}

// 分页

type Page struct {
	Count  int
	Offset int
}

func (p *Page) Action(service *elastic.SearchService) *elastic.SearchService {
	return service.From(p.Offset).Size(p.Count)
}

// 排序

type Sort struct {
	Params [][2]interface{}
}

func (s *Sort) Action(service *elastic.SearchService) *elastic.SearchService {
	for _, sort := range s.Params {
		service = service.Sort(sort[0].(string), sort[1].(bool))
	}
	return service
}

func (s *Sort) ActionScroll(service *elastic.SearchService) *elastic.SearchService {
	for _, sort := range s.Params {
		service = service.Sort(sort[0].(string), sort[1].(bool))
	}
	return service
}

// SortInfo

type SortInfo struct {
	Params []elastic.SortInfo
}

func (si *SortInfo) Action(service *elastic.SearchService) *elastic.SearchService {
	for _, sortInfo := range si.Params {
		service = service.SortWithInfo(sortInfo)
	}
	return service
}

func (si *SortInfo) ActionScroll(service *elastic.SearchService) *elastic.SearchService {
	for _, sortInfo := range si.Params {
		service = service.SortWithInfo(sortInfo)
	}
	return service
}

// Agg

type Agg struct {
	Name string
	Agg  elastic.Aggregation
}

func (a *Agg) Action(service *elastic.SearchService) *elastic.SearchService {
	service = service.Aggregation(a.Name, a.Agg)
	return service
}

// Count

type CountQuery struct {
	Query elastic.Query
}

func (c *CountQuery) Action(service *elastic.SearchService) *elastic.SearchService {
	service = service.Query(c.Query)
	return service
}

// Delete

type DeleteQuery struct {
	Query elastic.Query
}

func (d *DeleteQuery) Action(service *elastic.SearchService) *elastic.SearchService {
	service = service.Query(d.Query)
	return service
}

func (c *Client) GetQueryService(options ...QueryOption) *elastic.SearchService {
	service := c.Client.Search()
	for _, option := range options {
		service = option.Action(service)
	}
	return service
}

func (c *Client) GetScrollQueryService(options ...ScrollQueryOption) *elastic.ScrollService {
	service := c.Client.Scroll()
	for _, option := range options {
		service = option.ActionScroll(service)
	}
	return service
}

func (c *Client) GetCountService(options ...CountOption) *elastic.CountService {
	service := c.Client.Count()
	for _, option := range options {
		service = option.Action(service)
	}
	return service
}

func (c *Client) GetDeleteService(options ...DeleteOption) *elastic.DeleteService {
	service := c.Client.Delete()
	for _, option := range options {
		service = option.Action(service)
	}
	return service
}
