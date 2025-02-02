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

type Version struct {
	DockerVersion string `json:"docker_version"`
	APIVersion    string `json:"api_version"`
	MinAPIVersion string `json:"min_api_version"`
	GitCommit     string `json:"git_commit"`
	GoVersion     string `json:"go_version"`
	OS            string `json:"os"`
	Arch          string `json:"arch"`
}

func NewManager() (*Manager, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	return &Manager{Client: cli}, err
}

func (m *Manager) Version(ctx context.Context) (*Version, error) {
	serverVersion, err := m.Client.ServerVersion(ctx)
	if err != nil {
		return nil, err
	}

	return &Version{
		DockerVersion: serverVersion.Version,
		APIVersion:    serverVersion.APIVersion,
		MinAPIVersion: serverVersion.MinAPIVersion,
		GitCommit:     serverVersion.GitCommit,
		GoVersion:     serverVersion.GoVersion,
		OS:            serverVersion.Os,
		Arch:          serverVersion.Arch,
	}, nil
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
