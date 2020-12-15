# wine

wine是一个 [gin](https://github.com/gin-gonic/gin) 中间件，它简化了参数绑定和返回功能

## Installation

1. 首先你必须要安装Go语言并且设置Go语言工作环境，并且安装gin框架

```shell
$ go get -u github.com/helloteemo/wine
```

2. Import

```go
import "gitub.com/helloteemo/wine"
```

## Quick Start

你可以查看 `demo/main.go` 文件，这个文件展示了wine的所有功能

这个函数声明完全展示了wine的所有功能

```
func Print(c *gin.Context, req User, req2 *User) interface{}
```

### 行参

在这里我们可以省下参数绑定的代码，直接把你的参数写在行参上，wine会帮你自动把参数绑定上去，

**目前只实现了参数绑定和`*gin.Context`的传递，其它的以后再写吧**

### 返回值

返回值如果是 `wine.ErrorResult` 的话，那么就会返回错误的信息，即 `demo.go` 中的 `Error` 函数

否则参数会响应到 `data` 上
