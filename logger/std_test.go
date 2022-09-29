// Package logger
// Date: 2022/9/29 01:10
// Author: Amu
// Description:
package logger

import (
	"testing"

	"github.com/pkg/errors"
)

type User struct {
	Name string
}

func TestInfo(t *testing.T) {
	err := errors.New("bad request")
	user := &User{Name: "amu"}
	Info("hello", "status", 200, user.Name, err)
}
