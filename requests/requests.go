// Package requests
// Date:   2024/12/13 16:41
// Author: Amu
// Description:
package requests

import (
	"crypto/tls"

	"github.com/valyala/fasthttp"
)

type Option func(req *Requests)

type Requests struct {
	client *fasthttp.Client

	headers map[string]string
	cookies map[string]string
}

func SetHeader(key, value string) Option {
	return func(req *Requests) {
		req.headers[key] = value
	}
}

func SetCookie(key, value string) Option {
	return func(req *Requests) {
		req.cookies[key] = value
	}
}

func NewRequests() *Requests {
	return &Requests{
		client: &fasthttp.Client{
			TLSConfig: &tls.Config{InsecureSkipVerify: true},
		},
		headers: make(map[string]string),
		cookies: make(map[string]string),
	}
}

func (r *Requests) GET(url string, queryString string) ([]byte, error) {
	defer r.clear()

	uri := fasthttp.AcquireURI()
	uri.Update(url)
	uri.SetQueryString(queryString)
	defer fasthttp.ReleaseURI(uri)

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(uri.String())
	req.Header.SetMethod(fasthttp.MethodGet)
	req.SetBody([]byte(queryString))
	defer fasthttp.ReleaseRequest(req)

	for key, value := range r.headers {
		req.Header.Set(key, value)
	}
	for key, value := range r.cookies {
		req.Header.SetCookie(key, value)
	}

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	err := r.client.Do(req, resp)
	if err != nil {
		return nil, err
	}
	return resp.Body(), nil
}

func (r *Requests) POST(url string, body []byte) ([]byte, error) {
	defer r.clear()

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.SetMethod(fasthttp.MethodPost)
	req.Header.SetContentType("application/json")
	req.SetBody(body)
	defer fasthttp.ReleaseRequest(req)

	for key, value := range r.headers {
		req.Header.Set(key, value)
	}
	for key, value := range r.cookies {
		req.Header.SetCookie(key, value)
	}

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	err := r.client.Do(req, resp)
	if err != nil {
		return nil, err
	}
	return resp.Body(), nil
}

func (r *Requests) PUT(url string, body []byte) ([]byte, error) {
	defer r.clear()

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.SetMethod(fasthttp.MethodPut)
	req.Header.SetContentType("application/json")
	req.SetBody(body)
	defer fasthttp.ReleaseRequest(req)

	for key, value := range r.headers {
		req.Header.Set(key, value)
	}
	for key, value := range r.cookies {
		req.Header.SetCookie(key, value)
	}

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	err := r.client.Do(req, resp)
	if err != nil {
		return nil, err
	}
	return resp.Body(), nil
}

func (r *Requests) DELETE(url string, body []byte) ([]byte, error) {
	defer r.clear()

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.SetMethod(fasthttp.MethodDelete)
	req.Header.SetContentType("application/json")
	req.SetBody(body)
	defer fasthttp.ReleaseRequest(req)

	for key, value := range r.headers {
		req.Header.Set(key, value)
	}
	for key, value := range r.cookies {
		req.Header.SetCookie(key, value)
	}
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	err := r.client.Do(req, resp)
	if err != nil {
		return nil, err
	}
	return resp.Body(), nil
}

func (r *Requests) clear() {
	r.headers = make(map[string]string)
	r.cookies = make(map[string]string)
}
