// Package convert
// Date: 2022/10/1 02:07
// Author: Amu
// Description:
package convert

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func ToMap(content interface{}) map[string]interface{} {
	var result map[string]interface{}
	if marshalContent, err := json.Marshal(content); err != nil {
		fmt.Println(err)
	} else {
		d := json.NewDecoder(bytes.NewReader(marshalContent))
		d.UseNumber()
		if err := d.Decode(&result); err != nil {
			fmt.Println(err)
		} else {
			for k, v := range result {
				result[k] = v
			}
		}
	}
	return result
}

func ToJson(content interface{}) interface{} {
	if marshalContent, err := json.Marshal(content); err != nil {
		return nil
	} else {
		return string(marshalContent)
	}
}
