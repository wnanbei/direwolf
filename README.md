# Direwolf HTTP Client: Save your time

Package direwolf is a convient and esay to use http client written in Golang. 

## Feature Support

- Clean and Convient API
- Simple to Set Headers, Cookies, Parameters, Post Forms 
- Sessions with Cookie Persistence
- Elegant Key/Value Cookies
- Keep-Alive & Connection Pooling
- HTTP(S) Proxy Support
- Redirect Control
- Timeout Control
- Support extract result from response body with css selector, regexp, json
- Content Decoding
- More to come

## Installation

```
go get github.com/wnanbei/direwolf
```

## Quick Start

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

## How to Use

### 1. Make Request

First, you can start a request like this:

```go
import (
    dw "github.com/wnanbei/direwolf"
)

func main() {
    resp, err := dw.Get("https://httpbin.org/get")
}
```

Other HTTP request types: 

```go
resp, err := dw.Post("https://httpbin.org/post", dw.NewPostForm("key", "value"))
resp, err := dw.Head("https://httpbin.org/head")
resp, err := dw.Put("https://httpbin.org/put", dw.NewPostForm("key", "value"))
resp, err := dw.Delete("https://httpbin.org/delete")
```

### 2. Passing Parameters In URLs

Passing parameters in URLs is very easy, you only need to new a Params and pass it to request method.

```go
import (
    "fmt"

    dw "github.com/wnanbei/direwolf"
)

func main() {
    params := dw.NewParams("key", "value")
    resp, err := dw.Get("https://httpbin.org/get", params)
    if err != nil {
        return
    }
    fmt.Println(resp.URL)
}
```

Output:

```
https://httpbin.org/get?key=value
```

If you want pass more parameters to URLs, just like this:

```go
params := dw.NewParams(
    "key1", "value1",
    "key2", "value2",
)
```

**Note: Remember the comma between key and value.** 

**Note: Key must to match Value one by one, if not, will report an error.**

If the parameters have the same key, it`s ok:

```go
params := dw.NewParams(
    "key1", "value1",
    "key1", "value2",
)
```

Output:

```
https://httpbin.org/get?key1=value1&key1=value2
```

### 3. Set Headers

Set headers is similar to add parameters, use `NewHeaders()`:

```go
import (
    "fmt"

    dw "github.com/wnanbei/direwolf"
)

func main() {
    headers := dw.NewHeaders(
        "key", "value",
        "User-Agent", "direwolf",
    )
    resp, err := dw.Get("https://httpbin.org/get", headers)
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
    "Key": "value",
    "User-Agent": "direwolf"
  },
  "origin": "1.1.1.1, 1.1.1.1",
  "url": "https://httpbin.org/get"
}
```

If you did not set `User-Agent`, direwolf will use default `User-Agent`: `direwolf - winter is coming`.

### 4. Add Cookies

Add cookies is similar to add parameters, too.

```go
import (
	"fmt"

	dw "github.com/wnanbei/direwolf"
)

func main() {
	cookies := dw.NewCookies(
		"key1", "value1",
		"key2", "value2",
	)
	resp, err := dw.Get("https://httpbin.org/get", cookies)
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
import (
	"fmt"

	dw "github.com/wnanbei/direwolf"
)

func main() {
	postForm := dw.NewPostForm(
		"uid", "123456789",
		"pw", "666888",
	)
	resp, err := dw.Post("https://httpbin.org/post", postForm)
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
import (
	"fmt"

	dw "github.com/wnanbei/direwolf"
)

func main() {
	body := dw.Body("Hello World")
	resp, err := dw.Post("https://httpbin.org/post", body)
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

```go
import (
	"fmt"

	dw "github.com/wnanbei/direwolf"
)

func main() {
	timeout := dw.Timeout(5)
	resp, err := dw.Get("https://httpbin.org/delay/10", timeout)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Text())
}
```

