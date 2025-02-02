// Package amutool
// Date: 2022/9/7 10:12
// Author: Amu
// Description:
package basex

import (
	"fmt"
	"testing"
)

func TestBase64Encode(t *testing.T) {
	str := "hello world"
	res := Encode(str)
	fmt.Println(res)
}

func TestBase64Decode(t *testing.T) {
	str := "aGVsbG8gd29ybGQ="
	res, _ := Decode(str)
	fmt.Println(res)
}
