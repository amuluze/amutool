// Package docker
// Date: 2023/4/19 17:21
// Author: Amu
// Description:
package docker

import (
	"context"
	"net"
	"strings"
	"testing"
)

func TestIPParse(t *testing.T) {
	t.Logf("input: %s\n", strings.Split("172.80.0.0/24", "/")[0])

	ip := net.ParseIP(strings.Split("172.80.0.0/24", "/")[0])
	t.Logf("ip: %v\n", ip)
	mask := net.CIDRMask(24, 32)
	t.Logf("mask: %v\n", mask)
	nw := ip.Mask(mask)
	t.Logf("network: %v\n", nw)
	gateway := net.IPv4(nw[0], nw[1], nw[2], nw[3]+1).String()
	t.Logf("gateway: %s\n", gateway)
}

func TestListNetwork(t *testing.T) {
	manager, _ := NewManager()
	nets, _ := manager.ListNetwork(context.Background())
	for _, net := range nets {
		t.Logf("network: %#v\n", net)
	}
}

func TestQueryNetwork(t *testing.T) {
	manager, _ := NewManager()
	net, _ := manager.QueryNetwork(context.Background(), "7be8e024bcb58caff65d38b39e42dff05e292e3f2f30963ae51732250b45a33f")
	t.Logf("network detail: %#v\n", net)
}

func TestCreateNetwork(t *testing.T) {
	manager, _ := NewManager()
	network, err := manager.CreateNetwork(context.Background(), "test", "bridge", "172.20.0.0/24", map[string]string{AmprobeLabel: "true"})
	if err != nil {
		t.Fatalf("create network failed: %v\n", err)
	}
	t.Logf("network: %#v\n", network)
}

func TestCreateNetworkWithNetworkSegment(t *testing.T) {
	manager, _ := NewManager()
	network, err := manager.CreateNetwork(context.Background(), "test2", "bridge", "", map[string]string{AmprobeLabel: "true"})
	if err != nil {
		t.Fatalf("create network failed: %v\n", err)
	}
	t.Logf("network: %#v\n", network)
}

func TestDeleteNetwork(t *testing.T) {
	manager, _ := NewManager()
	err := manager.DeleteNetwork(context.Background(), "7be8e024bcb5")
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
