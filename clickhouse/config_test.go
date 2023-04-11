// Package clickhouse
// Date: 2023/4/11 11:43
// Author: Amu
// Description:
package clickhouse

import (
	"fmt"
	"testing"

	"gitee.com/amuluze/amutool/conf"
)

func TestConfig(t *testing.T) {
	var config Config
	conf.MustLoad(&config, "./config.toml")
	fmt.Println(config)
}
