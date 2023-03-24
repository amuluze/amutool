// Package conf
// Date: 2022/9/30 11:25
// Author: Amu
// Description:
package conf

import (
	"strings"
	"sync"

	"github.com/koding/multiconfig"
)

var once sync.Once

func MustLoad(model interface{}, filePaths ...string) {
	once.Do(func() {
		loaders := []multiconfig.Loader{
			&multiconfig.TagLoader{},
			&multiconfig.EnvironmentLoader{},
		}

		for _, filePath := range filePaths {
			if strings.HasSuffix(filePath, "toml") {
				loaders = append(loaders, &multiconfig.TOMLLoader{Path: filePath})
			}
			if strings.HasSuffix(filePath, "json") {
				loaders = append(loaders, &multiconfig.JSONLoader{Path: filePath})
			}
			if strings.HasSuffix(filePath, "yaml") {
				loaders = append(loaders, &multiconfig.YAMLLoader{Path: filePath})
			}
		}
		m := multiconfig.DefaultLoader{
			Loader:    multiconfig.MultiLoader(loaders...),
			Validator: multiconfig.MultiValidator(&multiconfig.RequiredValidator{}),
		}
		m.MustLoad(model)
	})
}
