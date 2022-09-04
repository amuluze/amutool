// Package requests
// Date: 2022/9/1 23:06
// Author: Amu
// Description:
package requests

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

type Request struct {
	Client  *http.Client
	Headers map[string]string // header
	Cookies map[string]string // cookies
	Data    map[string]string // data
	JSON    interface{}       // json
	Params  map[string]string // params
	Timeout time.Duration     //timeout
	Files   []File            // file
}

func NewRequest(options ...Options) *Request {
	r := Request{}
	for _, option := range options {
		option(&r)
	}
	return &r
}

func DoRequest(method, url string, options ...Options) (*Response, error) {
	return buildResponse(buildRequest(method, url, options...))
}

func buildRequest(method, url string, options ...Options) (*http.Response, error) {
	request := NewRequest(options...)
	client = buildClient(*request)
}

func buildClient(r Request) *http.Client {
	return &http.Client{
		Timeout: r.Timeout,
	}
}

func createTransport(r Request) *http.Transport {
	ourHTTPTransport := &http.Transport{
		// These are borrowed from the default transporter
		Proxy: http.ProxyFromEnvironment,
		Dial: (&net.Dialer{
			Timeout:   ro.DialTimeout,
			KeepAlive: ro.DialKeepAlive,
			LocalAddr: ro.LocalAddr,
		}).Dial,
		TLSHandshakeTimeout: ro.TLSHandshakeTimeout,

		// Here comes the user settings
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: ro.InsecureSkipVerify},
		DisableCompression: ro.DisableCompression,
	}
	EnsureTransporterFinalized(ourHTTPTransport)
	return ourHTTPTransport
}
