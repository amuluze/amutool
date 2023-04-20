// Package basex
// Date: 2022/9/7 01:01
// Author: Amu
// Description:
package basex

import "encoding/base64"

// Encode base64 编码
func Encode(src string) string {
	return base64.StdEncoding.EncodeToString([]byte(src))
}

// Decode base64 解码
func Decode(src string) (string, error) {
	if res, err := base64.StdEncoding.DecodeString(src); err != nil {
		return "", err
	} else {
		return string(res), nil
	}
}
