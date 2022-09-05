// Package requests
// Date: 2022/9/5 09:54
// Author: Amu
// Description:
package requests

import (
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Request struct {
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
}

type Options func(*Request)

func SetHeaders(headers map[string]string) Options {
	return func(c *Request) {
		c.Headers = headers
	}
}

func SetCookies(cookies map[string]string) Options {
	return func(c *Request) {
		c.Cookies = cookies
	}
}

func SetData(data map[string]string) Options {
	return func(r *Request) {
		r.Data = data
	}
}

func SetParam(param map[string]string) Options {
	return func(r *Request) {
		r.Param = param
	}
}

func SetTimeout(timeout time.Duration) Options {
	return func(r *Request) {
		r.Timeout = timeout
	}
}

func NewRequests(options ...Options) *Request {
	config := &Request{
		DialTimeout:   dialTimeout,
		DialKeepAlive: dialKeepAlive,
		Timeout:       requestTimeout,
	}
	for _, option := range options {
		option(config)
	}

	return &Request{}
}

func (r *Request) Get(url string) (*Response, error) {
	return DoRequest("GET", url, r)
}

func (r *Request) Post(url string) (*Response, error) {
	return DoRequest("POST", url, r)
}

func DoRequest(method, url string, r *Request) (*Response, error) {
	return buildResponse(buildRequest(method, url, r))
}

func buildRequest(method, url string, r *Request) (*http.Response, error) {
	client := buildHttpClient(r)

	var err error
	if len(r.Param) != 0 {
		url, err = buildUrlParams(url, r.Param)
		if err != nil {
			return nil, err
		}
	}
	req, _ := buildHttpRequest(method, url, r)
	if err != nil {
		return nil, err
	}

	addHeaders(r, req)
	addCookies(r, req)

	return client.Do(req)
}

func buildUrlParams(userURL string, params map[string]string) (string, error) {
	parsedURL, err := url.Parse(userURL)

	if err != nil {
		return "", err
	}

	parsedQuery, err := url.ParseQuery(parsedURL.RawQuery)

	if err != nil {
		return "", nil
	}

	for key, value := range params {
		parsedQuery.Set(key, value)
	}

	return addQueryParams(parsedURL, parsedQuery), nil
}

func buildHttpClient(r *Request) *http.Client {
	return &http.Client{
		Transport: buildHttpTransport(r),
		Timeout:   requestTimeout,
	}
}

func buildHttpRequest(method, url string, r *Request) (*http.Request, error) {
	if r.Data != nil {
		return createBasicRequest(method, url, r)
	}
	return http.NewRequest(method, url, nil)
}

func buildHttpTransport(r *Request) *http.Transport {
	httpTransport := &http.Transport{
		Proxy:       nil,
		DialContext: nil,
	}
	return httpTransport
}

func addHeaders(r *Request, req *http.Request) {
	for key, value := range r.Headers {
		req.Header.Set(key, value)
	}
}

func addCookies(r *Request, req *http.Request) {}

func addQueryParams(parsedURL *url.URL, parsedQuery url.Values) string {
	return strings.Join([]string{strings.Replace(parsedURL.String(), "?"+parsedURL.RawQuery, "", -1), parsedQuery.Encode()}, "?")
}

func createBasicRequest(httpMethod, userURL string, r *Request) (*http.Request, error) {

	req, err := http.NewRequest(httpMethod, userURL, strings.NewReader(encodePostValues(r.Data)))

	if err != nil {
		return nil, err
	}

	// The content type must be set to a regular form
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return req, nil
}

func encodePostValues(postValues map[string]string) string {
	urlValues := &url.Values{}

	for key, value := range postValues {
		urlValues.Set(key, value)
	}

	return urlValues.Encode() // This will sort all of the string values
}
