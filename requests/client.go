// Package requests
// Date: 2022/9/6 00:08
// Author: Amu
// Description:
package requests

import "net/http"

var httpClient *http.Client

func buildHttpClient(r *Requests) {
	httpClient = &http.Client{
		Transport: buildHttpTransport(r),
		Timeout:   requestTimeout,
	}
}

func buildHttpTransport(r *Requests) *http.Transport {
	httpTransport := &http.Transport{
		Proxy:       nil,
		DialContext: nil,
	}
	return httpTransport
}
