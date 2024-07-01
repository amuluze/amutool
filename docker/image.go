// Package docker
// Date: 2023/4/19 15:24
// Author: Amu
// Description:
package docker

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/registry"

	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/image"

	"github.com/pkg/errors"
)

type Image struct {
	ID      string
	Name    string
	Tag     string
	Created string
	Size    string
}

// ListImage 获取本地所有的镜像信息，类似 docker images
func (m *Manager) ListImage(ctx context.Context) ([]Image, error) {
	images, err := m.Client.ImageList(ctx, image.ListOptions{All: true})
	if err != nil {
		return nil, err
	}

	var imageList []Image
	for _, image := range images {
		if len(image.RepoTags) == 0 {
			continue
		}
		for _, repoTag := range image.RepoTags {
			tags := strings.Split(repoTag, ":")
			im := Image{
				ID:      image.ID,
				Name:    tags[0],
				Tag:     tags[1],
				Created: time.Unix(image.Created, 0).Format("2006-01-02 15:04:05"),
				Size:    strconv.FormatFloat(float64(image.Size)/(1000*1000), 'f', 2, 64) + "MB",
			}
			imageList = append(imageList, im)
		}
	}
	return imageList, nil
}

// GetImageByName 根据 imageName 获取 Image 详情, imageName -> image:latest
func (m *Manager) GetImageByName(ctx context.Context, imageName string) (*Image, error) {
	images, err := m.Client.ImageList(ctx, image.ListOptions{All: true})
	if err != nil {
		return nil, err
	}

	for _, v := range images {
		for _, t := range v.RepoTags {
			if t == imageName {
				tagsList := strings.Split(t, ":")
				return &Image{
					ID:      v.ID,
					Name:    tagsList[0],
					Tag:     tagsList[1],
					Created: time.Unix(v.Created, 0).Format("2006-01-02 15:04:05"),
					Size:    strconv.FormatFloat(float64(v.Size)/(1000*1000), 'f', 2, 64) + "MB",
				}, nil
			}
		}
	}
	return nil, errors.New("not found image")
}

// GetImageByID 根据 imageID 获取 Image 详情
func (m *Manager) GetImageByID(ctx context.Context, imageID string) (*Image, error) {
	imageResponse, _, err := m.Client.ImageInspectWithRaw(ctx, imageID)
	if err != nil {
		return nil, err
	}

	tagsList := strings.Split(imageResponse.RepoTags[0], ":")

	return &Image{
		ID:      imageResponse.ID,
		Name:    tagsList[0],
		Tag:     tagsList[1],
		Created: imageResponse.Created,
		Size:    strconv.FormatFloat(float64(imageResponse.Size)/(1000*1000), 'f', 2, 64) + "MB",
	}, nil
}

func (m *Manager) RemoveImage(ctx context.Context, imageID string) error {
	_, err := m.Client.ImageRemove(ctx, imageID, image.RemoveOptions{Force: true})
	return err
}

func (m *Manager) PruneImages(ctx context.Context) error {
	_, err := m.Client.ImagesPrune(ctx, filters.NewArgs(filters.Arg("dangling", "true")))
	return err
}

// SearchImage 通过关键词查找镜像
func (m *Manager) SearchImage(ctx context.Context, term string) ([]registry.SearchResult, error) {
	search, err := m.Client.ImageSearch(ctx, term, types.ImageSearchOptions{Limit: 10})
	if err != nil {
		return []registry.SearchResult{}, err
	}
	fmt.Println(search)
	return search, nil
}

// https://blog.csdn.net/u010918487/article/details/105788735

// PullImage 根据名称拉去镜像，term 可以是 镜像名称(ubuntu 会拉去 ubuntu:latest) 也可以是 镜像名称:tag(ubuntu:18.04)
func (m *Manager) PullImage(ctx context.Context, term string) error {
	pullReader, err := m.Client.ImagePull(ctx, term, image.PullOptions{All: false, PrivilegeFunc: nil, RegistryAuth: ""})
	if err != nil {
		return err
	}
	defer func(pullReader io.ReadCloser) {
		err := pullReader.Close()
		if err != nil {
			return
		}
	}(pullReader)
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(pullReader)
	if err != nil {
		return err
	}
	return nil
}

// TagImage 修改镜像 tag oldTag: ubuntu:latest  newTag ubuntu:22.04
func (m *Manager) TagImage(ctx context.Context, oldTag string, newTag string) error {
	return m.Client.ImageTag(ctx, oldTag, newTag)
}

// https://www.cnblogs.com/guangdelw/p/17567195.html

// ImportImage 镜像导入
func (m *Manager) ImportImage(ctx context.Context, sourceFile string) error {
	inputFile, err := os.Open(sourceFile)
	if err != nil {
		return err
	}
	defer func(inputFile *os.File) {
		err := inputFile.Close()
		if err != nil {
			return
		}
	}(inputFile)

	resp, err := m.Client.ImageLoad(ctx, inputFile, true)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	// 读取并输出导入过程
	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

// ExportImage 镜像导出
func (m *Manager) ExportImage(ctx context.Context, imageIDs []string, targetFile string) error {
	resp, err := m.Client.ImageSave(ctx, imageIDs)
	if err != nil {
		return err
	}
	defer func(resp io.ReadCloser) {
		err := resp.Close()
		if err != nil {
			return
		}
	}(resp)
	outputFile, err := os.Create(targetFile)
	if err != nil {
		return err
	}
	defer func(outputFile *os.File) {
		err := outputFile.Close()
		if err != nil {
			return
		}
	}(outputFile)

	_, err = io.Copy(outputFile, resp)
	if err != nil {
		return err
	}
	return nil
}
