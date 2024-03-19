// Package docker
// Date: 2023/4/19 17:21
// Author: Amu
// Description:
package docker

import (
	"context"
	"fmt"
	"testing"
)

func TestListNetwork(t *testing.T) {
	manager, _ := NewManager()
	nets, _ := manager.ListNetwork(context.Background())
	for _, net := range nets {
		fmt.Printf("net: %#v\n", net)
	}
}

func TestQueryNetwork(t *testing.T) {
	manager, _ := NewManager()
	net, _ := manager.QueryNetwork(context.Background(), "a5e3e066520b")
	fmt.Printf("net detail: %#v\n", net)
}
