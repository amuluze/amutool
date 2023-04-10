// Package gen
// Date: 2023/4/6 15:45
// Author: Amu
// Description:
package gen

import "github.com/mcuadros/go-defaults"

type Config struct {
	Subnet string `yaml:"subnet" default:"172.15.0.0/16"`
}

type DockerCompose struct {
	Version  string `yaml:"version" default:"2.4"`
	Services struct {
		Db struct {
			Image         string `yaml:"image" default:"postgres:12"`
			ContainerName string `yaml:"container_name" default:"db"`
			Cpuset        string `yaml:"cpuset" default:"1-2"`
			Logging       struct {
				Options struct {
					MaxSize string `yaml:"max-size" default:"10m"`
					MaxFile string `yaml:"max-file" default:"10"`
				} `yaml:"options"`
			} `yaml:"logging"`
			Environment struct {
				POSTGRESUSER     string `yaml:"POSTGRES_USER" default:"hera"`
				POSTGRESPASSWORD string `yaml:"POSTGRES_PASSWORD" default:"hera"`
				POSTGRESDB       string `yaml:"POSTGRES_DB" default:"hera"`
			} `yaml:"environment"`
			Command  []string `yaml:"command" default:"[postgres,-c,config_file=/etc/postgresql/postgresql.conf]"`
			Volumes  []string `yaml:"volumes" default:"[./workdir/pg_data:/var/lib/postgresql/data,./workdir/pg_host:/host,./config/hera/prod_init.sql:/host_init/prod_init.sql,/etc/localtime:/etc/localtime:ro,./config/postgresql/postgresql.conf:/etc/postgresql/postgresql.conf]"`
			MemLimit string   `yaml:"mem_limit" default:"0"`
			Restart  string   `yaml:"restart" default:"always"`
			Networks struct {
				AppNet struct {
					Ipv4Address string `yaml:"ipv4_address" default:"172.15.1.10"`
				} `yaml:"app_net"`
			} `yaml:"networks"`
			Ports []string `yaml:"ports" default:"[127.0.0.1:5432:5432]"`
		} `yaml:"db"`
		Redis struct {
			Image         string `yaml:"image" default:"redis:7.0"`
			ContainerName string `yaml:"container_name" default:"redis"`
			Restart       string `yaml:"restart" default:"always"`
			Ulimits       struct {
				Nproc  int `yaml:"nproc" default:"1048576"`
				Nofile int `yaml:"nofile" default:"1048576"`
			} `yaml:"ulimits"`
			Logging struct {
				Options struct {
					MaxSize string `yaml:"max-size" default:"10M"`
					MaxFile string `yaml:"max-file" default:"10"`
				} `yaml:"options"`
			} `yaml:"logging"`
			Command  []string `yaml:"command" default:"[redis-server,--requirepass,e585a8e68289e7899b624032303231]"`
			Volumes  []string `yaml:"volumes" default:"[./workdir/redis_data:/data,/etc/localtime:/etc/localtime:ro]"`
			MemLimit string   `yaml:"mem_limit" default:"0"`
			Networks struct {
				AppNet struct {
					Ipv4Address string `yaml:"ipv4_address" default:"172.15.0.61"`
				} `yaml:"app_net"`
			} `yaml:"networks"`
		} `yaml:"redis"`
	} `yaml:"services"`
	Networks struct {
		AppNet struct {
			Name string `yaml:"name" default:"app_net"`
			Ipam struct {
				Config []Config `yaml:"config"`
			} `yaml:"ipam"`
			//External bool `yaml:"external" default:"true"`
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
