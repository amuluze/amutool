# Amutool

![Go version](https://img.shields.io/badge/go-%3E%3Dv1.18-9cf)

golang 工具类封装

[//]: # ([![Release]&#40;https://img.shields.io/badge/release-2.2.3-green.svg&#41;]&#40;https://github.com/duke-git/lancet/releases&#41;)
[//]: # ([![GoDoc]&#40;https://godoc.org/github.com/duke-git/lancet/v2?status.svg&#41;]&#40;https://pkg.go.dev/github.com/duke-git/lancet/v2&#41;)

[//]: # ([![Go Report Card]&#40;https://goreportcard.com/badge/github.com/duke-git/lancet/v2&#41;]&#40;https://goreportcard.com/report/github.com/duke-git/lancet/v2&#41;)

[//]: # ([![test]&#40;https://github.com/duke-git/lancet/actions/workflows/codecov.yml/badge.svg?branch=main&event=push&#41;]&#40;https://github.com/duke-git/lancet/actions/workflows/codecov.yml&#41;)

[//]: # ([![codecov]&#40;https://codecov.io/gh/duke-git/lancet/branch/main/graph/badge.svg?token=FC48T1F078&#41;]&#40;https://codecov.io/gh/duke-git/lancet&#41;)

[//]: # ([![License]&#40;https://img.shields.io/badge/license-MIT-blue.svg&#41;]&#40;https://github.com/duke-git/lancet/blob/main/LICENSE&#41;)


## 文档

- [Bannerx](#bannerx)
- [Randx](#randx)

### bannerx
bannerx 根据输入字符串生成一张 banner。 [[doc](https://gitee.com/amuluze/amutool/main/docs/bannerx.md)]
**函数列表：**
- GenerateBanner: 根据输入字符串生成 banner。

### randx
randx 随机数生成包，可以随机生成随机 []byte,int,string。 [[doc](https://gitee.com/amuluze/amutool/main/docs/randx.md)]
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
- [x] bannerx: banner 生成
- [x] basex: base 64 编解码
- [x] randx: 随机整数、字符串生成
- [x] command: 运行 linux 命令
- [x] database: 数据库操作，支持 postgres mysql sqlite
- [x] logx: 日志 zap 封装
- [x] conf: 配置文件加载，支持 json yaml toml
- [x] envx: 环境变量获取
- [x] errors: error 封装

## TODO
- [ ] clickhouse: clickhouse 客户端
- [ ] convertx: 类型转换
- [ ] docker: golang 操作 docker
- [ ] es: Elasticsearch Client and BulkClient
- [ ] hashx:摘要算法
- [ ] httpx:简单的 http 客户端
- [ ] iohelper: io 操作相关
- [ ] kafka:操作 kafka
- [ ] redis: redis 客户端
- [ ] stringx: 字符串操作封装
- [ ] timex: 时间相关操作
- [ ] uuidx: uuid 相关操作
- [ ] jsonrpc
- [ ] doc: api 文档自动生成
- [ ] encrypt: 加解密
- [ ] executors
- [ ] requests
- [ ] rescue
- [ ] rpc
- [ ] task
- [ ] ws

