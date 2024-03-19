// Package es
// Date: 2023/3/28 14:19
// Author: Amu
// Description:
package es

import (
	"os"
)

func (c *Client) ReadFile(fileName string) (string, error) {
	_, err := os.Stat(fileName)
	if err != nil {
		return "", err
	}

	body, _ := os.ReadFile(fileName)
	return string(body), nil
}
