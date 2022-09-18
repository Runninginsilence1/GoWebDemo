# 自建 Go Web 的项目

# Git 迭代
每次commit的 -m 内容大概就是这些。
1. 第一次打卡：建立项目，并用 GitHub管理。
2. 第二次打卡：建立基本的Web服务，通过 Gin 的RESTful API展示一个欢迎页面。

# 打卡区
每次上传按照打卡的要求来。
基本上如果我没出错的话，每次commit就代表完成了一个打卡的任务。

+ [x] 创建一个属于自己的 GitHub仓库！
+ [x] 创建一个Web服务！
+ [ ] 让网站拥有短期的记忆！
+ [ ] 初步完成注册功能！

# 每次打卡前（Before Commit）需要进行的动作
1. 代码完成需求，并且无明显问题。
2. 打卡区打卡。
3. 迭代的地方写 提交日志。
4. 写出实现细节。
5. Commit

# 任务具体实现
## 项目集成 GitHub
比较简单。不写在这里了。


## 启动一个最基本的 Web页面
使用 Gin 完成基本的搭建。
Go标准库的网络包我应该也会用，不过因为我有需求，所以直接用 Gin。
基本的操作是：
依赖，导包，模板代码，启动项目

导入 gin的依赖：
`go get -u github.com/gin-gonic/gin`

导包：
```
import "github.com/gin-gonic/gin"
import "net/http"
```

官方文档是推荐直接用curl导入一个模板直接快速入门，不过**我用 curl会乱码**，所以暂时不考虑
```shell
curl https://raw.githubusercontent.com/gin-gonic/examples/master/basic/main.go > main.go
```

我的话使用的是后面贴出来的代码，然后通过 ResponseWriter 直接操作 响应体：
```go
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/FromZerotoExpert", func(c *gin.Context) {
		//c.JSON(200, gin.H{
		//	"message": "成功",
		//	"string":  `嗨，欢迎您来到 from zero to expert.`,
		//})
		// 拿到 Writer
		writer := c.Writer
		writer.Write([]byte("嗨，欢迎您来到 from zero to expert."))
	})
	r.Run(":8081") // 监听并在 0.0.0.0:8080 上启动服务
}

```
然后运行。

# 临时区域，做草稿的地方
Hello Gin!

## 项目的目的：
1. 学习 Go语言；
2. 通过功能导向，学会使用Go的主流框架；
3. 自己添加功能，正在做的那个 JavaWeb项目作为基础。
