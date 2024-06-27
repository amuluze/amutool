// Package compose
// Date: 2023/4/6 15:45
// Author: Amu
// Description:
package compose

import (
	"fmt"
	"os"

	"github.com/mcuadros/go-defaults"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Subnet string `yaml:"subnet" load:"172.15.0.0/16"`
}

type DockerCompose struct {
	Version  string `yaml:"version" load:"2.4"`
	Services struct {
		Db struct {
			Image         string `yaml:"image" load:"postgres:12"`
			ContainerName string `yaml:"container_name" load:"db"`
			Cpuset        string `yaml:"cpuset" load:"1-2"`
			Logging       struct {
				Options struct {
					MaxSize string `yaml:"max-size" load:"10m"`
					MaxFile string `yaml:"max-file" load:"10"`
				} `yaml:"options"`
			} `yaml:"logging"`
			Environment struct {
				POSTGRESUSER     string `yaml:"POSTGRES_USER" load:"hera"`
				POSTGRESPASSWORD string `yaml:"POSTGRES_PASSWORD" load:"hera"`
				POSTGRESDB       string `yaml:"POSTGRES_DB" load:"hera"`
			} `yaml:"environment"`
			Command  []string `yaml:"command" load:"[postgres,-c,config_file=/etc/postgresql/postgresql.conf]"`
			Volumes  []string `yaml:"volumes" load:"[./workdir/pg_data:/var/lib/postgresql/data,./workdir/pg_host:/host,./config/hera/prod_init.sql:/host_init/prod_init.sql,/etc/localtime:/etc/localtime:ro,./config/postgresql/postgresql.conf:/etc/postgresql/postgresql.conf]"`
			MemLimit string   `yaml:"mem_limit" load:"0"`
			Restart  string   `yaml:"restart" load:"always"`
			Networks struct {
				AppNet struct {
					Ipv4Address string `yaml:"ipv4_address" load:"172.15.1.10"`
				} `yaml:"app_net"`
			} `yaml:"networks"`
			Ports []string `yaml:"ports" load:"[127.0.0.1:5432:5432]"`
		} `yaml:"db"`
		Redis struct {
			Image         string `yaml:"image" load:"redis:7.0"`
			ContainerName string `yaml:"container_name" load:"redis"`
			Restart       string `yaml:"restart" load:"always"`
			Ulimits       struct {
				Nproc  int `yaml:"nproc" load:"1048576"`
				Nofile int `yaml:"nofile" load:"1048576"`
			} `yaml:"ulimits"`
			Logging struct {
				Options struct {
					MaxSize string `yaml:"max-size" load:"10M"`
					MaxFile string `yaml:"max-file" load:"10"`
				} `yaml:"options"`
			} `yaml:"logging"`
			Command  []string `yaml:"command" load:"[redis-server,--requirepass,e585a8e68289e7899b624032303231]"`
			Volumes  []string `yaml:"volumes" load:"[./workdir/redis_data:/data,/etc/localtime:/etc/localtime:ro]"`
			MemLimit string   `yaml:"mem_limit" load:"0"`
			Networks struct {
				AppNet struct {
					Ipv4Address string `yaml:"ipv4_address" load:"172.15.0.61"`
				} `yaml:"app_net"`
			} `yaml:"networks"`
		} `yaml:"redis"`
	} `yaml:"services"`
	Networks struct {
		AppNet struct {
			Name string `yaml:"name" load:"app_net"`
			Ipam struct {
				Config []Config `yaml:"config"`
			} `yaml:"ipam"`
			//External bool `yaml:"external" load:"true"`
		} `yaml:"app_net"`
	} `yaml:"networks"`
}

func DockerComposeConfig() DockerCompose {
	var dockerCompose DockerCompose
	defaults.SetDefaults(&dockerCompose)

	var subnetConfig Config
	defaults.SetDefaults(&subnetConfig)
	config := []Config{
		subnetConfig,
	}

	dockerCompose.Networks.AppNet.Ipam.Config = config

	return dockerCompose
}

func GenerateDockerComposeFile(dockerComposeFilePath string) {
	baseConfig := DockerComposeConfig()
	out, err := yaml.Marshal(baseConfig)
	if err == nil {
		err := os.WriteFile(dockerComposeFilePath, out, 0644)
		if err != nil {
			fmt.Printf("write error: %v\n", err)
		}
	}
}
