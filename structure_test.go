// Package amutool
// Date: 2022/9/7 11:16
// Author: Amu
// Description:
package amutool

import (
	"fmt"
	"testing"
)

type TestObjectOne struct {
	Name string
	Age  int
}

type TestObjectTwo struct {
	Name   string
	gender string
}

func TestCopy(t *testing.T) {

	too := &TestObjectOne{Name: "amu", Age: 12}
	tot := &TestObjectTwo{}
	Copy(too, tot)
	fmt.Printf("tot: %+v", tot)
}
