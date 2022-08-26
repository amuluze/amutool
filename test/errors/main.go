// Package main
// Date: 2022/8/26 15:46
// Author: Amu
// Description:
package main

import (
	"fmt"

	"gitee.com/amuluze/amutool/errors"
)

func main() {
	err := errors.New("new error")
	fmt.Println(err)

	err = errors.Wrap400Response(err, "400 error")
	fmt.Println(err)
}
