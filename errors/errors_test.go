// Package errors
// Date: 2023/11/23 15:23
// Author: Amu
// Description:
package errors

import (
	"fmt"
	"testing"
)

func TestNewError(t *testing.T) {
	err := NewError(400, "bad request")
	fmt.Println(err.Error(), err.CodeStatus())
}
