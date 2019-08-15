# Direwolf HTTP 客户端

Direwolf 是一个由 Golang 编写的简单易用的 HTTP 客户端。

![direwolf](cover.png)

## 目录

- [Feature Support](#Feature Support)
- [Installation](#Installation)
- [Quick Start](#Quick Start)
- [How to Use](#How to Use)
  - [1. Make Request](#1. Make Request)
  - [2. Passing Parameters In URLs](#2. Passing Parameters In URLs)
  - [3. Set Headers](#3. Set Headers)
  - [4. Add Cookies](#4. Add Cookies)
  - [5. Post Form](#5. Post Form)
  - [6. Post Body](#6. Post Body)
  - [7. Set Timeout](#7. Set Timeout)
  - [8. Redirect](#8. Redirect)
  - [9. Proxy](#9. Proxy)
  - [10. Response](#10. Response)
  - [11. Extract Data](#11. Extract Data)
  - [12. Extract Data by CSS Selector](#12. Extract Data by CSS Selector)
  - [13. Extract Data by Regexp](#13. Extract Data by Regexp)

## 支持特性

- 干净方便的 API
- 设置 Headers, Cookies, URL参数, Post表单非常简单
- 支持 Cookie 管理的 Session
- Keep-Alive 和连接池
- 支持 HTTP(S) 代理
- 重定向控制
- 超时控制
- 支持使用正则表达式、CSS选择器从响应提取内容
- 响应内容解码
- 更多...

## 安装

```
go get github.com/wnanbei/direwolf
```

## 速览

你可以像下方这样非常简单的发起一个请求：

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

输出:

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

## 基础用法

首先，你可以像这样导入 direwolf 以方便之后使用：

```go
import (
    dw "github.com/wnanbei/direwolf"
)
```

这只是推荐用法。

### 1. 发起请求

你可以像这样发起一个请求：

```go
resp, err := dw.Get("https://httpbin.org/get")
if err != nil {
    return
}
```

如果 err 等于 nil，那么你会得到一个 `Response`。

其他请求方法：

```go
resp, err := dw.Post("https://httpbin.org/post", dw.NewPostForm("key", "value"))
resp, err := dw.Head("https://httpbin.org/head")
resp, err := dw.Put("https://httpbin.org/put", dw.NewPostForm("key", "value"))
resp, err := dw.Delete("https://httpbin.org/delete")
```

### 2. 传递URL参数

在请求中加入URL参数非常简单，你只需要使用 `NewParams()` 创建一个URL参数对象，并将其传入请求方法中即可：

```go
params := dw.NewParams("key", "value")
resp, err := dw.Get("https://httpbin.org/get", params)
if err != nil {
    return
}
fmt.Println(resp.URL)
```

输出:

```
https://httpbin.org/get?key=value
```

如果希望传入多个参数，可以像这样:

```go
params := dw.NewParams(
    "key1", "value1",
    "key2", "value2",
)
```

**注: 记住 Key 和 Value 之间的逗号.** 

**注: Key 必须与 Value 成对匹配, 如果没有的话将会报错.**

参数中有同名的 Key 是没有问题的：

```go
params := dw.NewParams(
    "key1", "value1",
    "key1", "value2",
)
```

输出:

```
https://httpbin.org/get?key1=value1&key1=value2
```

### 3. 设置Headers

设置 Headers 与传入URL参数非常相似, 使用`NewHeaders()`:

```go
headers := dw.NewHeaders(
    "key", "value",
    "User-Agent", "direwolf",
)
resp, err := dw.Get("https://httpbin.org/get", headers)
if err != nil {
    return
}
fmt.Println(resp.Text())
```

Output:

```json
{
  "args": {},
  "headers": {
    "Accept-Encoding": "gzip",
    "Host": "httpbin.org",
    "Key": "value",
    "User-Agent": "direwolf"
  },
  "origin": "1.1.1.1, 1.1.1.1",
  "url": "https://httpbin.org/get"
}
```

This method will return a `http.Header`, it`s ok if you want to construct it by yourself.

If you did not set `User-Agent`, direwolf will use default `User-Agent`: `direwolf - winter is coming`.

### 4. Add Cookies

Add cookies is similar to add parameters, too.

```go
cookies := dw.NewCookies(
    "key1", "value1",
    "key2", "value2",
)
resp, err := dw.Get("https://httpbin.org/get", cookies)
if err != nil {
    return
}
fmt.Println(resp.Text())
```

Output:

```json
{
  "args": {},
  "headers": {
    "Accept-Encoding": "gzip",
    "Cookie": "key2=value2; key1=value1",
    "Host": "httpbin.org",
    "User-Agent": "direwolf - winter is coming"
  },
  "origin": "1.1.1.1, 1.1.1.1",
  "url": "https://httpbin.org/get"
}
```

### 5. Post Form

If you want post form data, use `direwolf.NewPostForm()`:

```go
postForm := dw.NewPostForm(
    "uid", "123456789",
    "pw", "666888",
)
resp, err := dw.Post("https://httpbin.org/post", postForm)
if err != nil {
    return
}
fmt.Println(resp.Text())
```

Output:

```json
{
  "args": {},
  "data": "",
  "files": {},
  "form": {
    "pw": "666888",
    "uid": "123456789"
  },
  "headers": {
    "Accept-Encoding": "gzip",
    "Content-Length": "23",
    "Content-Type": "application/x-www-form-urlencoded",
    "Host": "httpbin.org",
    "User-Agent": "direwolf - winter is coming"
  },
  "json": null,
  "origin": "1.1.1.1, 1.1.1.1",
  "url": "https://httpbin.org/post"
}
```

### 6. Post Body

If you want post bytes type data, you can use `direwolf.Body`, its original type is `[]byte`, like this:

```go
body := dw.Body("Hello World")
resp, err := dw.Post("https://httpbin.org/post", body)
if err != nil {
    return
}
fmt.Println(resp.Text())
```

Output:

```json
{
  "args": {},
  "data": "Hello World",
  "files": {},
  "form": {},
  "headers": {
    "Accept-Encoding": "gzip",
    "Content-Length": "11",
    "Host": "httpbin.org",
    "User-Agent": "direwolf - winter is coming"
  },
  "json": null,
  "origin": "1.1.1.1, 1.1.1.1",
  "url": "https://httpbin.org/post"
}
```

### 7. Set Timeout

`Timeout` specifies a time limit for request. The timeout includes connection time, any redirects, and reading the response body. 

The timer remains running after Get, Head, Post, or Do return and will interrupt reading of the Response.Body.

- if timeout > 0, it means a time limit for requests.
- if timeout < 0, it means no limit.
- if timeout = 0, it means keep default 30 second timeout.

```go
timeout := dw.Timeout(5)
resp, err := dw.Get("https://httpbin.org/delay/10", timeout)
```

or

```go
resp, err := dw.Get(
    "https://httpbin.org/delay/10",
    dw.Timeout(5),
)
```

### 8. Redirect

RedirectNum is the number of request redirect allowed.

- If RedirectNum > 0, it means a redirect number limit for requests.

- If RedirectNum = 0, it means ban redirect.

- If RedirectNum is not set, it means default 5 times redirect limit.

```go
redirect := dw.RedirectNum(10)
resp, err := dw.Get("https://httpbin.org/delay/10", redirect)
```

or

```go
resp, err := dw.Get(
    "https://httpbin.org/delay/10",
    dw.RedirectNum(5),
)
```

### 9. Proxy

Set proxy is esay, too. You can set different proxies for HTTP and HTTPS sites.

```go
proxies := dw.Proxy{
    HTTP:  "http://127.0.0.1:8888",
    HTTPS: "http://127.0.0.1:8888",
}
resp, err := dw.Get("https://httpbin.org/get", proxies)
if err != nil {
	return
}
fmt.Println(resp.Text())
```

### 10. Response

After request, you will get a `Response` object if no error return.

You can get Original request url from response:

```go
resp.URL
```

You can also get request status code from response, only numbers:

```go
resp.StatusCode
```

Get response headers:

```go
resp.Headers
```

If you want get request of response:

```go
resp.Request
```

### 11. Extract Data

You can easily extract data using direwolf after sending a request, as we did above:

```go
resp, err := dw.Get("https://httpbin.org/get")
if err != nil {
	return
}
fmt.Println(resp.Text())
```

`Text()` will use `UTF8` to decode content by default. You can also specify decode method yourself:

```go
resp.Text("GBK")
```

It only support `UTF8`, `GBK`, `GB18030`, `Latin1` now.

Note: Text() will decode content everytime you call it. If you want reuse content, you would better store the content in a variable.

```go
text := resp.Text()
```

Besides, if you want get raw content, you can use `Content()` method, it will return a `[]byte`:

```go
resp.Content()
```

### 12. Extract Data by CSS Selector

#### Text

Direwolf has a built-in Css selector via `goquery`, which makes it easy to extract data.

```go
text := resp.CSS("a").Text()
```

It will find all matching values, put them in a slice and return it. If no matching values was found, it will return a empty slice.

In many cases, we only looking for a single match result, then we can use `First()` or `At()` to extract single result:

```go
text1 := resp.CSS("a").First().Text()
text2 := resp.CSS("a").At(3).Text()
```

Using these two methods will return a single string. If no matching values was found, it will return a empty string.

`Text()` method only return the text under the current node, not contains the text in the child node. If you need text in all child node, consider `TextAll()`.

```go
text := resp.CSS("a").TextAll()
```

#### Attribute

In addition to text, direwolf can alse extract attributes.

```go
attr := resp.CSS("a").Attr("href")
```

The same with `Text()`, it retrun a slice of attribute values. It can use `First()` or `At()` to extract single value, too.

`Attr()` can set a default value, if not match value is found, it will return the default value.

```go
attr := resp.CSS("a").Attr("class", "default value")
```

### 13. Extract Data by Regexp

Except css selector, direwolf also integrates regular expressions to extract data. It has two methods.

This is sample data:

```go
fmt.Println(resp.Text())
// Output:
// -Hello--World--direwolf--wnanbei-
```

First is `Re()`, it returns a slice of all match strings.

```go
fmt.Println(resp.Re("-.*?-"))
// Output:
// [-Hello- -World- -direwolf- -wnanbei-]
```

Then is `ReSubmatch()`, it will return a two-dimensional slice that contains all the sub matching results (Data only in brackets).

```go
fmt.Println(resp.ReSubmatch("-(.*?)--(.*?)-"))
// Output:
// [[Hello World] [direwolf wnanbei]]
```

***

## How to Contribute

Because I am not a native English speaker, please tell me in the Issues if there is something wrong or unclear in the introduction.

- Check for open issues or open a fresh issue to start a discussion around a feature idea or a bug.
- Send a pull request and bug the maintainer until it gets merged and published.
- Write a test which shows that the bug was fixed or that the feature works as expected.