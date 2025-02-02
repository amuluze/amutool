// Package docker
// Date: 2024/03/21 16:36:54
// Author: Amu
// Description:
package docker

import (
	"context"
	"testing"
)

func TestVersion(t *testing.T) {
	manager, _ := NewManager()
	t.Log(manager.Version(context.Background()))
}
