// Package amutool
// Date: 2022/9/23 13:38
// Author: Amu
// Description:
package randx

import (
	"fmt"
	"testing"
)

func TestGetRandInt(t *testing.T) {
	res := GetRandInt(1, 100)
	fmt.Println(res)
}

func TestGetRangeNum(t *testing.T) {
	res := GetRangeNum(3)
	fmt.Println(res)
}

func TestGetRandomString(t *testing.T) {
	res := RandomString(3)
	fmt.Println(res)
}

func TestGetRangeNumString(t *testing.T) {
	res := GetRangeNumString(4)
	fmt.Println(res)
}
