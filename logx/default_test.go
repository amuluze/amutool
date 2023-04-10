// Package logx
// Date: 2023/4/10 17:52
// Author: Amu
// Description:
package logx

import "testing"

func TestDebug(t *testing.T) {
	defaultLogger.Debug("hello golang")
}
