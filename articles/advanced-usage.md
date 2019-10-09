---
layout: single
title: Advanced Usage

toc: true
toc_label: Content

sidebar:
  nav: "sidebar_en"

permalink: /docs/advanced-usage
---

## 1. Session Object

Request method such as `Get()` and `Post()`, which use short connection by default, do not reuse connections. You can use `Session` if you want to reuse connections for efficiency.

The session integrates `http.Client`. When made a large number of requests to a single domain, the connections can be reused to improve efficiency through the underlying connection pool.

```go
session := dw.NewSession()
session.Get("http://httpbin.org/get")
```

A Session object has all the request methods of the main Direwolf API.

```go
session := dw.NewSession()
resp, err := session.Post("https://httpbin.org/post", dw.NewPostForm("key", "value"))
resp, err := session.Head("https://httpbin.org/head")
resp, err := session.Put("https://httpbin.org/put", dw.NewPostForm("key", "value"))
resp, err := session.Delete("https://httpbin.org/delete")
```

and `Request()`:

```go
req := dw.NewRequestSetting("Get", "https://httpbin.org/get")
resp, err := session.Request(req)
```

### Parameter Priority

Session can persists some parameters across requests, like headers, timeout and proxy. But method-level parameters will override session parameters if method-level parameters coexists with session parameters.

In other words, method-level parameters has a higher priority. Example:

```go
session := dw.NewSession()
sessionHeaders := dw.NewHeaders("User-Agent", "Chrome 88.0")
session.Headers = sessionHeaders  // Set session headers

normalHeaders := dw.NewHeaders("User-Agent", "Chrome 66.0")
resp, err := session.Get("http://httpbin.org/headers", normalHeaders)
if err != nil {
    return
}
fmt.Println(resp.Text())
```

Output：

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

**However, that method-level parameters will not be persisted across requests, even if using a session.** 

## 2. Session Cookies

Session can persists cookies across requests made from it：

```go
session := dw.NewSession()
session.Get("http://httpbin.org/cookies/set/name/direwolf")  // get cookie
resp, err := session.Get("http://httpbin.org/get")
if err != nil {
    return
}
fmt.Println(resp.Text())
```

Output：

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

### Add Cookies

If you need to manually add cookies to session, you can use `SetCookies()` method.

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

Output：

```json
{
  "cookies": {
    "key": "value"
  }
}
```

### Get Cookies

If you need to get cookies from session, you can use `Cookies()` method.

```go
session := dw.NewSession()
_, err := session.Get("http://httpbin.org/cookies/set/key/value")
if err != nil {
    return
}
cookies := session.Cookies("http://httpbin.org") // Input cookies scheme and domain
fmt.Println(cookies)
```

You will get a  `Cookies` type method. Output:

```go
[key=value]
```

### Disable CookieJar

If you want use session to improve efficiency but do not need CookieJar to persist cookie，you can use `DisableCookieJar()` to disable cookiejar.

```go
session.DisableCookieJar()
```

## 3. Session set Headers, Proxy, Timeout

Session can persists some parameters across requests, like headers, timeout and proxy. But method-level parameters will override session parameters if method-level parameters coexists with session parameters.

### Headers

The type of Session Headers field is `http.Header`, you can simply construct it with `dw.NewHeaders()` and assign it.

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

Output：

```json
{
  "headers": {
    "Accept-Encoding": "gzip",
    "Host": "httpbin.org",
    "User-Agent": "Chrome 76.0"
  }
}
```

Different with other parameters, they will be merged if method-level parameters coexists with session parameters. And if there are same name parameters, the method-level parameter will override the same name session parameters.

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

Output:

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

### Proxy

The type of Session Proxy field is a struct of `*dw.Proxy`. This struct has two fields `HTTP` and `HTTPS`, which means you can set different proxy when you request HTTP or HTTPS website.

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

Output：

```json
{
  "origin": "88.88.88.88, 88.88.88.88"
}
```

### Timeout

The type of Session Timeout field is a simple `int`.

```go
session := dw.NewSession()
session.Timeout = 5
```

## 4. Session Options

There are more settings that can be changed in Session. Such as the maximum number of connections each session support, or various detailed timeouts.

Since Golang lacks support of default values, if you need modify these settings, you need to get a default `SessionOptions` object first:

```go
option := dw.DefaultSessionOptions()
```

Then you can change one or more of these settings, and pass in these settings when creating `Session` object:

```go
option.DialTimeout = 10
session := dw.NewSession(option)
```

This is default value of `DefaultSessionOptions`:

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

### Timeout

**DialTimeout** - DialTimeout is the maximum amount of time a dial will wait for a connect to complete.

When using TCP and dialing a host name with multiple IP addresses, the timeout may be divided between them. 

With or without a timeout, the operating system may impose its own earlier timeout. For instance, TCP timeouts are often around 3 minutes.

**DialKeepAlive** - DialKeepAlive specifies the interval between keep-alive probes for an active network connection.

Network protocols or operating systems that do not support keep-alives ignore this field. If negative, keep-alive probes are disabled.

**IdleConnTimeout** - IdleConnTimeout is the maximum amount of time an idle (keep-alive) connection will remain idle before closing itself.

Zero means no limit.

**TLSHandshakeTimeout** - TLSHandshakeTimeout specifies the maximum amount of time waiting to wait for a TLS handshake. Zero means no timeout.

**ExpectContinueTimeout** - ExpectContinueTimeout, if non-zero, specifies the amount of time to wait for a server's first response headers after fully writing the request headers if the request has an "Expect: 100-continue" header. 

Zero means no timeout and causes the body to be sent immediately, without waiting for the server to approve. 

This time does not include the time to send the request header.

### Connections

**MaxConnsPerHost** - MaxConnsPerHost optionally limits the total number of connections per host, including connections in the dialing, active, and idle states. On limit violation, dials will block.

Zero means no limit.

**MaxIdleConns** - MaxIdleConns controls the maximum number of idle (keep-alive) connections across all hosts. Zero means no limit.

If you need to make a large number of requests, it is recommended to increase the value of this field.

**MaxIdleConnsPerHost** - MaxConnsPerHost optionally limits the total number of connections per host, including connections in the dialing, active, and idle states. On limit violation, dials will block.

Zero means no limit.

If you need to make a large number of requests for a single domain, it is recommended to increase the value of this field.

### Other

**DisableCookieJar** - DisableCookieJar specifies whether disable session cookiejar.

**DisableDialKeepAlives** - DisableDialKeepAlives, if true, disables HTTP keep-alives and will only use the connection to the server for a single HTTP request.

This is unrelated to the similarly named TCP keep-alives.