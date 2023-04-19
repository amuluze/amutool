// Package hashx
// Date: 2022/9/7 09:45
// Author: Amu
// Description:
package hashx

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

func MD5(b []byte) string {
	md5h := md5.New()
	_, _ = md5h.Write(b)
	return fmt.Sprintf("%x", md5h.Sum(nil))
}

func MD5String(s string) string {
	return MD5([]byte(s))
}

func SHA1(b []byte) string {
	md5h := sha1.New()
	_, _ = md5h.Write(b)
	return fmt.Sprintf("%x", md5h.Sum(nil))
}

func SHA1String(s string) string {
	return SHA1([]byte(s))
}
