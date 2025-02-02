// Package requests
// Date: 2024/12/16 15:22:00
// Author: Amu
// Description:
package requests

import (
	"fmt"
	"net/url"
	"reflect"
)

func ToQuery(params any) string {
	if params == nil {
		return ""
	}
	values := url.Values{}
	v := reflect.ValueOf(params).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		jsonTag := field.Tag.Get("json")
		value := v.Field(i).Interface()
		values.Add(jsonTag, fmt.Sprintf("%v", value))
	}

	return values.Encode()
}
