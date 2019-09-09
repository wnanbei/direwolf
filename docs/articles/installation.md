---
layout: single
title: Installation

toc: true
toc_label: Content

sidebar:
  nav: sidebar_en

permalink: /docs/installation
---

The easist way to install direwolf is to use `go get`:

```text
go get github.com/wnanbei/direwolf
```

**Note: Please make sure that `GOPATH` and other environment have been setted before installation.**

## Install the Specified Version

Golang released version `1.13` in recent 2019.09.03. This version uses `Go module` as the default mode.

So if your Golang version is equal to or higher than `1.13`, or if you enabled `Go module` mode, you can use this method to specific the version of direwolf installed:

```text
go get github.com/wnanbei/direwolf@v0.4.0
```

## Manual installation

If you can\`t install direwolf online with `go get` because of some reasons. You can use `git clone` to get source code of direwolf from github:

```text
git clone https://github.com/wnanbei/direwolf.git
```

Then move it to your `GOPATH/src` directory.

## Import

You can import `direwolf` like this for later use:

```go
import (
    dw "github.com/wnanbei/direwolf"
)
```

I will use this way to demonstrate in all documentation on this site. But this is only a recommended usage, not mandatory.

You can normally import direwolf like other package:

```go
import (
    "github.com/wnanbei/direwolf"
)
```
