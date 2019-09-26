# Direwolf HTTP Client: Save your time

Package direwolf is a convenient and easy to use http client written in Golang.

If you want find more info, please go here: [Direwolf HTTP Client: Save your time](https://wnanbei.github.io/direwolf/)

![direwolf](docs/assets/cover.jpg)

## Feature Support

- Clean and Convenient API
- Simple to Set Headers, Cookies, Parameters, Post Forms
- Sessions with Cookie Persistence
- Keep-Alive & Connection Pooling
- HTTP(S) Proxy Support
- Redirect Control
- Timeout Control
- Support extract result from response body with css selector, regexp, json
- Content Decoding
- More to come...

## Installation

```text
go get github.com/wnanbei/direwolf
```

## Take a glance

You can easily send a request like this:

```go
import (
    "fmt"

    dw "github.com/wnanbei/direwolf"
)

func main() {
    resp, err := dw.Get("https://www.google.com")
    if err != nil {
        return
    }
    fmt.Println(resp.Text())
}
```

Besides, direwolf provide a convenient way to add parameters to request. Such
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
        return
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
    "origin": "1.1.1.1, 1.1.1.1",
    "url": "https://httpbin.org/get?age=18&name=wnanbei"
}
```
