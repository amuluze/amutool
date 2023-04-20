## amutool

### Done
- [x] bannerx: banner 生成
- [x] basex: base 64 编解码
- [x] clickhouse: clickhouse 客户端
- [x] command: 运行 linux 命令
- [x] conf: 配置文件加载，支持 json yaml toml
- [x] convertx: 类型转换
  - [x] AsString: any to string
  - [x] ToInt64: any to int64
  - [x] BytesToInt64: []byte to int64
  - [x] Int64ToBytes: int64 to []byte
  - [x] StringToInt: string to int
  - [x] IntToInt64: int to int64
- [ ] docker: golang 操作 docker
  - [ ] container: 容器相关操作
  - [x] image: 镜像相关操作
  - [x] network: 网络相关操作
- [x] database: 数据库操作，支持 postgres mysql sqlite
- [x] envx: 环境变量获取
- [x] errors: error 封装
- [ ] hashx:摘要算法
  - [x] MD5
  - [x] SHA1
  - [ ] SHA256
- [x] httpx:简单的 http 客户端
  - [x] Get
  - [x] PostParams
  - [x] PostJson
- [x] iohelper: io 操作相关
  - [x] file: 文件相关
  - [x] path:路径相关
  - [x] md5_change: 每次运行时检查，指定文件是否变动
- [x] kafka:操作 kafka
  - [x] Consumer: 消费者
  - [x] ConsumerGroup: 消费者组
  - [x] Producer: 生产者
- [x] logx: 日志 zap 封装
- [x] randx: 随机整数、字符串生成
- [x] redis: redis 客户端
- [x] stringx: 字符串操作封装
- [x] timex: 时间相关操作
  - [ ] time.Time string timestamp 相关转换
  - [x] time.Ticker 封装
- [x] uuidx: uuid 相关操作

### Todo
- [ ] es: Elasticsearch Client and BulkClient
- [ ] doc: api 文档自动生成
- [ ] encrypt: 加解密
- [ ] executors
- [ ] requests
- [ ] rescue
- [ ] rpc
- [ ] task
- [ ] ws

### Examples
- [ ] examples