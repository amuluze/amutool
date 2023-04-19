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

	"github.com/docker/docker/api/types"
)

type Image struct {
	ID      string
	Name    string
	Tag     string
	Created string
	Size    string
}

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
