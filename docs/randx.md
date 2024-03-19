# Randx

randx 随机数生成器包，可以生成随机[]bytes, int, string。

## 源码:


## 用法:

```go
import (
    "github.com/amuluze/amutool/randx"
)
```

## 目录

-   [RandBytes](#RandBytes)
-   [RandInt](#RandInt)
-   [RandString](#RandString)
-   [RandUpper](#RandUpper)
-   [RandLower](#RandLower)
-   [RandNumeral](#RandNumeral)
-   [RandNumeralOrLetter](#RandNumeralOrLetter)
-   [UUIdV4](#UUIdV4)
-   [RandUniqueIntSlice](#RandUniqueIntSlice)


## 文档

### RandBytes
生成随机字节切片
函数签名:
```go
func RandBytes(length int) []byte
```
示例:
```go
package main

import (
    "fmt"
	"github.com/amuluze/amutool/randx"
)

func main() {
    randBytes := randx.RandBytes(4)
    fmt.Println(randBytes)
}
```

### RandInt

生成随机int, 范围[min, max)
函数签名:
```go
func RandInt(min, max int) int
```

示例:

```go
package main

import (
    "fmt"
	"github.com/amuluze/amutool/randx"
)

func main() {
    rInt := randx.RandInt(1, 10)
    fmt.Println(rInt)
}
```

### RandString

生成给定长度的随机字符串，只包含字母(a-zA-Z)

函数签名:

```go
func RandString(length int) string
```

示例:

```go
package main

import (
    "fmt"
	"github.com/amuluze/amutool/randx"
)

func main() {
    randStr := randx.RandString(6)
    fmt.Println(randStr) //pGWsze
}
```

### RandUpper

生成给定长度的随机大写字母字符串

函数签名:

```go
func RandUpper(length int) string
```

示例:

```go
package main

import (
    "fmt"
	"github.com/amuluze/amutool/randx"
)

func main() {
    randStr := randx.RandString(6)
    fmt.Println(randStr) //PACWGF
}
```

### RandLower

生成给定长度的随机小写字母字符串

函数签名:

```go
func RandLower(length int) string
```

示例:

```go
package main

import (
    "fmt"
	"github.com/amuluze/amutool/randx"
)

func main() {
    randStr := randx.RandLower(6)
    fmt.Println(randStr) //siqbew
}
```

### RandNumeral

生成给定长度的随机数字字符串

函数签名:

```go
func RandNumeral(length int) string
```

示例:

```go
package main

import (
    "fmt"
	"github.com/amuluze/amutool/randx"
)

func main() {
    randStr := randx.RandNumeral(6)
    fmt.Println(randStr) //035172
}
```

### RandNumeralOrLetter

生成给定长度的随机字符串（数字+字母)

函数签名:

```go
func RandNumeralOrLetter(length int) string
```

示例:

```go
package main

import (
    "fmt"
	"github.com/amuluze/amutool/randx"
)

func main() {
    randStr := randx.RandNumeralOrLetter(6)
    fmt.Println(randStr) //0aW7cQ
}
```

### UUIdV4

生成UUID v4字符串

函数签名:

```go
func UUIdV4() (string, error)
```

示例:

```go
package main

import (
    "fmt"
	"github.com/amuluze/amutool/randx"
)

func main() {
    uuid, err := randx.UUID4()
    if err != nil {
        return
    }
    fmt.Println(uuid)
}
```
