// Package iohelper
// Date: 2023/4/20 15:43
// Author: Amu
// Description:
package iohelper

import (
	"fmt"
	"testing"
)

func TestTextMd5Change(t *testing.T) {
	textMD5 := NewCheckTextMd5("./test.txt", "", ".", "123")
	res := textMD5.Change()
	fmt.Printf("change res: %v\n", res)
}
