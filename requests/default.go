// Package requests
// Date: 2022/9/5 10:12
// Author: Amu
// Description:
package requests

var requests = NewRequests()

func Get(url string) (*Response, error) {
	return requests.Get(url)
}

func Post(url string) (*Response, error) {
	return requests.Post(url)
}
