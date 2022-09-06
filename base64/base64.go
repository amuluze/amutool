// Package base64
// Date: 2022/9/7 01:01
// Author: Amu
// Description:
package base64

import "encoding/base64"

// Base64Encode base64 编码
func Base64Encode(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

// Base64Decode base64 解码
func Base64Decode(src string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(src)
}
