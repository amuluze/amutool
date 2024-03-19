// Package docker
// Date: 2023/4/19 16:58
// Author: Amu
// Description:
package docker

import (
	"context"
	"fmt"
	"testing"
)

func TestListImage(t *testing.T) {
	manager, _ := NewManager()
	images, _ := manager.ListImage(context.Background())
	for _, image := range images {
		fmt.Printf("image: %#v\n", image)
	}
}
