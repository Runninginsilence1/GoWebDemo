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
