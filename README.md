# Direwolf HTTP Client: Save your time

Package direwolf is a convient and esay to use http client written in Golang. 

## Feature Support

- Clean and Convient API
- Simple to Set Headers, Cookies, Parameters, Post Forms 
- Sessions with Cookie Persistence
- Elegant Key/Value Cookies
- Keep-Alive & Connection Pooling
- HTTP(S) Proxy Support
- Redirect Control
- Timeout Control
- Support extract result from response body with css selector, regexp, json
- Content Decoding
- More to come

## Installation

```
go get github.com/wnanbei/direwolf
```

## Quick Start

You can easily send a request like this:

```go
import (
    "fmt"

    dw "github.com/wnanbei/direwolf"
)

func main() {
    resp, err := dw.Get("https://www.google.com")
    if err != nil {
        ...
    }
    fmt.Println(resp.Text())
}
```

Besides, direwolf provide a convient way to add parameters to request. Such
as Headers, Cookies, Params, etc.

```go
import (
    "fmt"

    dw "github.com/wnanbei/direwolf"
)

func main() {
    headers := dw.NewHeaders(
        "User-Agent", "direwolf",
    )
    params := dw.NewParams(
        "name", "wnanbei",
        "age", "18",
    )
    cookies := dw.NewCookies(
        "sign", "kzhxciuvyqwekhiuxcyvnkjdhiue",
    )
    resp, err := dw.Get("https://httpbin.org/get", headers, params, cookies)
    if err != nil {
        ...
    }
    fmt.Println(resp.Text())
}
```

Output:

```json
{
    "args": {
        "age": "18",
        "name": "wnanbei"
    },
    "headers": {
        "Accept-Encoding": "gzip",
        "Cookie": "sign=kzhxciuvyqwekhiuxcyvnkjdhiue",
        "Host": "httpbin.org",
        "User-Agent": "direwolf"
    },
    "origin": "118.116.15.151, 118.116.15.151",
    "url": "https://httpbin.org/get?age=18&name=wnanbei"
}
```

## API examples

### 1. Make Request

First, you can start a request like this:

```go
import (
    dw "github.com/wnanbei/direwolf"
)

func main() {
    resp, err := dw.Get("https://httpbin.org/get")
}
```

Other HTTP request types: 

```go
resp, err := dw.Post("https://httpbin.org/post", dw.NewPostForm("key", "value"))
resp, err := dw.Head("https://httpbin.org/head")
resp, err := dw.Put("https://httpbin.org/put", dw.NewPostForm("key", "value"))
resp, err := dw.Delete("https://httpbin.org/delete")
```

### 2. Passing Parameters In URLs

Passing parameters in URLs is very easy, you only need to new a Params and pass it to request method.

```go
import (
    "fmt"

    dw "github.com/wnanbei/direwolf"
)

func main() {
    params := dw.NewParams("key", "value")
    resp, err := dw.Get("https://httpbin.org/get", params)
    if err != nil {
        ...
    }
    fmt.Println(resp.URL)
}
```

Output:

```
https://httpbin.org/get?key=value
```

If you want pass more parameters to URLs, just like this:

```go
params := dw.NewParams(
    "key1", "value1",
    "key2", "value2",
)
```

**Note: Remember the comma between key and value.** 

**Note: Key must to match Value one by one, if not, will report an error.**

If the parameters have the same key, it`s ok:

```go
params := dw.NewParams(
    "key1", "value1",
    "key1", "value2",
)
```

Output:

```
https://httpbin.org/get?key1=value1&key1=value2
```

