// Package kafka
// Date: 2023/3/31 11:20
// Author: Amu
// Description:
package kafka

import (
	"fmt"
	"testing"

	"gitee.com/amuluze/amutool/conf"
)

func TestConfig(t *testing.T) {
	var cfg = new(Config)
	conf.MustLoad(cfg, "./config.toml")
	fmt.Printf("cfg: %#v\n", cfg.Kafka)
}
