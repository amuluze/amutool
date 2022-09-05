// Package requests
// Date: 2022/9/5 13:42
// Author: Amu
// Description:
package requests

import "time"

const (
	localUserAgent      = "Requests/0.10"
	dialTimeout         = 30 * time.Second
	dialKeepAlive       = 30 * time.Second
	tlshandshakeTimeout = 10 * time.Second
	requestTimeout      = 30 * time.Second
)
