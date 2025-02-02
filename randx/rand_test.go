// Package amutool
// Date: 2022/9/23 13:38
// Author: Amu
// Description:
package randx

import (
	"fmt"
	"testing"
)

func TestRandInt(t *testing.T) {
	res := RandInt(1, 100)
	fmt.Println(res)
}

func TestRandNumeral(t *testing.T) {
	res := RandNumeral(3)
	fmt.Println(res)
}

func TestRandString(t *testing.T) {
	res := RandString(3)
	fmt.Println(res)
}

func TestRandNumeralOrLetter(t *testing.T) {
	res := RandNumeralOrLetter(4)
	fmt.Println(res)
}
