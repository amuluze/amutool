// Package requests
// Date: 2024/12/16 15:23:12
// Author: Amu
// Description:
package requests

import "testing"

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestToQuery(t *testing.T) {
	user := User{
		Name: "Amu",
		Age:  18,
	}
	res := ToQuery(&user)
	t.Logf("%v", res)
}
