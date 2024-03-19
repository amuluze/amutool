// Package requests
// Date: 2022/9/7 23:56
// Author: Amu
// Description:
package requests

import (
	"fmt"
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	params := map[string]string{
		"good": "job",
	}

	resp, err := Get("http://localhost:9000/get?hello=world", SetParams(params))
	if err != nil {
		return
	}

	fmt.Printf("resp: %#v\n", resp)
	fmt.Printf("response ok: %+v\n", resp.Ok)
	fmt.Printf("response err: %+v\n", resp.Error)
	fmt.Printf("response raw: %+v\n", resp.RawResponse)
	fmt.Printf("response status: %+v\n", resp.StatusCode)
	fmt.Printf("response header: %+v\n", resp.Header)
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>><<<<<<<<<<<<<<<<<")

	fmt.Printf("resp: %#v\n", resp)
}

func TestPost(t *testing.T) {
	data := map[string]string{
		"name": "amuluze",
		"age":  "12",
	}

	resp, err := Post("http://localhost:9000/post", SetData(data))
	if err != nil {
		return
	}
	fmt.Printf("data resp: %#v\n", resp)
}

func TestSetHeaders(t *testing.T) {
	type Result struct {
		Hello    string `json:"hello"`
		Duration int    `json:"duration"`
		IsDelete bool   `json:"is_delete"`
	}
	headers := map[string]string{
		"token": "123456",
	}
	resp, _ := Get("http://localhost:9000/header", SetHeaders(headers))

	fmt.Printf("header resp: %#v\n", resp.String())
	var res Result
	//_ = json.Unmarshal(resp.Bytes(), &res)
	//fmt.Printf("res: %#v\n", res)
	resp.JSON(res)
	fmt.Printf("header resp: %#v\n", res)
}

func TestSetCookies(t *testing.T) {
	cookies := map[string]string{
		"session_id": "asdfghjjkl",
	}
	resp, _ := Get("http://localhost:9000/header", SetCookies(cookies))
	fmt.Printf("header resp: %#v\n", resp.String())
}

func TestSetJson(t *testing.T) {
	load := map[string]interface{}{
		"name":        "amuluze",
		"age":         12,
		"is_delete":   false,
		"create_time": time.Now(),
	}

	resp, err := Post("http://localhost:9000/json/post", SetJson(load))
	if err != nil {
		return
	}
	fmt.Printf("json resp: %#v\n", resp)
}
