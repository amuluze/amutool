// Package gormx
// Date: 2022/9/30 17:34
// Author: Amu
// Description:
package gormx

import (
	"gorm.io/gorm"
)

func Exec(fn func() error) error {
	return db.Transaction(func(tx *gorm.DB) error {
		return fn()
	})
}
