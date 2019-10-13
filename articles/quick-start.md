---
layout: single
title: Quick Start

toc: true
toc_label: "Quick Start"

sidebar:
  nav: "sidebar_en"

permalink: /docs/quick-start
---

## Send Request

You can easily send a request like this:

```go
import (
    "fmt"

    dw "github.com/wnanbei/direwolf"
)

func main() {
    resp, err := dw.Get("http://httpbin.org/get")
    if err != nil {
        return
    }
    fmt.Println(resp.Text())
}
```

Output:

```json
{
  "args": {},
  "headers": {
    "Accept-Encoding": "gzip",
    "Host": "httpbin.org",
    "User-Agent": "direwolf - winter is coming"
  },
  "origin": "171.217.52.188, 171.217.52.188",
  "url": "https://httpbin.org/get"
}
```

## Add Request Parameters

Besides, direwolf provide a convenient way to add parameters to request. Such as Headers, Cookies, Params, etc.

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
