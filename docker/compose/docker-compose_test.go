// Package compose
// Date: 2023/4/6 16:48
// Author: Amu
// Description:
package compose

import (
	"testing"
)

func TestDockerComposeConfig(t *testing.T) {
	dockerComposeFilePath := "./docker-compose.yaml"
	GenerateDockerComposeFile(dockerComposeFilePath)
}
