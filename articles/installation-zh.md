---
layout: single
title: 安装
locale: zh
toc: true
toc_label: Content

sidebar:
  nav: sidebar_zh

permalink: /docs/installation-zh
---

安装 direwolf 最简单的方式是使用 `go get`：

```text
go get github.com/wnanbei/direwolf
```

**注：请在安装前确保 `GOPATH` 等环境变量已经配置完毕。**

## 安装指定版本

在最近的 2019.09.03，Golang 发布了 1.13 版本，此版本将 `Go module` 做为默认模式。

所以如果你的 Go 版本等于或高于 1.13，或者你开启了 `Go module`模式，那么你可以使用这种方法指定 direwolf 安装的版本：

```text
go get github.com/wnanbei/direwolf@v0.7.0
```

由于 Go 1.13 默认使用中国国内无法访问的 `https://proxy.golang.org` 作为 `GOPROXY`，所以要使用 `Go module`，需要将 `goproxy.cn` 或者 `goproxy.io` 设置为 `GOPROXY`。

## 手动安装

如果因为某些原因，你不能使用 `go get` 在线安装 direwolf，那么你可以使用 `git clone` 从 github 获取 direwolf 的源代码：

```text
git clone https://github.com/wnanbei/direwolf.git
```

并将其放入你的 `GOPATH/src` 目录中。

## 导入

你可以像这样导入 direwolf 以方便之后使用：

```go
import (
    dw "github.com/wnanbei/direwolf"
)
```

我会在此网站的所有文档中都使用这种方式进行演示，但这只是推荐用法，并不是强制性的。

你同样可以像正常导入其他包一样导入 direwolf :

```go
import (
    "github.com/wnanbei/direwolf"
)
```
