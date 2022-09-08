// Package requests
// Date: 2022/9/8 23:50
// Author: Amu
// Description:
package requests

import "time"

type Option func(req *Requests)

func SetHeaders(rh map[string]string) Option {
	return func(req *Requests) {
		req.Headers = rh
	}
}

func SetCookies(rc map[string]string) Option {
	return func(req *Requests) {
		req.Cookies = rc
	}
}

func SetParam(rp map[string]string) Option {
	return func(req *Requests) {
		req.Param = rp
	}
}

func SetData(rd map[string]string) Option {
	return func(req *Requests) {
		req.Data = rd
	}
}

func SetJson(rj interface{}) Option {
	return func(req *Requests) {
		req.Json = rj
	}
}

func SetRequestTimeout(rt time.Duration) Option {
	return func(req *Requests) {
		req.Timeout = rt
	}
}

func SetMaxConns(mc int) Option {
	return func(req *Requests) {
		req.MaxConnsPerHost = mc
	}
}

func SetMaxIdle(mi int) Option {
	return func(req *Requests) {
		req.MaxIdleConnsPerHost = mi
	}
}

func SetIdleConnTimeout(ic time.Duration) Option {
	return func(req *Requests) {
		req.IdleConnTimeout = ic
	}
}
