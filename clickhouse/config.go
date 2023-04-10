// Package clickhouse
// Date: 2023/4/7 18:06
// Author: Amu
// Description:
package clickhouse

type Config struct {
	ClickHouse ClickHouse
}

type ClickHouse struct {
	Addrs           []string // redis 地址，兼容单机和集群
	Username        string   // 用户名
	Password        string   // 密码，没有则为空
	Database        string   // 使用数据库
	DialTimeout     string   // 读取超时，默认 3s，-1 表示取消读超时
	MaxOpenConn     int      // 最大连接数
	MaxIdleConns    int
	ConnMaxLifeTime string
}
