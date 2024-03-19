// Package docker
// Date: 2023/4/19 15:29
// Author: Amu
// Description:
package docker

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type Manager struct {
	Client *client.Client
}

func NewManager() (*Manager, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	return &Manager{Client: cli}, err
}

func (m *Manager) JoinNetwork(ctx context.Context, containerID string, networkID string) error {
	// check if network is existing
	var err error
	_, err = m.Client.NetworkInspect(ctx, networkID, types.NetworkInspectOptions{})
	if err != nil {
		return err
	}

	// check if container is existing
	_, err = m.Client.ContainerInspect(ctx, containerID)
	if err != nil {
		return err
	}

	return m.Client.NetworkConnect(ctx, networkID, containerID, nil)
}

func (m *Manager) LeaveNetwork(ctx context.Context, containerID string, networkID string) error {
	var err error
	_, err = m.Client.NetworkInspect(ctx, networkID, types.NetworkInspectOptions{})
	if err != nil {
		return err
	}

	// check if container is existing
	_, err = m.Client.ContainerInspect(ctx, containerID)
	if err != nil {
		return err
	}

	return m.Client.NetworkDisconnect(ctx, networkID, containerID, true)
}
