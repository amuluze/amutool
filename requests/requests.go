// Package requests
// Date: 2022/9/6 00:11
// Author: Amu
// Description:
package requests

import (
	"net/http"
	"net/url"
	"time"
)

type Requests struct {
	Headers map[string]string // header
	Cookies map[string]string // cookies
	Data    map[string]string // data
	Param   map[string]string // params

	// DialTimeout is the maximum amount of time a dial will wait for a connect to complete
	DialTimeout time.Duration

	// KeepAlive specifies the keep-alive period for an active
	// network connection. If zero, keep-alive are not enabled.
	DialKeepAlive time.Duration

	// TLSHandshakeTimeout specifies the maximum amount of time waiting to
	// wait for a TLS handshake. Zero means no timeout.
	TLSHandshakeTimeout time.Duration

	// RequestTimeout is the maximum amount of time a whole request(include dial / request / redirect)
	// will wait.
	Timeout time.Duration

	// Proxies is a map in the following format
	// *protocol* => proxy address e.g http => http://127.0.0.1:8080
	Proxies map[string]*url.URL
}

func NewRequests(options ...Options) *Requests {
	requestsConfig := &Requests{
		DialTimeout:   dialTimeout,
		DialKeepAlive: dialKeepAlive,
		Timeout:       requestTimeout,
	}
	for _, option := range options {
		option(requestsConfig)
	}

	buildHttpClient(requestsConfig)

	return requestsConfig
}

func (r *Requests) Get(url string) (*Responses, error) {
	return DoRequests("GET", url, r)
}

func (r *Requests) Post(url string) (*Responses, error) {
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
