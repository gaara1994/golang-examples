package main

import (
	"github.com/gaara1994/demo/pprof"
	"github.com/gin-gonic/gin"
	_ "net/http/pprof" // 导入pprof包，但因为是匿名导入，所以不会自动注册处理器
)

func main() {
	// 创建一个默认的路由引擎并命名为router
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	pprof.UseByGin(router)
	// 启动服务
	router.Run()
	// 默认监听 :8080，也可以指定如 router.Run(":8081")
}