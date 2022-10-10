// Package convert
// Date: 2022/10/1 02:13
// Author: Amu
// Description:
package convert

import (
	"fmt"
	"testing"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestToJson(t *testing.T) {
	user := &User{Name: "amu", Age: 32}
	res := ToJson(user)
	fmt.Printf("%s\n", res)
}

func TestToMap(t *testing.T) {
	user := &User{Name: "amu", Age: 32}
	res := ToMap(user)
	fmt.Printf("%s\n", res)
}
