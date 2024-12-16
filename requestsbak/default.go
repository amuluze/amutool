// Package requests
// Date: 2022/9/5 10:12
// Author: Amu
// Description:
package requests

var requests = NewRequests()

func Get(url string, options ...Option) (*Responses, error) {
	return requests.Get(url, options...)
}

func Post(url string, options ...Option) (*Responses, error) {
	return requests.Post(url, options...)
}
