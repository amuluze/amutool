// Package es
// Date: 2023/12/1 16:05
// Author: Amu
// Description:
package es

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Age      int    `json:"age"`
	Sex      int    `json:"sex"`
}

func (u User) GetIndexName() string {
	return "am_users"
}

func (u User) GetId() string {
	return u.ID
}

func TestQueryAll(t *testing.T) {
	client, err := NewEsClient(
		WithAddr("http://localhost:9200"),
		WithUsername("elastic"),
		WithPassword("123456"),
		WithSniff(true),
		WithDebug(true),
		WithHealthcheck(true),
	)
	if err != nil {
		panic(err)
	}
	res, err := client.Search().Index(currentIndexName).Do(context.TODO())
	if err != nil {
		panic(err)
	}
	var u User
	for _, item := range res.Each(reflect.TypeOf(u)) {
		t := item.(User)
		fmt.Printf("user search: %#v\n", t)
	}
}

func TestFind(t *testing.T) {
	client, _ := NewEsClient(
		WithAddr("http://localhost:9200"),
		WithUsername("elastic"),
		WithPassword("123456"),
		WithSniff(true),
		WithDebug(true),
		WithHealthcheck(true),
	)
	query := NewBoolQuery().QueryName("test-query")
	query.Filter(NewRangeQuery("age").Gte(200).Lte(478))
	queryService := client.GetQueryService(
		&Page{Offset: 20, Count: 20},
		&Sort{Params: [][2]interface{}{{"age", false}}}, // true 升序 false 降序
		&Attr{Attr: []string{"username"}, Include: true},
	)
	fmt.Printf("query: %#v\n", query)
	result, err := queryService.Query(query).Do(context.TODO())
	if err != nil {
		panic(err)
	}
	for _, item := range result.Each(reflect.TypeOf(User{})) {
		t := item.(User)
		fmt.Printf("user search: %#v\n", t)
	}
	//
	//res, err := client.Find(context.TODO(), User{}, &Page{Offset: 200, Count: 20})
	//if err != nil {
	//	panic(err)
	//}
	//for _, item := range res.Each(reflect.TypeOf(User{})) {
	//	t := item.(User)
	//	fmt.Printf("user search: %#v\n", t)
	//}
}

func TestCountQuery(t *testing.T) {
	client, _ := NewEsClient(
		WithAddr("http://localhost:9200"),
		WithUsername("elastic"),
		WithPassword("123456"),
		WithSniff(true),
		WithDebug(true),
		WithHealthcheck(true),
	)
	query := NewBoolQuery().QueryName("count-query")
	query.Filter(NewRangeQuery("age").Gte(3890).Lte(10000))
	countQuery := client.GetCountService().Index(currentIndexName).Query(query)
	do, err := countQuery.Do(context.TODO())
	if err != nil {
		panic(err)
	}
	fmt.Printf("count: %v\n", do)
}
