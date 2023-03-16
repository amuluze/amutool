// Package config
// Date: 2022/9/30 11:53
// Author: Amu
// Description:
package config

import (
	"fmt"
	"testing"
)

func TestGetConfigs(t *testing.T) {
	res := GetConfigs()
	fmt.Printf("%#v\n", res.Servers.Prod)
	fmt.Printf("%#v\n", res.Servers.Dev)
}

func TestMustLoad(t *testing.T) {
	MustLoad(Cfg, "./config.toml")
	fmt.Printf("prod: %#v\n", Cfg)
}
