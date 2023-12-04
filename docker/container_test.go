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
	mem, err := manager.GetContainerMem(context.Background(), "dcf6136614908235918c7d2b20e0d619b69ad446fb32c2274bf6c159cc573412")
	if err != nil {
		panic(err)
	}
	fmt.Printf("container mem: %v MB\n", mem)
}

func TestContainerCPU(t *testing.T) {
	manager, _ := NewManager()
	cpu, err := manager.GetContainerCPU(context.Background(), "dcf6136614908235918c7d2b20e0d619b69ad446fb32c2274bf6c159cc573412")
	if err != nil {
		panic(err)
	}
	fmt.Printf("cpu percent: %v\n", cpu)
}
