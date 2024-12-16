// Package requests
// Date:   2024/12/13 18:15
// Author: Amu
// Description:
package requests

import "testing"

func TestGET(t *testing.T) {
	requests := NewRequests()
	requests.GET("http://example.com", "param1=value1&param2=value2")
}
