// Package requests
// Date: 2022/9/5 13:42
// Author: Amu
// Description:
package requests

import "time"

const (
	localUserAgent      = "Requests/0.10"  // userAgent
	requestTimeout      = 30 * time.Second // 请求超时
	maxConnsPerHost     = 10               // 某一host的最大连接数
	maxIdleConnsPerHost = 3                // 某一host的最大空闲连接数
	idleConnTimeout     = 10 * time.Second // 空闲连接超时时间
)
