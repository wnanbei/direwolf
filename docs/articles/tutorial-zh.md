---
layout: single
title: 教程

toc: true
toc_label: Content

sidebar:
  nav: sidebar_zh

permalink: /docs/tutorial-zh
---

## 1. 发起请求

你可以像这样发起一个请求：

```go
resp, err := dw.Get("https://httpbin.org/get")
if err != nil {
    return
}
```

如果 err 等于 nil，那么你会得到一个 `Response` 对象。

其他请求方法：

```go
resp, err := dw.Post("https://httpbin.org/post", dw.NewPostForm("key", "value"))
resp, err := dw.Head("https://httpbin.org/head")
resp, err := dw.Put("https://httpbin.org/put", dw.NewPostForm("key", "value"))
resp, err := dw.Delete("https://httpbin.org/delete")
```

你还可以使用一个更加通用的函数 `Request()` 来发起请求，使用这个方法需要你先构造一个 `RequestSetting` 对象：

```go
req := dw.NewRequestSetting("Get", "https://httpbin.org/get")
resp, err := dw.Request(req)
if err != nil {
    return
}
```

使用这种方法发起请求，那么 `RequestSetting` 对象将是可以被复用的。除了需要将请求方法的字符串作为第一个参数传入之外，`NewRequestSetting()` 方法与普通的 `Get()` 和 `Post()` 等方法需要的参数是一致的。

## 2. 传递URL参数

在请求中加入URL参数非常简单，你只需要使用 `NewParams()` 创建一个URL参数对象，并将其传入请求方法中即可：

```go
params := dw.NewParams("key", "value")
resp, err := dw.Get("https://httpbin.org/get", params)
if err != nil {
    return
}
fmt.Println(resp.URL)
```

输出：

```text
https://httpbin.org/get?key=value
```

如果希望传入多个参数，可以像这样：

```go
params := dw.NewParams(
    "key1", "value1",
    "key2", "value2",
)
```

**注：记住 Key 和 Value 之间的逗号。**

**注：Key 必须与 Value 成对匹配, 如果没有的话将会报错。**

参数中有同名的 Key 是没有问题的：

```go
params := dw.NewParams(
    "key1", "value1",
    "key1", "value2",
)
```

输出：

```text
https://httpbin.org/get?key1=value1&key1=value2
```

## 3. 设置 Headers

设置 Headers 与传入URL参数非常相似, 使用 `NewHeaders()`：

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

输出：

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

这个 `NewHeaders()` 方法返回的是一个 `http.Header` 对象，如果你想要自己构造也是可以的。

如果你没有设置 `User-Agent`，direwolf 会自动使用默认的 `User-Agent`: `direwolf - winter is coming`。

## 4. 添加 Cookies

添加 Cookies 与传入URL参数也是类似的：

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

输出：

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

## 5. Post 表单

如果你想要使用 Post 方法提交表单，请使用 `NewPostForm()`：

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

输出：

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

## 6. Post 请求体

如果你想要使用 Post 直接提交数据，你可以使用 `Body`，它的原始类型是 `[]byte`，如下所示：

```go
body := dw.Body("Hello World")
resp, err := dw.Post("https://httpbin.org/post", body)
if err != nil {
    return
}
fmt.Println(resp.Text())
```

输出：

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

## 7. 设置超时

`Timeout` 指定了一个请求的超时时间，这个超时包含了连接时间、任何的重定向、和读取响应体的时间。

计时器在 Get、Head、Post 等方法返回之后仍在运行，并且可能会打断 Response.Body 的读取，在 Response.Body 读取完毕后计时结束。

- 如果 timeout > 0, 表示设置了一个超时时间。
- 如果 timeout < 0, 表示不设置超时。
- 如果 timeout = 0, 表示使用默认的30秒超时。

```go
timeout := dw.Timeout(5)
resp, err := dw.Get("https://httpbin.org/delay/10", timeout)
```

或者

```go
resp, err := dw.Get(
    "https://httpbin.org/delay/10",
    dw.Timeout(5),
)
```

## 8. 重定向

RedirectNum 是允许重定向的次数。

- 如果 RedirectNum > 0, 表示设置一个允许重定向的次数。

- 如果 RedirectNum = 0, 表示禁止重定向。

