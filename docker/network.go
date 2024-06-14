// Package docker
// Date: 2023/4/19 14:19
// Author: Amu
// Description: docker 网络操作
package docker

import (
	"context"
	"github.com/docker/docker/api/types/filters"
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

type SubNetworkConfig struct {
	Subnet  string
	Gateway string
}

// ListNetwork lists all networks.
func (m *Manager) ListNetwork(ctx context.Context) ([]Network, error) {
	nets, err := m.Client.NetworkList(ctx, types.NetworkListOptions{})
	if err != nil {
		return nil, err
	}

	var networkList []Network
	for _, net := range nets {
		containers := make(map[string]string)
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

// QueryNetwork queries a network by its ID.
func (m *Manager) QueryNetwork(ctx context.Context, networkID string) (*Network, error) {
	nr, err := m.Client.NetworkInspect(ctx, networkID, types.NetworkInspectOptions{})
	if err != nil {
		return nil, err
	}
	containers := make(map[string]string)
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

// CreateNetwork creates a new network.
func (m *Manager) CreateNetwork(ctx context.Context, name string, driver string) (string, error) {
	//ipam := &network.IPAM{
	//	Driver:  "default",
	//	Options: nil,
	//	Config: []network.IPAMConfig{
	//		{
	//			Subnet:  "172.18.0.0/16",
	//			Gateway: "172.18.0.1",
	//		},
	//	},
	//}
	//if len(subConfig) > 0 {
	//	conf := make([]network.IPAMConfig, 0, len(subConfig))
	//	for sc := range subConfig {
	//		conf = append(conf, network.IPAMConfig{
	//			Subnet:  subConfig[sc].Subnet,
	//			Gateway: subConfig[sc].Gateway,
	//		})
	//	}
	//	ipam.Config = conf
	//}
	response, err := m.Client.NetworkCreate(ctx, name, types.NetworkCreate{
		Driver:     driver,
		EnableIPv6: false,
		Internal:   true,
	})
	if err != nil {
		return "", err
	}
	return response.ID, nil
}

// DeleteNetwork removes a network by network id.
func (m *Manager) DeleteNetwork(ctx context.Context, networkID string) error {
	return m.Client.NetworkRemove(ctx, networkID)
}

// PruneNetwork removes all dangling networks.
func (m *Manager) PruneNetwork(ctx context.Context) error {
	_, err := m.Client.NetworksPrune(ctx, filters.NewArgs(filters.Arg("dangling", "true")))
	return err
}
