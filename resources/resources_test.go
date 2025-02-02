// Package resources
// Date: 2024/6/26 18:19
// Author: Amu
// Description:
package resources

import "testing"

func TestResources(t *testing.T) {
	t.Logf("resources: %s", RootPath)
	t.Logf("Nginx config: %s", NginxConfigPath)
}
