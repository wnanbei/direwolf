---
layout: single
title: Session 会话

toc: true
toc_label: Content

sidebar:
  nav: "sidebar_zh"

permalink: /docs/session-zh
---

## 1. 会话 Session

在快速上手中提到的 `Get()` 等单次请求，默认使用的是短连接，不会复用连接，如果希望复用连接以提升效率的话，可以使用 `Session`。

Session 中集成了连接池，在对单个域名发起大量请求时，可以通过复用连接极大的提升效率。

```go
session := dw.NewSession()
session.Get("http://httpbin.org/get")
```

Session 同样可以使用多种请求方法，例如 `Post()`, `Request()`, `Put()` 等，请求方法中需要的参数也保持一致。

```go
session := dw.NewSession()
resp, err := session.Post("https://httpbin.org/post", dw.NewPostForm("key", "value"))
resp, err := session.Head("https://httpbin.org/head")
resp, err := session.Put("https://httpbin.org/put", dw.NewPostForm("key", "value"))
resp, err := session.Delete("https://httpbin.org/delete")
```

## 2. Session Cookies

Session 可以跨请求地自动管理请求获取的 Cookies：

```go
session := dw.NewSession()
session.Get("http://httpbin.org/cookies/set/name/direwolf")  // 获取Cookie
resp, err := session.Get("http://httpbin.org/get")
if err != nil {
    return
}
fmt.Println(resp.Text())
```

输出：

```json
{
  "args": {},
  "headers": {
    "Accept-Encoding": "gzip",
    "Cookie": "name=direwolf",
    "Host": "httpbin.org",
    "User-Agent": "direwolf - winter is coming"
  },
  "origin": "222.209.233.36, 222.209.233.36",
  "url": "https://httpbin.org/get"
}
```

### 添加 Cookies

如果需要手动添加 Cookies 到 Session 中的话，那么可以使用 `SetCookies()`方法：

```go
session := dw.NewSession()
cookies := dw.NewCookies("key", "value")
session.SetCookies("http://httpbin.org", cookies)
resp, err := session.Get("http://httpbin.org/cookies")
if err != nil {
    return
}
fmt.Println(resp.Text())
```

输出：

```json
{
  "cookies": {
    "key": "value"
  }
}
```

### 获取 Cookies

如果需要获取 Session 中的 Cookies，则可以使用 `Cookies()` 方法：

```go
session := dw.NewSession()
_, err := session.Get("http://httpbin.org/cookies/set/key/value")
if err != nil {
    return
}
cookies := session.Cookies("http://httpbin.org") // 输入cookies对应的协议和域名
fmt.Println(cookies)
```

得到的是一个 `Cookies` 类型的对象。输出：

```go
[key=value]
```

### 禁用 CookieJar

如果你想要使用 Session 以得到更高的效率，但又不需要自动管理 Cookie 的话，可以使用 `DisableCookieJar()` 这个方法禁用掉 CookieJar。

```go
session.DisableCookieJar()
```

## 2. Session 设置 Headers，Proxy，Timeout

在 Session 中可以设定一些参数，例如 Headers，Proxy，Timeout，在 Session 每次发起请求时都会带上这些参数。

### 请求头 Headers

```go
session := dw.NewSession()
headers := dw.NewHeaders("User-Agent", "Chrome 76.0")
session.Headers = headers
resp, err := session.Get("http://httpbin.org/headers")
if err != nil {
    return
}
fmt.Println(resp.Text())
```

输出：

```json
{
  "headers": {
    "Accept-Encoding": "gzip",
    "Host": "httpbin.org",
    "User-Agent": "Chrome 76.0"
  }
}
```

如果在请求方法中也传入了 Headers 参数，direwolf 会在发起请求时将其与 Session 的 Headers 合并，如果有同名 Header，则请求方法中传入的 Headers 优先。

```go
session := dw.NewSession()
sessionHeaders := dw.NewHeaders(
    "User-Agent", "Chrome 88.0",
    "session", "on",
)
session.Headers = sessionHeaders

normalHeaders := dw.NewHeaders(
    "User-Agent", "Chrome 66.0",
    "normal", "on",
)
resp, err := session.Get("http://httpbin.org/headers", normalHeaders)
if err != nil {
    return
}
fmt.Println(resp.Text())
```

输出：

```json
{
  "headers": {
    "Accept-Encoding": "gzip",
    "Host": "httpbin.org",
    "Normal": "on",
    "Session": "on",
    "User-Agent": "Chrome 66.0"
  }
}
```

### 代理 Proxy

```go
session := dw.NewSession()
proxy := &dw.Proxy{
    HTTP:  "http://127.0.0.1:12333",
    HTTPS: "http://127.0.0.1:12333",
}
session.Proxy = proxy
resp, err := session.Get("http://httpbin.org/ip")
if err != nil {
    return
}
fmt.Println(resp.Text())
```

输出：

```json
{
  "origin": "88.88.88.88, 88.88.88.88"
}
```

如果在请求方法中传入了 Proxy 参数，则优先级高于 Session 的 Proxy。

### 超时 Timeout

```go
session := dw.NewSession()
session.Timeout = 5
```

如果在请求方法中传入了 Timeout 参数，则优先级高于 Session 的 Timeout。
