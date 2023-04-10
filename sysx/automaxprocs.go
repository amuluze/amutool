// Package sysx
// Date: 2023/4/10 14:56
// Author: Amu
// Description:
package sysx

import "go.uber.org/automaxprocs/maxprocs"

func init() {
	maxprocs.Set(maxprocs.Logger(nil))
}
