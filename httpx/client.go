// Package httpx
// Date: 2023/4/6 10:28
// Author: Amu
// Description:
package httpx

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type ResponseWrapper struct {
	StatusCode int
	Body       string
	Header     http.Header
}

func Get(url string, params *string, timeout int) ResponseWrapper {
	var body *bytes.Buffer
	if params != nil {
		body = bytes.NewBufferString(*params)
	}
	req, err := http.NewRequest("GET", url, body)
	if err != nil {
		return createRequestError(err)
	}

	return request(req, timeout)
}

// PostParams post form data
func PostParams(url string, params *string, timeout int) ResponseWrapper {
	var body *bytes.Buffer
	if params != nil {
		body = bytes.NewBufferString(*params)
	}
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return createRequestError(err)
	}
	req.Header.Set("Content-type", "application/x-www-form-urlencoded")

	return request(req, timeout)
}

// PostJson post json
func PostJson(url string, params *string, timeout int) ResponseWrapper {
	var body *bytes.Buffer
	if params != nil {
		body = bytes.NewBufferString(*params)
	}
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return createRequestError(err)
	}
	req.Header.Set("Content-type", "application/json")

	return request(req, timeout)
}

func request(req *http.Request, timeout int) ResponseWrapper {
	wrapper := ResponseWrapper{StatusCode: 0, Body: "", Header: make(http.Header)}
	client := &http.Client{}
	if timeout > 0 {
		client.Timeout = time.Duration(timeout) * time.Second
	}
	setRequestHeader(req)
	resp, err := client.Do(req)
	if err != nil {
		wrapper.Body = fmt.Sprintf("执行HTTP请求错误-%s", err.Error())
		return wrapper
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		wrapper.Body = fmt.Sprintf("读取HTTP请求返回值失败-%s", err.Error())
		return wrapper
	}
	wrapper.StatusCode = resp.StatusCode
	wrapper.Body = string(body)
	wrapper.Header = resp.Header

	return wrapper
}

func setRequestHeader(req *http.Request) {
	req.Header.Set("User-Agent", "golang/gocron")
}

func createRequestError(err error) ResponseWrapper {
	errorMessage := fmt.Sprintf("创建HTTP请求错误-%s", err.Error())
	return ResponseWrapper{0, errorMessage, make(http.Header)}
}
