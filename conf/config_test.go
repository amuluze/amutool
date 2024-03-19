// Package conf
// Date: 2022/9/30 11:53
// Author: Amu
// Description:
package conf

import (
	"fmt"
	"testing"
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

func TestMustLoad(t *testing.T) {
	//MustLoad(Cfg, "./config.toml")
	MustLoad(Cfg, "./config.toml")
	fmt.Printf("dev: %#v\n", Cfg.Servers.Dev)
	fmt.Printf("prod: %#v\n", Cfg.Servers.Prod)
}
