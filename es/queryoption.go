// Package es
// Date: 2023/12/1 16:04
// Author: Amu
// Description:
package es

import "github.com/olivere/elastic/v7"

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
	service := c.Search()
	for _, option := range options {
		service = option.Action(service)
	}
	return service
}

func (c *Client) GetScrollQueryService(options ...ScrollQueryOption) *elastic.ScrollService {
	service := c.Scroll()
	for _, option := range options {
		service = option.ActionScroll(service)
	}
	return service
}

func (c *Client) GetCountService(options ...CountOption) *elastic.CountService {
	service := c.Count()
	for _, option := range options {
		service = option.Action(service)
	}
	return service
}

func (c *Client) GetDeleteService(options ...DeleteOption) *elastic.DeleteService {
	service := c.Delete()
	for _, option := range options {
		service = option.Action(service)
	}
	return service
}
