// Package requests
// Date: 2022/9/7 23:56
// Author: Amu
// Description:
package requests

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	params := map[string]string{
		"Good": "Job",
	}

	resp, err := Get("http://httpbin.org/get?Hello=World", SetParam(params))
	if err != nil {
		return
	}

	fmt.Printf("resp: %#v\n", resp)
	fmt.Printf("response ok: %+v\n", resp.Ok)
	fmt.Printf("response err: %+v\n", resp.Error)
	fmt.Printf("response raw: %+v\n", resp.RawResponse)
	fmt.Printf("response status: %+v\n", resp.StatusCode)
	fmt.Printf("response header: %+v\n", resp.Header)
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>><<<<<<<<<<<<<<<<<")

	fmt.Println(resp.RawResponse.Body)
}
