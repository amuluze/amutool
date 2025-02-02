// Package requests
// Date:   2024/12/13 18:27
// Author: Amu
// Description:
package requests

import (
	"testing"
)

type PostResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []User `json:"data"`
}

type MessageReponse struct {
	Message string `json:"message"`
}

func TestGetTwo(t *testing.T) {
	reply := &MessageReponse{}
	err := Get("http://1xxx.xxx.xxx.xxx:8090/aaa", nil, reply)
	if err != nil {
		t.Fatalf("get error: %v", err)
	}
	t.Logf("get success: %#v", reply)
}

func TestPost(t *testing.T) {
	params := &User{
		Name: "Amu",
		Age:  18,
	}
	reply := &PostResponse{}
	err := Post("http://1xxx.xxx.xxx.xxx:8090/users", params, reply, SetHeader("Referer", "https://example.com/test"), SetCookie("cid", "123456"))
	if err != nil {
		t.Fatalf("post error: %v", err)
	}
	t.Logf("post success: %#v", reply)
}
