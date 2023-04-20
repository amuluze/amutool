// Package docker
// Date: 2023/4/19 14:19
// Author: Amu
// Description: docker 容器操作
package docker

import (
	"context"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/tidwall/gjson"
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

// CreateContainer 根据条件创建容器（各种条件会比较复杂），创建成功后返回 containerID，此时容器状态为 created
func (m *Manager) CreateContainer(ctx context.Context, imageTag string, networkID string, containerName string) (string, error) {
	containerConfig := &container.Config{
		Image: imageTag,
	}
	//hostConfig := &container.HostConfig{
	//
	//}
	//networkConfig := &network.NetworkingConfig{
	//
	//}
	//platform := &v1.Platform{
	//
	//}
	createResponse, err := m.Client.ContainerCreate(ctx, containerConfig, nil, nil, nil, containerName)
	if err != nil {
		return "", err
	}
	return createResponse.ID, nil
}

// StartContainer 根据 containerID 启动容器
func (m *Manager) StartContainer(ctx context.Context, containerID string) error {
	return m.Client.ContainerStart(ctx, containerID, types.ContainerStartOptions{})
}

// StopContainer stop 指定容器
func (m *Manager) StopContainer(ctx context.Context, containerID string) error {
	return m.Client.ContainerStop(ctx, containerID, container.StopOptions{})
}

// RestartContainer 重启指定容器
func (m *Manager) RestartContainer(ctx context.Context, containerID string) error {
	return m.Client.ContainerRestart(ctx, containerID, container.StopOptions{})
}

// DeleteContainer 删除指定容器
func (m *Manager) DeleteContainer(ctx context.Context, containerID string) error {
	return m.Client.ContainerRemove(ctx, containerID, types.ContainerRemoveOptions{Force: true, RemoveVolumes: false, RemoveLinks: false})
}

// CopyFileToContainer 向容器中拷贝文件
func (m *Manager) CopyFileToContainer(ctx context.Context, containerID, srcFile, dstFile string) error {
	file, _ := os.Open(srcFile)
	return m.Client.CopyToContainer(ctx, containerID, dstFile, file, types.CopyToContainerOptions{AllowOverwriteDirWithFile: true, CopyUIDGID: false})
}

// GetContainerMem 获取指定容器的内存使用情况，单位 MB
func (m *Manager) GetContainerMem(ctx context.Context, containerID string) (int64, error) {
	stats, _ := m.Client.ContainerStats(ctx, containerID, false)
	body, _ := ioutil.ReadAll(stats.Body)
	memUsage := gjson.Get(string(body), "memory_stats.usage").Int() / 1024 / 1024
	memCache := gjson.Get(string(body), "memory_stats.stats.cache").Int() / 1024 / 1024
	mem := memUsage - memCache
	return mem, nil
}
