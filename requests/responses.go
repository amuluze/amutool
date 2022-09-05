// Package requests
// Date: 2022/9/1 23:06
// Author: Amu
// Description:
package requests

import "net/http"

type Responses struct {
	Ok          bool
	Error       error
	RawResponse *http.Response
	StatusCode  int
	Header      http.Header
}

func buildResponses(resp *http.Response, err error) (*Responses, error) {
	if err != nil {
		return &Responses{Error: err}, err
	}

	goodResp := &Responses{
		Ok:          resp.StatusCode >= 200 && resp.StatusCode < 300,
		Error:       nil,
		RawResponse: resp,
		StatusCode:  resp.StatusCode,
		Header:      resp.Header,
	}

	return goodResp, nil
}
