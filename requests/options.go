// Package requests
// Date: 2022/9/6 00:39
// Author: Amu
// Description:
package requests

import "time"

type Options func(*Requests)

func SetHeaders(headers map[string]string) Options {
	return func(c *Requests) {
		c.Headers = headers
	}
}

func SetCookies(cookies map[string]string) Options {
	return func(c *Requests) {
		c.Cookies = cookies
	}
}

func SetData(data map[string]string) Options {
	return func(r *Requests) {
		r.Data = data
	}
}

func SetParam(param map[string]string) Options {
	return func(r *Requests) {
		r.Param = param
	}
}

func SetTimeout(timeout time.Duration) Options {
	return func(r *Requests) {
		r.Timeout = timeout
	}
}
