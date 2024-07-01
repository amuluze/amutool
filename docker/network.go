// Package docker
// Date: 2023/4/19 14:19
// Author: Amu
// Description: docker 网络操作
package docker

import (
	"context"
	"fmt"
	"net"
	"strings"

	"github.com/docker/docker/api/types/filters"
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
	SubNet     []SubNetworkConfig
	Containers map[string]string // map[cid]ipaddr
	Labels     map[string]string
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
		subNet := make([]SubNetworkConfig, 0)
		for _, ncf := range net.IPAM.Config {
			subNet = append(subNet, SubNetworkConfig{
				Subnet:  ncf.Subnet,
				Gateway: ncf.Gateway,
			})
		}
		n := Network{
			ID:         net.ID,
			Name:       net.Name,
			Driver:     net.Driver,
			Scope:      net.Scope,
			Created:    net.Created.Format("2006-01-02 15:04:05"),
			SubNet:     subNet,
			Containers: containers,
			Labels:     net.Labels,
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
	subNet := make([]SubNetworkConfig, 0)
	for _, ncf := range nr.IPAM.Config {
		subNet = append(subNet, SubNetworkConfig{
			Subnet:  ncf.Subnet,
			Gateway: ncf.Gateway,
		})
	}
	nw := &Network{
		ID:         nr.ID,
		Name:       nr.Name,
		Driver:     nr.Driver,
		Scope:      nr.Scope,
		Created:    nr.Created.Format("2006-01-02 15:04:05"),
		SubNet:     subNet,
		Containers: containers,
		Labels:     nr.Labels,
	}
	return nw, nil
}

// CreateNetwork creates a new network.
func (m *Manager) CreateNetwork(ctx context.Context, name string, driver string, networkSegment string, labels map[string]string) (string, error) {
	var options types.NetworkCreate
	if networkSegment != "" {
		// 根据网段 networkSegment 计算网关
		ip := net.ParseIP(strings.Split(networkSegment, "/")[0])
		mask := net.CIDRMask(24, 32)
		nw := ip.Mask(mask)
		gateway := net.IPv4(nw[0], nw[1], nw[2], nw[3]+1).String()

		ipam := &network.IPAM{
			Driver:  "default",
			Options: nil,
			Config: []network.IPAMConfig{
				{
					Subnet:  networkSegment,
					Gateway: gateway,
				},
			},
		}
		options = types.NetworkCreate{
			Driver:     driver,
			EnableIPv6: false,
			Internal:   false,
			IPAM:       ipam,
		}
	} else {
		options = types.NetworkCreate{
			Driver:     driver,
			EnableIPv6: false,
			Internal:   true,
			Options:    map[string]string{"com.docker.network.bridge.name": name},
		}
	}
	if len(labels) != 0 {
		options.Labels = labels
	}
	fmt.Printf("options: %#v\n", options.IPAM)
	response, err := m.Client.NetworkCreate(ctx, name, options)
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
