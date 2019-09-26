# logfox

## 简介

logfox是一个简单的go log库，可以按小时自动切分日志，设置日志保存的期限

**V1.1使用go mod做依赖管理**

## 安装

```
go get github.com/zer0131/logfox
```

## 使用示例

```
package main

import "github.com/zer0131/logfox"

func init() {
    logfox.Init("./log", "test", "DEBUG", 7)
}

func main() {
    logfox.Info("test info")
    logfox.Close()
}
```

