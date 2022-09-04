// Package requests
// Date: 2022/9/1 23:05
// Author: Amu
// Description:
package requests

func Get(url string, options ...Options) (*Response, error) {
	return DoRequest("GET", url, options...)
}
