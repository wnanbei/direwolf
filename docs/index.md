---
layout: index
title: "Direwolf HTTP Client: Save your time"
sidebar_link: true
---

Package direwolf is a convient and esay to use http client written in Golang. 

[中文文档请点此处](/direwolf/ch)

![direwolf](cover.png)

## Contents

- [Feature Support](#Feature-Support)
- [Installation](#Installation)
- [Getting started](#Getting-started)
- [Quick Start](#Quick-Start)
  - [1. Make Request](#1.-Make-Request)
  - [2. Passing Parameters In URLs](#2.-Passing-Parameters-In-URLs)
  - [3. Set Headers](#3.-Set-Headers)
  - [4. Add Cookies](#4.-Add-Cookies)
  - [5. Post Form](#5.-Post-Form)
  - [6. Post Body](#6.-Post-Body)
  - [7. Set Timeout](#7.-Set-Timeout)
  - [8. Redirect](#8.-Redirect)
  - [9. Proxy](#9.-Proxy)
  - [10. Response](#10.-Response)
  - [11. Extract Data](#11.-Extract-Data)
  - [12. Extract Data by CSS Selector](#12.-Extract-Data-by-CSS-Selector)
  - [13. Extract Data by Regexp](#13.-Extract-Data-by-Regexp)
- [Advanced Usage](#Advanced-Usage)
  - [1. Session](#1.-Session)
  - [2. Session set Headers, Porxy, Timeout](#2.-Session-set-Headers,-Porxy,-Timeout)
- [How to Contribute](#How-to-Contribute)

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