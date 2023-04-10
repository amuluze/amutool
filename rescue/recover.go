// Package rescue
// Date: 2023/4/10 14:16
// Author: Amu
// Description:
package rescue

import "fmt"

func Recover(cleanups ...func()) {
	for _, cleanup := range cleanups {
		cleanup()
	}

	if p := recover(); p != nil {
		fmt.Println(p)
	}
}
