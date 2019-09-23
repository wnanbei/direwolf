---
layout: single
title: 高级用法

toc: true
toc_label: Content

sidebar:
  nav: "sidebar_zh"

permalink: /docs/advanced-usage-zh
---

## 1. 会话 Session

`Get()`，`Post()` 等请求方法，默认使用的是短连接，不会复用连接，如果希望复用连接以提升效率的话，可以使用 `Session`。

Session 中集成了 `http.Client`，通过其底层的连接池，在对单个域名发起大量请求时，可以复用连接来极大的提升效率。

```go
session := dw.NewSession()
session.Get("http://httpbin.org/get")
```

Session 对象拥有 Direwolf API 所有的请求方法。

```go
session := dw.NewSession()
resp, err := session.Post("https://httpbin.org/post", dw.NewPostForm("key", "value"))
resp, err := session.Head("https://httpbin.org/head")
resp, err := session.Put("https://httpbin.org/put", dw.NewPostForm("key", "value"))
resp, err := session.Delete("https://httpbin.org/delete")
```

以及 `Request()`：

```go
req := dw.NewRequestSetting("Get", "https://httpbin.org/get")
resp, err := session.Request(req)
```

### 参数优先级

Session 可以跨请求地保持某些参数，例如 headers，超时和代理。但是如果方法的参数和 Session 的参数共存的话，方法的参数会覆盖掉 Session 的参数。

换句话说，方法的参数拥有更高的优先级。例子：

```go
session := dw.NewSession()
sessionHeaders := dw.NewHeaders("User-Agent", "Chrome 88.0")
session.Headers = sessionHeaders  // 设置Session的Headers

normalHeaders := dw.NewHeaders("User-Agent", "Chrome 66.0")
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
    "User-Agent": "Chrome 66.0"
  }
}
```

**但是，方法的参数将不会被跨请求的保持，它仅会被使用一次，即使使用的是 Session**

## 2. Session Cookies

Session 可以跨请求地自动管理请求获取的 Cookies：

```go
session := dw.NewSession()
session.Get("http://httpbin.org/cookies/set/name/direwolf")  // 获取cookie
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

## 3. Session 设置 Headers，Proxy，Timeout

Session 可以跨请求地保持某些参数，例如 headers，超时和代理。但是如果方法的参数和 Session 的参数共存的话，方法的参数会覆盖掉 Session 的参数。

### 请求头 Headers

Session 的 Headers 字段类型为 `http.Header`, 使用 `dw.NewHeaders()` 方法构造并赋值即可。

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

与其他参数不同，如果方法的 Headers 和 Session 的 Headers 共存，那么它们将会被合并，而如果有同名的 Header，那么方法的 Header 将会覆盖掉同名的 Session 的 Header。

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

Session 的 Proxy 字段类型为 `*dw.Proxy` 的一个结构体，这个结构体有两个字段 `HTTP` 和 `HTTPS`，表示你访问 HTTP 和 HTTPS 网页时可以分别设置不同的代理。

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

### 超时 Timeout

Session 的 Timeout 字段类型为一个简单的整形。

```go
session := dw.NewSession()
session.Timeout = 5
```
