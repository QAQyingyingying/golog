# golog - 简易的彩色日志打印工具

golog 是一个用于在 Go 项目中进行彩色日志打印的简易工具。

## 安装

使用以下命令来安装 golog：

```shell
go get github.com/QAQyingyingying/golog
```
## 使用方法
以下是一个简单的例子，展示如何在你的 Go 项目中使用 golog
```golang
package main

import "github.com/QAQyingyingying/golog"

func main() {
	golog.Info("This is an info message")
	golog.Warn("This is a warning message")
	golog.Error("This is an error message")
	golog.ErrorAndExit("This is an error message and will exit the program")
}
```
## 效果
![image](https://github.com/QAQyingyingying/golog/assets/49606783/75606725-38bb-4564-a7aa-b56acff32d1d)

