// Package docker
// Date: 2023/4/19 14:19
// Author: Amu
// Description: docker 容器操作
package docker

import (
	"context"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
)

type ContainerSummary struct {
	ID      string `json:"id"`      // ID
	Name    string `json:"name"`    // Name
	Image   string `json:"image"`   // Image
	State   string `json:"state"`   // State: created running paused restarting removing exited dead
	Created string `json:"created"` // create time
	Uptime  string `json:"uptime"`  // uptime in seconds
}

type Container struct {
}

func (m *Manager) getUptime(ctx context.Context, containerID string) string {
	inspect, _ := m.Client.ContainerInspect(ctx, containerID)
	started, _ := time.Parse(time.RFC3339Nano, inspect.State.StartedAt)
	return started.Format("2006-01-02 15:04:05")
}

// ListContainer 获取所有容器 []ContainerSummary
func (m *Manager) ListContainer(ctx context.Context) ([]ContainerSummary, error) {
	containers, err := m.Client.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		return nil, err
	}

	var containerSummaryList []ContainerSummary
	for _, c := range containers {
		var uptime string
		if c.State == "running" {
			uptime = m.getUptime(ctx, c.ID)
		}
		cs := ContainerSummary{
			ID:      c.ID,
			Name:    strings.Trim(c.Names[0], "/"),
			Image:   c.Image,
			State:   c.State,
			Created: time.Unix(c.Created, 0).Format("2006-01-02 15:04:05"),
			Uptime:  uptime,
		}
		containerSummaryList = append(containerSummaryList, cs)
	}
	return containerSummaryList, nil
}

func (m *Manager) CreateContainer(ctx context.Context, imageTag string, networkID string, imageID string) (string, error) {
	return "", nil
}
