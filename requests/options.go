// Package requests
// Date: 2022/9/4 00:58
// Author: Amu
// Description:
package requests

type Options func(*Request)

func SetHeaders(headers map[string]string) Options {
	return func(request *Request) {

	}
}

func SetCookies(cookies map[string]string) Options {
	return func(request *Request) {

	}
}
