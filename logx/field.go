// Package logx
// Date: 2023/4/10 17:17
// Author: Amu
// Description:
package logx

import (
	"go.uber.org/zap"
)

func Field(key string, value interface{}) zap.Field {
	return zap.Any(key, value)
}
