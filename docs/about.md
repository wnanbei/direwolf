---
layout: single
title: About

toc: true
toc_label: "Content"

author_profile: true
---

## Introduction

As described in the title, direwolf is a library written in Golang and is dedicated to the convient and esay to use HTTP client.

When i first learned about Golang, i was suprised by its powerful Goroutine and complete support for network programming.But i alse found that Golang does not have a relatively complete library similar to Requests in Python. The native `net/http` is pretty good, but it is always need many annoying steps to set a simple configuration while send a request.

In the third-part libraries, it seems that most people are more concerned about performance than ease of use, such as very powerful library `fasthttp`.

But i think that ease of use is also important. Sometimes we may not need such strong performance, this is my original intention to develop direwolf.

正如标题所描述的那样，direwolf 是一个使用 Golang 编写的，致力于简单易用的 HTTP 客户端。

当我在刚接触 Golang 时，惊喜于其强大的 Goroutine 功能和完善的对网络编程的支持，但是我同时也发现，Golang 还没有一个比较完善的类似于 Python 中的 Requests 一样的库。原生的 `net/http` 虽然很不错，但是在发起网络请求时，常常一个简单的设置需要非常多麻烦的步骤来实现。

在第三方库中，好像大部分人也更关注性能，而不是易用性，例如非常厉害的 `fasthttp`。

但我认为，易用性同样也非常重要，有一些时候我们或许不需要那么出色的性能，所以这是我开发 direwolf 的初衷。

## License

Direwolf is released under terms of MIT License.

>MIT License
>
>Copyright (c) 2019 wnanbei
>
>Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
>
>The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
>
>THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
