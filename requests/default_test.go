// Package requests
// Date:   2024/12/13 18:27
// Author: Amu
// Description:
package requests

import "testing"

func TestGet(t *testing.T) {
	user := &User{Name: "jack", Age: 18}
	Get("http://example.com", user, nil)
}
