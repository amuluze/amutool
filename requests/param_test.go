// Package requests
// Date:   2024/12/13 18:22
// Author: Amu
// Description:
package requests

import (
	"fmt"
	"net/url"
	"reflect"
	"testing"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (u *User) ToQuery() string {
	values := url.Values{}
	v := reflect.ValueOf(u).Elem()
	t := v.Type()
	fmt.Println(v, t)

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		jsonTag := field.Tag.Get("json")
		value := v.Field(i).Interface()
		values.Add(jsonTag, fmt.Sprintf("%v", value))
	}

	return values.Encode()
}

func TestFormData(t *testing.T) {
	user := &User{Name: "jack", Age: 18}
	t.Logf("user to query: %s", user.ToQuery())
}
