// Package hash
// Date: 2022/9/7 10:26
// Author: Amu
// Description:
package hash

import (
	"fmt"
	"testing"
)

func TestMD5(t *testing.T) {
	str := "hello world"
	res := MD5([]byte(str))
	fmt.Println(res)
}

func TestMD5String(t *testing.T) {
	str := "hello world"
	res := MD5String(str)
	fmt.Println(res)
}

func TestSHA1(t *testing.T) {
	str := "123456"
	res := SHA1([]byte(str))
	fmt.Println(res)
}

func TestSHA1String(t *testing.T) {
	str := "123456"
	res := SHA1String(str)
	fmt.Println(res)
}
