// Package amutool
// Date: 2022/9/7 10:26
// Author: Amu
// Description:
package hashx

import (
	"fmt"
	"testing"
)

func TestMD5(t *testing.T) {
	str := "hello world"
	res := MD5(str)
	fmt.Println(res)
}

func TestSHA1(t *testing.T) {
	str := "123456"
	res := SHA1(str)
	fmt.Println(res)
}
