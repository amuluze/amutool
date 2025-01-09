// Package docker
// Date: 2023/4/19 14:19
// Author: Amu
// Description: docker 容器操作
package docker

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
	"github.com/tidwall/gjson"
	goyaml "gopkg.in/yaml.v2"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/libcompose/yaml"
)

type ContainerSummary struct {
	ID      string            `json:"id"`      // ID
	Name    string            `json:"name"`    // Name
	Image   string            `json:"image"`   // Image
	State   string            `json:"state"`   // State: created running paused restarting removing exited dead
	Created string            `json:"created"` // create time
	Uptime  string            `json:"uptime"`  // uptime in seconds
	IP      string            `json:"ip"`      // ip
	Labels  map[string]string `json:"labels"`
}

type PortMapping struct {
	Proto         string
	IP            string
	HostPort      string
	ContainerPort string
}

// getUptime 获取指定容器的启动时间
func (m *Manager) getUptime(ctx context.Context, containerID string) string {
	inspect, _ := m.Client.ContainerInspect(ctx, containerID)
	started, _ := time.Parse(time.RFC3339Nano, inspect.State.StartedAt)
	return started.Format("2006-01-02 15:04:05")
}

// GetContainerIDByName 根据名称获取指定 container ID
func (m *Manager) GetContainerIDByName(ctx context.Context, name string) (string, error) {
	containers, err := m.Client.ContainerList(ctx, container.ListOptions{All: true})
	if err != nil {
		return "", err
	}
	for _, ct := range containers {
		if ct.Names[0] == name {
			return ct.ID, nil
		}
	}
	return "", nil
}

// GetContainerByName 根据名称获取指定 container
func (m *Manager) GetContainerByName(ctx context.Context, name string) (ContainerSummary, error) {
	containers, err := m.Client.ContainerList(ctx, container.ListOptions{All: true})
	if err != nil {
		return ContainerSummary{}, err
	}
	for _, ct := range containers {
		if ct.Names[0] == name {
			var uptime string
			if ct.State == "running" {
				uptime = m.getUptime(ctx, ct.ID)
			}

			var ip string
			for _, nt := range ct.NetworkSettings.Networks {
				if nt.IPAddress != "" {
					ip = nt.IPAddress
					break
				}
			}

			state := ct.State
			inspect, err := m.Client.ContainerInspect(ctx, ct.ID)
			if err == nil {
				if inspect.ContainerJSONBase.State.Health != nil && inspect.ContainerJSONBase.State.Health.Status == "healthy" {
					state = "running"
				}
			}
			return ContainerSummary{
				ID:      ct.ID,
				Name:    strings.Trim(ct.Names[0], "/"),
				Image:   ct.Image,
				State:   state,
				Created: time.Unix(ct.Created, 0).Format("2006-01-02 15:04:05"),
				Uptime:  uptime,
				IP:      ip,
				Labels:  ct.Labels,
			}, nil
		}
	}
	return ContainerSummary{}, errors.New("not found")
}

// ListContainer 获取所有容器 []ContainerSummary
func (m *Manager) ListContainer(ctx context.Context) ([]ContainerSummary, error) {
	containers, err := m.Client.ContainerList(ctx, container.ListOptions{All: true})
	if err != nil {
		return nil, err
	}

	var containerSummaryList []ContainerSummary
	for _, c := range containers {
		var uptime string
		if c.State == "running" {
			uptime = m.getUptime(ctx, c.ID)
		}

		var ip string
		for _, nt := range c.NetworkSettings.Networks {
			if nt.IPAddress != "" {
				ip = nt.IPAddress
				break
			}
		}

		state := c.State
		inspect, err := m.Client.ContainerInspect(ctx, c.ID)
		if err == nil {
			if inspect.ContainerJSONBase.State.Health != nil && inspect.ContainerJSONBase.State.Health.Status == "healthy" {
				state = "running"
			}
		}

		cs := ContainerSummary{
			ID:      c.ID,
			Name:    strings.Trim(c.Names[0], "/"),
			Image:   c.Image,
			State:   state,
			Created: time.Unix(c.Created, 0).Format("2006-01-02 15:04:05"),
			Uptime:  uptime,
			IP:      ip,
			Labels:  c.Labels,
		}
		containerSummaryList = append(containerSummaryList, cs)
	}
	return containerSummaryList, nil
}

