// Package uuidx
// Date: 2022/9/7 11:04
// Author: Amu
// Description:
package uuidx

import (
	"fmt"
	"testing"
)

func TestNewUUID(t *testing.T) {
	res, _ := NewUUID()
	fmt.Println(res)
}

func TestMustUUID(t *testing.T) {
	res := MustUUID()
	fmt.Println(res)
}

func TestMustString(t *testing.T) {
	res := MustString()
	fmt.Println(res)
}

func TestMustParseUUIToString(t *testing.T) {
	str := "422ba45f-a9f8-4ec2-b048-eb9d902df5ad"
	res := MustParseUUIToString(str)
	fmt.Println(res)
}

func TestSnowID(t *testing.T) {
	res := SnowID()
	fmt.Println(res)
}
