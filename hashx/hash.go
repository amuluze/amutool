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

func MD5(b string) string {
	md5h := md5.New()
	_, _ = md5h.Write([]byte(b))
	return fmt.Sprintf("%x", md5h.Sum(nil))
}

func SHA1(b string) string {
	md5h := sha1.New()
	_, _ = md5h.Write([]byte(b))
	return fmt.Sprintf("%x", md5h.Sum(nil))
}
