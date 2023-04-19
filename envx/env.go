// Package envx
// Date: 2022/9/7 01:02
// Author: Amu
// Description:
package envx

import (
	"os"
)

// GetEnv 获取本地系统变量
func GetEnv(key string) string {
	return os.Getenv(key)
}
