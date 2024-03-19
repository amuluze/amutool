// Package docker
// Date: 2023/4/19 14:19
// Author: Amu
// Description: docker 网络操作
package docker

import (
	"context"
	"strings"

	"github.com/docker/docker/api/types/network"

	"github.com/docker/docker/api/types"
)

type Network struct {
	ID         string
	Name       string
	Driver     string
	Scope      string
	Created    string
	Internal   bool
	IPAM       network.IPAM
	Containers map[string]string // map[cid]ipaddr
}

func (m *Manager) ListNetwork(ctx context.Context) ([]Network, error) {
	nets, err := m.Client.NetworkList(ctx, types.NetworkListOptions{})
	if err != nil {
		return nil, err
	}

	var networkList []Network
	for _, net := range nets {
		var containers map[string]string
		for id, container := range net.Containers {
			ipAddr := container.IPv4Address
			if slashIdx := strings.IndexByte(ipAddr, '/'); slashIdx != -1 {
				ipAddr = ipAddr[:slashIdx]
			}
			containers[id] = ipAddr
		}
		n := Network{
			ID:         net.ID,
			Name:       net.Name,
			Driver:     net.Driver,
			Scope:      net.Scope,
			Created:    net.Created.Format("2006-01-02 15:04:05"),
			IPAM:       net.IPAM,
			Containers: containers,
		}
		networkList = append(networkList, n)
	}
	return networkList, nil
}

func (m *Manager) QueryNetwork(ctx context.Context, networkID string) (*Network, error) {
	nr, err := m.Client.NetworkInspect(ctx, networkID, types.NetworkInspectOptions{})
	if err != nil {
		return nil, err
	}
	var containers map[string]string
	for id, container := range nr.Containers {
		ipAddr := container.IPv4Address
		if slashIdx := strings.IndexByte(ipAddr, '/'); slashIdx != -1 {
			ipAddr = ipAddr[:slashIdx]
		}
		containers[id] = ipAddr
	}
	nw := &Network{
		ID:         nr.ID,
		Name:       nr.Name,
		Driver:     nr.Driver,
		Scope:      nr.Scope,
		Created:    nr.Created.Format("2006-01-02 15:04:05"),
		IPAM:       nr.IPAM,
		Containers: containers,
	}
	return nw, nil
}

func (m *Manager) CreateNetwork(ctx context.Context, name string, internal bool) (string, error) {
	response, err := m.Client.NetworkCreate(ctx, name, types.NetworkCreate{
		CheckDuplicate: true,
		Driver:         "bridge",
		EnableIPv6:     false,
		Internal:       internal,
	})
	if err != nil {
		return "", err
	}
	return response.ID, nil
}

func (m *Manager) UpdateNetwork(ctx context.Context) error { return nil }

func (m *Manager) DeleteNetwork(ctx context.Context, networkID string) error {
	err := m.Client.NetworkRemove(ctx, networkID)
	if err != nil {
		return err
	}
	return nil
}
