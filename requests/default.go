// Package requests
// Date:   2024/12/13 16:41
// Author: Amu
// Description:
package requests

import "fmt"

var requests = NewRequests()

func Get(url string, params FormData, options ...Option) {
	get, err := requests.GET(url, params.ToQuery(), nil)
	if err != nil {
		return
	}
	fmt.Println(get)
}

func Post(url string, data JsonData) {}
