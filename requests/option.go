// Package requests
// Date:   2024/12/13 16:44
// Author: Amu
// Description:
package requests

type Option func(req *Requests)

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
