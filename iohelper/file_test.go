// Package iohelper
// Date: 2023/3/28 17:20
// Author: Amu
// Description:
package iohelper

import (
	"fmt"
	"testing"
)

func TestFileExist(t *testing.T) {
	path := "/Users/amu/Desktop/amuluze/amutool/base64.go"
	res := FileExist(path)
	fmt.Printf("res: %#v\n", res)
}
