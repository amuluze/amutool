# requests

## install

```golang
go get github.com/amuluze/amutool/requests
```

## usage

### server api

```python
import uvicorn
from fastapi import FastAPI, Request
from pydantic import BaseModel


app = FastAPI()

@app.get("/aaa")
def home():
    return {"message": "Hello World"}


class User(BaseModel):
    name: str
    age: int

@app.post("/users")
def greet(request: Request, item: User):
    print(item)
    headers = request.headers
    cookies = request.cookies
    print(headers)
    print(cookies)
    return {"message": "success", "code": 10000, "data": [{"name": "jack", "age": 12}, {"name": "nick", "age": 15}]}

if __name__ == "__main__":
    uvicorn.run(app=app, host="0.0.0.0", port=8090)
```

> 运行服务: `python3 main.py`

### client

```golang
package main

import (
 "fmt"
 "github.com/amuluze/amutool/requests"
)

type MessageReponse struct {
 Message string `json:"message"`
}

func main() {
 reply := &MessageReponse{}
 err := requests.Get("http://1xxx.xxx.xxx.xxx5:8090/aaa", nil, reply)
    if err!= nil {
        fmt.Println(err)
    }
    fmt.Println(reply)
}

```

```golang
package main

import (
 "fmt"
 "github.com/amuluze/amutool/requests"
)

type User struct {
 Name string `json:"name"`
 Age  int    `json:"age"`
}

type PostResponse struct {
 Code    int    `json:"code"`
 Message string `json:"message"`
 Data    []User `json:"data"`
}

func main() {
 params := &User{
  Name: "Amu",
  Age:  18,
 }
 reply := &PostResponse{}
 err := Post("http://1xxx.xxx.xxx.xxx:8090/users", params, reply, SetHeader("Referer", "https://example.com/test"), SetCookie("cid", "123456"))
 if err != nil {
  fmt.Println(err)
 }
 fmt.Println(reply)
 }

```
