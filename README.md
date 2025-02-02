# Amutool

![Go version](https://img.shields.io/badge/go-%3E%3Dv1.18-9cf)

golang 工具类封装

## 文档

- [Bannerx](#bannerx)
- [Randx](#randx)

### bannerx

bannerx 根据输入字符串生成一张 banner。 [[doc](https://github.com/amuluze/amutool/blob/master/docs/bannerx.md)]
**函数列表：**

- GenerateBanner: 根据输入字符串生成 banner。

### randx

randx 随机数生成包，可以随机生成随机 []byte,int,string。 [[doc](https://github.com/amuluze/amutool/blob/master/docs/randx.md)]
**函数列表：**

- RandBytes: 生成随机字节切片。
- RandInt: 生成随机 int, 范围[min, max)。
- RandString: 生成给定长度的随机字符串，只包含字母(a-zA-Z)。
- RandUpper: 生成给定长度的随机大写字母字符串(A-Z)。
- RandLower: 生成给定长度的随机小写字母字符串(a-z)。
- RandNumeral: 生成给定长度的随机数字字符串(0-9)。
- RandNumeralOrLetter: 生成给定长度的随机字符串（数字+字母)。
- UUID4: 生成 UUID v4 字符串。

## DONE

- [X] bannerx: banner 生成
- [X] basex: base 64 编解码
- [X] randx: 随机整数、字符串生成
- [X] command: 运行 linux 命令
- [X] database: 数据库操作，支持 postgres mysql sqlite
- [X] logx: 日志 zap 封装
- [X] conf: 配置文件加载，支持 json yaml toml
- [X] envx: 环境变量获取
- [X] errors: error 封装
- [X] hashx:摘要算法
- [X] uuidx: uuid 相关操作
- [X] kafka:操作 kafka
- [X] es: Elasticsearch Client and BulkClient
- [X] timex: 时间相关操作
- [X] redis: redis 客户端
- [X] iohelper: io 操作相关
- [X] docker: golang 操作 docker
- [X] gpool: 协程池，用于并发除了简单任务
- [X] clickhousex: clickhouse 客户端，包含 BatchProcessor，支持批量写入

## TODO

- [ ] convertx: 类型转换
- [ ] httpx:简单的 http 客户端
- [ ] stringx: 字符串操作封装
- [ ] jsonrpc
- [ ] doc: api 文档自动生成
- [ ] encrypt: 加解密
- [ ] executors
- [ ] requests
- [ ] rescue
- [ ] rpc
- [ ] task
- [ ] ws
