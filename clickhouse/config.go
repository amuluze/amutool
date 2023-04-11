// Package clickhouse
// Date: 2023/4/7 18:06
// Author: Amu
// Description:
package clickhouse

type Config struct {
	Addr            string // 服务地址
	Username        string // 用户名
	Password        string // 密码，没有则为空
	Database        string // 使用数据库
	DialTimeout     string // 读取超时，默认 3s，-1 表示取消读超时
	MaxOpenConns    int    // 最大连接数
	MaxIdleConns    int    // 最大空闲连接数
	ConnMaxLifeTime string // 连接存活时间
}
