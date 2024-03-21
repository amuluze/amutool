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

func TestContainerMem(t *testing.T) {
	manager, _ := NewManager()
	percent, used, limit, err := manager.GetContainerMem(context.Background(), "dc505c86389c")
	if err != nil {
		panic(err)
	}
	fmt.Printf("container mem percent: %v, used: %v, limit: %v \n", percent, used, limit)
}

func TestContainerCPU(t *testing.T) {
	manager, _ := NewManager()
	cpu, err := manager.GetContainerCPU(context.Background(), "dc505c86389c")
	if err != nil {
		panic(err)
	}
	fmt.Printf("cpu percent: %v\n", cpu)
}