- 如果没有设置 RedirectNum, 表示默认允许5次重定向。

```go
redirect := dw.RedirectNum(10)
resp, err := dw.Get("https://httpbin.org/delay/10", redirect)
```

或者

```go
resp, err := dw.Get(
    "https://httpbin.org/delay/10",
    dw.RedirectNum(10),
)
```

## 9. 代理

设置代理同样非常简单，你可以为 HTTP 和 HTTPS 网页分别设置不同的代理：

```go
proxies := &dw.Proxy{
    HTTP:  "http://127.0.0.1:8888",
    HTTPS: "http://127.0.0.1:8888",
}
resp, err := dw.Get("https://httpbin.org/get", proxies)
if err != nil {
  return
}
fmt.Println(resp.Text())
```

## 10. Response 响应

发起请求之后，如果没有返回异常，那么你会得到一个 `Response` 对象。

你可以从 response 得到原始的请求地址：

```go
resp.URL
```

也可以获取请求的状态码, 仅数字部分:

```go
resp.StatusCode
```

获取请求返回的 headers:

```go
resp.Headers
```

获取请求返回的 cookies：

```go
resp.Cookies
```

获取得到这个响应的请求:

```go
resp.Request
```

## 11. 提取数据

你使用 direwolf 发送请求之后可以非常方便的提取数据，正如我们上面所做的一样：

```go
resp, err := dw.Get("https://httpbin.org/get")
if err != nil {
  return
}
fmt.Println(resp.Text())
```

`Text()` 会默认使用 `UTF8` 编码来解码内容，你也可以自行指定解码的编码：

```go
resp.Text("GBK")
```

目前仅支持 `UTF8`, `GBK`, `GB18030`, `Latin1` 这几种编码。

注：Text() 在你每次调用时都会解码一次响应内容，如果你希望重用 text，你最好将 text 存到一个变量中。

```go
text := resp.Text()
```

除此之外，如果你想要获取原始的 content，可以使用 `Content()` 方法，它会返回一个 `[]byte`:

```go
resp.Content()
```

## 12. 使用 CSS 选择器提取数据

### Text 文本

Direwolf 使用 `goquery` 在内部集成了 Css 选择器，可以使提取数据更加简单。

```go
text := resp.CSS("a").Text()
```

这会查找所有符合匹配的数据结果, 将其放入一个切片中并返回。如果没有找到匹配的数据，它会返回一个空切片。

在很多情况下，我们仅仅查找一个单个的匹配结果, 这样我们可以使用 `First()` 或者 `At()` 来提取单个匹配结果：

```go
text1 := resp.CSS("a").First().Text()
text2 := resp.CSS("a").At(3).Text()
```

使用这两个方法会返回单个的字符串，如果没有找到结果，会返回一个空字符串。

`Text()` 方法仅返回当前节点下的所有文本内容，不包含子节点中的文本。如果你需要所有子节点中的文本，请考虑 `TextAll()`。

```go
text := resp.CSS("a").TextAll()
```

### Attribute 属性

除了文本内容，direwolf 也可以提取属性内容：

```go
attr := resp.CSS("a").Attr("href")
```

与 `Text()` 相同，它返回一个包含属性值的切片。它也可以使用 `First()` 或者 `At()` 来提取单个数据。

`Attr()` 可以设置一个默认值，如果没有找到匹配的值，就会返回默认值。

```go
attr := resp.CSS("a").Attr("class", "default value")
```

## 13. 使用正则提取数据

Direwolf 也支持使用正则表达式提取数据，有两个方法。

这是示例数据：

```go
fmt.Println(resp.Text())
// Output:
// -Hello--World--direwolf--wnanbei-
```

首先是 `Re()`，它返回一个包含所有匹配数据的列表。

```go
fmt.Println(resp.Re("-.*?-"))
// Output:
// [-Hello- -World- -direwolf- -wnanbei-]
```

然后是 `ReSubmatch()`，它会返回一个二维列表，包含着所有的子匹配结果（正则表达式里括号中匹配的数据）。

```go
fmt.Println(resp.ReSubmatch("-(.*?)--(.*?)-"))
// Output:
// [[Hello World] [direwolf wnanbei]]
```
