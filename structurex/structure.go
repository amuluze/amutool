// Package structurex
// Date: 2022/9/7 09:50
// Author: Amu
// Description:
package structurex

import (
	"github.com/jinzhu/copier"
)

// Copy 结构体映射，只会映射结构体中交集字段的值
func Copy(s, ts interface{}) error {
	return copier.Copy(ts, s)
}
