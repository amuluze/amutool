// Package docker
// Date: 2023/4/19 16:35
// Author: Amu
// Description:
package docker

import (
	"context"
	"fmt"
	"testing"
)

func TestListContainer(t *testing.T) {
	manager, _ := NewManager()
	containers, _ := manager.ListContainer(context.Background())
	for _, c := range containers {
		fmt.Printf("%#v\n", c)
	}
}
