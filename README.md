# golog - 简易的彩色日志打印工具

golog 是一个用于在 Go 项目中进行彩色日志打印的简易工具。

## 安装

使用以下命令来安装 golog：

```shell
go get gitee.com/Mystic_wangl/golog
```
## 使用方法
以下是一个简单的例子，展示如何在你的 Go 项目中使用 golog
```go
package main

import (
    "gitee.com/Mystic_wangl/golog"
)

func main() {
    logger := golog.NewLogger()

    logger.Red("This is an error message.")
    logger.Green("This is a success message.")
    logger.Yellow("This is a warning message.")
    logger.Blue("This is an informational message.")
}

```