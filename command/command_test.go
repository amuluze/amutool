// Package command
// Date: 2023/4/6 15:39
// Author: Amu
// Description:
package command

import (
	"context"
	"fmt"
	"testing"
)

func TestRunCommand(t *testing.T) {
	if res, err := RunCommand(context.Background(), "pwd"); err != nil {
		fmt.Printf("error: %v\n", err)
		return
	} else {
		fmt.Printf("res: %v\n", string(res))
	}
}