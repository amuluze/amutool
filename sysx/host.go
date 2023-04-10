// Package sysx
// Date: 2023/4/10 14:24
// Author: Amu
// Description:
package sysx

import (
	"os"

	"gitee.com/amuluze/amutool/stringx"
)

var hostname string

func init() {
	var err error
	hostname, err = os.Hostname()
	if err != nil {
		hostname = stringx.RandId()
	}
}

func Hostname() string {
	return hostname
}
