// Package docker
// Date: 2023/4/19 15:24
// Author: Amu
// Description:
package docker

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/image"
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
	images, err := m.Client.ImageList(ctx, types.ImageListOptions{All: true})
	if err != nil {
		return nil, err
	}

	var imageList []Image
	for _, image := range images {
		if len(image.RepoTags) == 0 {
			continue
		}
		tagsList := strings.Split(image.RepoTags[0], ":")
		im := Image{
			ID:      image.ID,
			Name:    tagsList[0],
			Tag:     tagsList[1],
			Created: time.Unix(image.Created, 0).Format("2006-01-02 15:04:05"),
			Size:    strconv.FormatFloat(float64(image.Size)/(1000*1000), 'f', 2, 64) + "MB",
		}
		imageList = append(imageList, im)
	}
	return imageList, nil
}

// GetImageByName 根据 imageName 获取 Image 详情, imageName -> image:latest
func (m *Manager) GetImageByName(ctx context.Context, imageName string) (*Image, error) {
	images, err := m.Client.ImageList(ctx, types.ImageListOptions{All: true})
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
