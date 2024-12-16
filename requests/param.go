// Package requests
// Date:   2024/12/13 18:22
// Author: Amu
// Description:
package requests

type FormData interface {
	ToQuery() string
}

type JsonData interface{}
