// Package requests
// Date: 2022/9/5 09:54
// Author: Amu
// Description:
package requests

import (
	"net/http"
	"net/url"
	"strings"
)

func buildHttpRequest(method, url string, r *Requests) (*http.Request, error) {
	return http.NewRequest(method, url, nil)
}

func addHeaders(r *Requests, req *http.Request) {
	for key, value := range r.Headers {
		req.Header.Set(key, value)
	}
}

func addCookies(r *Requests, req *http.Request) {}

func addQueryParams(parsedURL *url.URL, parsedQuery url.Values) string {
	return strings.Join([]string{strings.Replace(parsedURL.String(), "?"+parsedURL.RawQuery, "", -1), parsedQuery.Encode()}, "?")
}

func encodePostValues(postValues map[string]string) string {
	urlValues := &url.Values{}

	for key, value := range postValues {
		urlValues.Set(key, value)
	}

	return urlValues.Encode() // This will sort all of the string values
}
