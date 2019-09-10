---
layout: single
title: 快速上手

toc: true
toc_label: Content

sidebar:
  nav: sidebar_zh

permalink: /docs/quick-start-zh
---

## 发起请求

你可以像下方这样非常简单的发起一个请求：

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

输出：

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

## 添加请求参数

除此之外，direwolf 可以很方便的给一个请求添加参数，例如 Headers、Cookies、Params。

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

输出：

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
