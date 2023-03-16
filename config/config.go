// Package config
// Date: 2022/9/30 11:25
// Author: Amu
// Description:
package config

import (
	"fmt"
	"strings"
	"sync"

	"github.com/fsnotify/fsnotify"

	"github.com/koding/multiconfig"
	"github.com/spf13/viper"
)

type Config struct {
	//Database Database
	Servers Servers
}

type Servers struct {
	Prod Server
	Dev  Server
}

type Server struct {
	Host            string
	Port            string
	ShutDownTimeout int64
}

var Cfg = new(Config)
var once sync.Once

func loadConfigs() {
	//configs := viper.New()
	//configs.AddConfigPath(configFilePath)
	//
	//configFileName := configFilePrefix
	//configs.SetConfigName(configFileName)
	//configs.SetConfigType(configFileType)
	viper.SetConfigFile("./config.toml")
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件变化： ", in.Name)
		err := viper.Unmarshal(&Cfg)
		if err != nil {
			fmt.Println("更新配置错误：", err)
		}
	})

	// 尝试进行读取
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&Cfg); err != nil {
		fmt.Println(err)
	}
}

func GetConfigs() *Config {
	if Cfg != nil {
		fmt.Println("Config is not nil")
		return Cfg
	}
	fmt.Println("Config is nil")
	loadConfigs()
	return Cfg
}

func MustLoad(model interface{}, filePaths ...string) {
	once.Do(func() {
		loaders := []multiconfig.Loader{
			&multiconfig.TagLoader{},
			&multiconfig.EnvironmentLoader{},
		}

		for _, filePath := range filePaths {
			fmt.Println(filePath)
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
