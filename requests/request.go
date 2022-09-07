// Package requests
// Date: 2022/9/5 09:54
// Author: Amu
// Description:
package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func buildHttpRequest(method, url string, r *Requests) (*http.Request, error) {
	var err error
	if r.Param != nil {
		if url, err = buildURLParams(url, r.Param); err != nil {
			return nil, err
		}
	}
	fmt.Printf("parsed url: %s\n", url)

	if r.Data != nil {
		return createFormDataRequest(method, url, r)
	}

	if r.Json != nil {
		return createJsonRequest(method, url, r)
	}

	return http.NewRequest(method, url, nil)
}

// buildURLParams 将 params 中的参数拼接到 url 中
func buildURLParams(requestURL string, params map[string]string) (string, error) {
	parsedURL, err := url.Parse(requestURL)
	if err != nil {
		return "", nil
	}

	parsedQuery, err := url.ParseQuery(parsedURL.RawQuery)

	if err != nil {
		return "", err
	}

	for key, value := range params {
		parsedQuery.Set(key, value)
	}
	fmt.Printf("parsed query: %+v\n", parsedQuery)
	return addQueryParams(parsedURL, parsedQuery), nil
}

func addQueryParams(parsedURL *url.URL, parsedQuery url.Values) string {
	return strings.Join([]string{strings.Replace(parsedURL.String(), "?"+parsedURL.RawQuery, "", -1), parsedQuery.Encode()}, "?")
}

func createFormDataRequest(method, url string, r *Requests) (*http.Request, error) {
	req, err := http.NewRequest(method, url, strings.NewReader(encodeDataValues(r.Data)))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return req, nil
}

func encodeDataValues(data map[string]string) string {
	urlValues := &url.Values{}

	for key, value := range data {
		urlValues.Set(key, value)
	}
	return urlValues.Encode()
}

func createJsonRequest(method, url string, r *Requests) (*http.Request, error) {
	var reader io.Reader

	switch r.Json.(type) {
	case string:
		reader = strings.NewReader(r.Json.(string))
	case []byte:
		reader = bytes.NewReader(r.Json.([]byte))
	default:
		byteSlice, err := json.Marshal(r.Json)
		if err != nil {
			return nil, err
		}
		reader = bytes.NewReader(byteSlice)
	}

	req, err := http.NewRequest(method, url, reader)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}
