// Package docker
// Date: 2023/4/19 17:21
// Author: Amu
// Description:
package docker

import (
	"context"
	"testing"
)

func TestListNetwork(t *testing.T) {
	manager, _ := NewManager()
	nets, _ := manager.ListNetwork(context.Background())
	for _, net := range nets {
		t.Logf("network: %#v\n", net)
	}
}

func TestQueryNetwork(t *testing.T) {
	manager, _ := NewManager()
	net, _ := manager.QueryNetwork(context.Background(), "e01b3ce4efdccc7efdffdfffeb39136f00ef9b9d6ea7c1c1098b08472e0e22d2")
	t.Logf("network detail: %#v\n", net)
}

func TestCreateNetwork(t *testing.T) {
	manager, _ := NewManager()
	network, err := manager.CreateNetwork(context.Background(), "test-network", "bridge", true)
	if err != nil {
		t.Fatalf("create network failed: %v\n", err)
	}
	t.Logf("network: %#v\n", network)
}

func TestDeleteNetwork(t *testing.T) {
	manager, _ := NewManager()
	err := manager.DeleteNetwork(context.Background(), "e01b3ce4efdccc7efdffdfffeb39136f00ef9b9d6ea7c1c1098b08472e0e22d2")
	if err != nil {
		t.Errorf("delete network failed: %v\n", err)
	}
	t.Log("delete network success")
}

func TestPruneNetwork(t *testing.T) {
	manager, _ := NewManager()
	err := manager.PruneNetwork(context.Background())
	if err != nil {
		t.Errorf("prune network failed: %v\n", err)
	}
	t.Log("prune network success")
}
