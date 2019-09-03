---
layout: single
title: "Direwolf HTTP Client: Save your time"
toc: true
toc_label: "Direwolf"

excerpt: "Package direwolf is a convient and esay to use http client written in Golang."
header:
  overlay_image: cover.png

author_profile: true
sidebar:
  nav: "docs"
---

## Feature Support

- Clean and Convient API
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

```
go get github.com/wnanbei/direwolf
```

## Getting started

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




***

## How to Contribute

Because I am not a native English speaker, please tell me in the Issues if there is something wrong or unclear in the introduction.

- Open a fresh issue to start a discussion around a feature idea or a bug.
- Send a pull request and bug the maintainer until it gets merged and published.
- Write a test which shows that the bug was fixed or that the feature works as expected.