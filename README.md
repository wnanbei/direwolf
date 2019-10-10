# Direwolf HTTP Client: Save your time

Package direwolf is a convenient and easy to use http client written in Golang.

If you want find more info, please go here: [Direwolf HTTP Client: Save your time](https://wnanbei.github.io/direwolf/)，内有中文文档。

![direwolf](docs/assets/cover.jpg)

## Feature Support

- Clean and Convenient API
- Simple to Set Headers, Cookies, Parameters, Post Forms
- Sessions Control
- Keep-Alive & Connection Pooling
- HTTP(S) Proxy Support
- Redirect Control
- Timeout Control
- Support extract result from response body with css selector, regexp, json
- Content Decoding
- More to come...

## Installation

```text
go get github.com/wnanbei/direwolf
```

## Take a glance

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

Besides, direwolf provide a convenient way to add parameters to request. Such
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

## Contribute

Direwolf is a personal project now, but all contributions, bug reports, bug fixes, documentation improvements, enhancements, and ideas are welcome.

If you find a bug in direwolf or have some good ideas:

 - Go to [GitHub “issues” tab](https://github.com/wnanbei/direwolf/issues) and open a fresh issue to start a discussion around a feature idea or a bug.
 - Send a pull request and bug the maintainer until it gets merged and published.
 - Write a test which shows that the bug was fixed or that the feature works as expected.

If you need to discuss about direwolf with me, you can send me a e-mail.
