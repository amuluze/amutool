// Package requests
// Date:   2024/12/13 16:41
// Author: Amu
// Description:
package requests

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

type Requests struct {
	client *fasthttp.Client

	headers map[string]string
	cookies map[string]string
}

func NewRequests() *Requests {
	return &Requests{
		client: &fasthttp.Client{},
	}
}

func (r *Requests) GET(url string, queryString string) (*fasthttp.Response, error) {

	uri := fasthttp.AcquireURI()
	defer fasthttp.ReleaseURI(uri)
	uri.Update(url)

	uri.SetQueryString(queryString)
	fmt.Printf("uri with query string: %s\n", uri.String())

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(uri.String())
	req.Header.SetMethod(fasthttp.MethodGet)
	req.SetBody([]byte(queryString))
	for key, value := range r.headers {
		req.Header.Set(key, value)
	}
	for key, value := range r.cookies {
		req.Header.SetCookie(key, value)
	}

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	return resp, r.client.Do(req, resp)
}
