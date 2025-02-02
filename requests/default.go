// Package requests
// Date:   2024/12/13 16:41
// Author: Amu
// Description:
package requests

import (
	"encoding/json"
)

var req = NewRequests()

func Get(url string, params any, reply any, options ...Option) error {
	if len(options) > 0 {
		for _, option := range options {
			option(req)
		}
	}

	resp, err := req.GET(url, ToQuery(params))
	if err != nil {
		return err
	}
	if err := json.Unmarshal(resp, reply); err != nil {
		return err
	}
	return nil
}

func Post(url string, data any, reply any, options ...Option) error {
	if len(options) > 0 {
		for _, option := range options {
			option(req)
		}
	}
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}
	resp, err := req.POST(url, body)
	if err != nil {

		return err
	}
	if err := json.Unmarshal(resp, reply); err != nil {
		return err
	}
	return nil
}

func Put(url string, data any, reply any, options ...Option) error {
	if len(options) > 0 {
		for _, option := range options {
			option(req)
		}
	}
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}
	resp, err := req.PUT(url, body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(resp, reply); err != nil {
		return err
	}
	return nil
}

func Delete(url string, data any, reply any, options ...Option) error {
	if len(options) > 0 {
		for _, option := range options {
			option(req)
		}
	}
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}
	resp, err := req.DELETE(url, body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(resp, reply); err != nil {
		return err
	}
	return nil
}
