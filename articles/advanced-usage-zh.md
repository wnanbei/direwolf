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

以及 `Send()`：

```go
req, err := dw.NewRequest("Get", "https://httpbin.org/get")
if err != nil {
    return
}
resp, err := session.Send(req)
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

## 4. Session 高级设置

在 Session 中还有更多可以详细更改的设置，例如每个 Session 可以支持的最大连接数量，或者各类详细的超时时间。

由于 Golang 缺乏对默认值的支持，如果你需要修改这些设置，那么你需要像这样先获取一个默认的 `SessionOptions` 对象:

```go
option := dw.DefaultSessionOptions()
```

然后你就可以更改其中的某个或多项设置，并在创建 `Session` 时将设置传入:

```go
option.DialTimeout = 10
session := dw.NewSession(option)
```

这是 `DefaultSessionOptions` 的默认值:

```go
&SessionOptions{
    DialTimeout: 30 * time.Second,		
    DialKeepAlive: 30 * time.Second,		
    MaxConnsPerHost: 0,		
    MaxIdleConns: 100,		
    MaxIdleConnsPerHost: 2,		
    IdleConnTimeout: 90 * time.Second,		
    TLSHandshakeTimeout: 10 * time.Second,		
    ExpectContinueTimeout: 1 * time.Second,		
    DisableCookieJar: false,		
    DisableDialKeepAlives: false,	
}
```

### Timeout 超时

**DialTimeout** - 建立一个新连接的超时时间。默认值为 30。如果在访问一个域名的过程中需要与多个 IP 地址建立连接，`DialTimeout` 将会被分成多个部分。

**DialKeepAlive** - 每次探测活跃连接是否 keep-alives 的时间间隔。如果网络协议或系统不支持 keep-alives，则将会忽略此字段。如果为负值，将会禁用 keep-alives 探测。

**IdleConnTimeout** - 连接池中持久连接保持在空闲状态的时间，如果超过这个时间，持久连接将会被关闭。值为 `0` 表示不限制超时。

**TLSHandshakeTimeout** - TLS 握手的超时时间。值为 `0` 表示不限制超时。

**ExpectContinueTimeout** - 如果请求的 header 中有 `Expect: 100-continue`，从完全发送请求的 header 开始，直到服务端返回响应的第一条 header 的超时时间。如果值为 `0`，表示发送请求后，不等待服务端的响应，立即发送响应体。

### Connections 连接数量

**MaxConnsPerHost** - 与每个域名建立的最大连接数，包括正在拨号、活跃的、空闲的连接。如果超过此连接数量，那么拨号将会堵塞。值为 `0` 表示不限制。

**MaxIdleConns** - 控制 Session 持久连接的最大总数量。值为 `0` 表示不限制。

如果需要发起大量请求，推荐提高此项设置的值。

**MaxIdleConnsPerHost** - 控制 Session 与每个域名的持久连接的最大数量。值为 `0` 将会使用默认值 `2`。

如果需要对单一域名发起大量请求，推荐提高此项设置的值。

### 其他

**DisableCookieJar** - Session 默认会启用 CookieJar 来管理 Cookie，如果你不希望启用 CookieJar，可以将此项设置改为 `true` 来禁用 CookieJar。

**DisableDialKeepAlives** - 默认情况下，HTTP1.1 会启用持久连接。如果你希望禁用持久连接，强制使用短连接，可以将此项设置改为 `true` 来禁用持久连接。