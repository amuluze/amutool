// Package requests
// Date: 2022/9/6 00:11
// Author: Amu
// Description:
package requests

import (
	"fmt"
	"net/http"
	"time"
)

type Requests struct {
	// 参数相关
	Headers map[string]string // header
	Cookies map[string]string // cookies
	Param   map[string]string // params
	Data    map[string]string // data
	Json    interface{}       // json

	// request 相关
	DialTimeout         time.Duration
	DialKeepAlive       time.Duration
	TLSHandshakeTimeout time.Duration

	// Client 相关
	Timeout             time.Duration
	MaxConnsPerHost     int
	MaxIdleConnsPerHost int
	IdleConnTimeout     time.Duration
}

func NewRequests() *Requests {
	requestsConfig := &Requests{
		MaxConnsPerHost:     maxConnsPerHost,
		MaxIdleConnsPerHost: maxIdleConnsPerHost,
		IdleConnTimeout:     idleConnTimeout,
		Timeout:             requestTimeout,
	}

	buildHttpClient(requestsConfig)

	return requestsConfig
}

func (r *Requests) Get(url string, options ...Option) (*Responses, error) {
	for _, option := range options {
		option(r)
	}

	fmt.Printf("get requests: %+v\n", r)
	return DoRequests("GET", url, r)
}

func (r *Requests) Post(url string, options ...Option) (*Responses, error) {
	for _, option := range options {
		option(r)
	}

	fmt.Printf("post requests: %+v\n", r)
	return DoRequests("POST", url, r)
}

func DoRequests(method, url string, r *Requests) (*Responses, error) {
	return buildResponses(buildRequests(method, url, r))
}

func buildRequests(method, url string, r *Requests) (*http.Response, error) {
	req, err := buildHttpRequest(method, url, r)
	if err != nil {
		return nil, err
	}

	addHeaders(r, req)
	addCookies(r, req)

	return httpClient.Do(req)
}

func addHeaders(r *Requests, req *http.Request) {
	for key, value := range r.Headers {
		req.Header.Add(key, value)
	}
}

func addCookies(r *Requests, req *http.Request) {
	for key, value := range r.Cookies {
		cookie := http.Cookie{Name: key, Value: value}
		req.AddCookie(&cookie)
	}
}
