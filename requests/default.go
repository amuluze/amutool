// Package requests
// Date: 2022/9/5 10:12
// Author: Amu
// Description:
package requests

var requests = NewRequests()

func Get(url string, headers *requestsHeaders, cookies *requestsCookies, params *requestsParam) (*Responses, error) {
	return requests.Get(url, headers, cookies, params)
}

func Post(url string, headers *requestsHeaders, cookies *requestsCookies, data *requestsData, json *requestsJson) (*Responses, error) {
	return requests.Post(url, headers, cookies, data, json)
}
