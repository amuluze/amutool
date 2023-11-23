### db

支持的数据库：
- postgresql
- mysql
- sqlite

toml 配置示例：
```toml
[DB]
# 连接地址
Host = "localhost"
# 连接端口
Port = 5432
# 用户名
User = "root"
# 密码
Password = "123456"
# 数据库
DBName = "test"
# SSL模式
SSLMode = "disable"
# 是否开启调试模式
Debug = true
# 数据库类型(目前支持的数据库类型：postgres, mysql, sqlite)
DBType = "postgres"
# 设置连接可以重用的最长时间(单位：秒)
MaxLifetime = 7200
# 设置数据库的最大打开连接数
MaxOpenConns = 150
# 设置空闲连接池中的最大连接数
MaxIdleConns = 50
# 数据库表名前缀
TablePrefix = "s_"
# 是否启用自动映射数据库表结构
EnableAutoMigrate = true

```