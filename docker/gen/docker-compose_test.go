// Package gen
// Date: 2023/4/6 16:48
// Author: Amu
// Description:
package gen

import (
	"fmt"
	"os"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestDockerComposeConfig(t *testing.T) {
	baseConfig := DockerComposeConfig()
	out, err := yaml.Marshal(baseConfig)
	if err == nil {
		fmt.Printf("docker-compose: %v\n", string(out))
		err := os.WriteFile("./docker-compose.yaml", out, 0644)
		if err != nil {
			fmt.Printf("write error: %v\n", err)
		}
	}
}
