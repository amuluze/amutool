// Package resources
// Date: 2024/6/26 18:13
// Author: Amu
// Description:
package resources

import "path/filepath"

var (
	RootPath = "resources"

	PostgresFolder     = "postgres"
	PostgresDataFolder = filepath.Join(PostgresFolder, "data")

	RedisFolder = "redis"

	ESFolder     = "es"
	ESDataFolder = filepath.Join(ESFolder, "data")

	NginxFolder      = "nginx"
	NginxConfigPath  = filepath.Join(NginxFolder, "nginx.conf")
	NginxLogsFolder  = filepath.Join(NginxFolder, "logs")
	NginxCertsFolder = filepath.Join(NginxFolder, "certs")
)