// CreateContainer 根据条件创建容器（各种条件会比较复杂），创建成功后返回 containerID，此时容器状态为 created
func (m *Manager) CreateContainer(ctx context.Context, imageName string, networkID string, networkMode string, networkName string, containerName string, ports []string, vols []string, labels map[string]string) (string, error) {
	config := &container.Config{}
	config.Hostname = containerName
	config.Image = imageName
	config.Labels = labels
	config.Tty = true

	hostConfig := &container.HostConfig{}
	hostConfig.NetworkMode = container.NetworkMode(networkMode)
	hostConfig.RestartPolicy = container.RestartPolicy{Name: "always"}
	hostConfig.PortBindings = make(nat.PortMap)

	networkConfig := &network.NetworkingConfig{}
	networkConfig.EndpointsConfig = make(map[string]*network.EndpointSettings)
	networkConfig.EndpointsConfig[networkName] = &network.EndpointSettings{
		NetworkID: networkID,
	}

	for _, port := range ports {
		portsMapping, err := nat.ParsePortSpec(port)
		if err != nil {
			return "", err
		}
		for _, portMapping := range portsMapping {
			port, err := nat.NewPort(portMapping.Port.Proto(), portMapping.Port.Port())
			if err != nil {
				return "", err
			}
			hostIP := portMapping.Binding.HostIP
			if hostIP == "" {
				hostIP = "0.0.0.0"
			}
			hostConfig.PortBindings[port] = append(hostConfig.PortBindings[port], nat.PortBinding{
				HostIP:   hostIP,
				HostPort: portMapping.Binding.HostPort,
			})
		}
	}

	config.ExposedPorts = make(nat.PortSet)
	for port := range hostConfig.PortBindings {
		config.ExposedPorts[port] = struct{}{}
	}

	for _, vol := range vols {
		vol := "- " + vol
		volumes := &yaml.Volumes{}

		err := goyaml.Unmarshal([]byte(vol), volumes)
		if err != nil {
			return "", err
		}
		for _, volume := range volumes.Volumes {
			if volume.AccessMode != "ro" {
				volume.AccessMode = "rw"
			}
			volString := fmt.Sprintf("%s:%s:%s", volume.Source, volume.Destination, volume.AccessMode)
			hostConfig.Binds = append(hostConfig.Binds, volString)
		}
	}

	createResponse, err := m.Client.ContainerCreate(ctx, config, hostConfig, networkConfig, nil, containerName)
	if err != nil {
		return "", err
	}
	for _, w := range createResponse.Warnings {
		fmt.Printf("Container Create Warning: %s\n", w)
	}
	if err := m.Client.NetworkConnect(ctx, networkID, createResponse.ID, nil); err != nil {
		return "", err
	}

	if err := m.Client.ContainerStart(ctx, createResponse.ID, container.StartOptions{}); err != nil {
		return "", err
	}
	return createResponse.ID, nil
}

// StartContainer 根据 containerID 启动容器
func (m *Manager) StartContainer(ctx context.Context, containerID string) error {
	return m.Client.ContainerStart(ctx, containerID, container.StartOptions{})
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
	return m.Client.ContainerRemove(ctx, containerID, container.RemoveOptions{Force: true, RemoveVolumes: false, RemoveLinks: false})
}

// CopyFileToContainer 向容器中拷贝文件
func (m *Manager) CopyFileToContainer(ctx context.Context, containerID, srcFile, dstFile string) error {
	file, _ := os.Open(srcFile)
	return m.Client.CopyToContainer(ctx, containerID, dstFile, file, types.CopyToContainerOptions{AllowOverwriteDirWithFile: true, CopyUIDGID: false})
}

// GetContainerMem 获取指定容器的内存使用情况，单位 MB
func (m *Manager) GetContainerMem(ctx context.Context, containerID string) (float64, float64, float64, error) {
	stats, err := m.Client.ContainerStats(ctx, containerID, false)
	if err != nil {
		return 0.0, 0.0, 0.0, err
	}
	body, err := io.ReadAll(stats.Body)
	if err != nil {
		return 0.0, 0.0, 0.0, err
	}
	memUsage := gjson.Get(string(body), "memory_stats.usage").Float()
	memLimit := gjson.Get(string(body), "memory_stats.limit").Float()
	memPercent := (memUsage / memLimit) * 100
	return memPercent, memUsage, memLimit, nil
}

// GetContainerCPU 获取指定容器 CPU 使用情况，单位百分比
func (m *Manager) GetContainerCPU(ctx context.Context, containerID string) (float64, error) {
	stats, err := m.Client.ContainerStats(ctx, containerID, false)
	if err != nil {
		return 0.0, err
	}
	body, err := io.ReadAll(stats.Body)
	if err != nil {
		return 0.0, err
	}

	cpuDelta := gjson.Get(string(body), "cpu_stats.cpu_usage.total_usage").Float() - gjson.Get(string(body), "precpu_stats.cpu_usage.total_usage").Float()
	systemDelta := gjson.Get(string(body), "cpu_stats.system_cpu_usage").Float() - gjson.Get(string(body), "precpu_stats.system_cpu_usage").Float()
	cpuPercent := (cpuDelta / systemDelta) * 100.0
	return cpuPercent, nil
}

func (m *Manager) ContainerLogs(ctx context.Context, containerID string) (io.ReadCloser, error) {
	reader, err := m.Client.ContainerLogs(ctx, containerID, container.LogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Follow:     true,
		Timestamps: false,
		Tail:       "any",
	})
	return reader, err
}

func (m *Manager) RenameContainer(ctx context.Context, containerID, newName string) error {
	return m.Client.ContainerRename(ctx, containerID, newName)
}

// ExecCommand 在容器中执行命令
func (m *Manager) ExecCommand(ctx context.Context, containerID string, cmd []string) ([]byte, error) {
	create, err := m.Client.ContainerExecCreate(ctx, containerID, types.ExecConfig{
		AttachStdout: true,
		AttachStderr: true,
		Cmd:          cmd,
	})
	if err != nil {
		return nil, err
	}

	resp, err := m.Client.ContainerExecAttach(ctx, create.ID, types.ExecStartCheck{})
	if err != nil {
		return nil, err
	}
	defer resp.Close()

	output, err := io.ReadAll(resp.Reader)
	if err != nil {
		return nil, err
	}
	inspect, err := m.Client.ContainerExecInspect(ctx, create.ID)
	if err != nil {
		return nil, err
	}
	if inspect.ExitCode != 0 {
		return nil, fmt.Errorf("container exited with non-zero exit code: %d, output: %s", inspect.ExitCode, string(output))
	}
	return output, nil
}
