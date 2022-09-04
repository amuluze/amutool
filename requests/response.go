// Package requests
// Date: 2022/9/1 23:06
// Author: Amu
// Description:
package requests

import "net/http"

type Response struct {
	Ok          bool
	Error       error
	RawResponse *http.Response
	StatusCode  int
	Header      http.Header
}

func buildResponse(resp *http.Response, err error) (*Response, error) {
	if err != nil {
		return &Response{Error: err}, err
	}

	goodResp := &Response{
		Ok:          resp.StatusCode >= 200 && resp.StatusCode < 300,
		Error:       nil,
		RawResponse: resp,
		StatusCode:  resp.StatusCode,
		Header:      resp.Header,
	}

	return goodResp, nil
}
