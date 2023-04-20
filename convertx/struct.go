// Package convertx
// Date: 2022/10/1 02:07
// Author: Amu
// Description:
package convertx

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func StructToMap(content interface{}) map[string]interface{} {
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

func StructToJson(content interface{}) interface{} {
	if marshalContent, err := json.Marshal(content); err != nil {
		return nil
	} else {
		return string(marshalContent)
	}
}
