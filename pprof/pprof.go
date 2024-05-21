package pprof

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/pprof"
)

// 将标准库（stdlib）中的中间件适配到 Gin 框架中
func adapter(f func(http.ResponseWriter, *http.Request)) gin.HandlerFunc {
	return func(c *gin.Context) {
		f(c.Writer, c.Request)
	}
}

// UseByGin 适用于Gin框架的web程序
func UseByGin(router *gin.Engine) {
	router.GET("/debug/pprof/", adapter(pprof.Index))
	router.GET("/debug/pprof/cmdline", adapter(pprof.Cmdline))
	router.GET("/debug/pprof/profile", adapter(pprof.Profile))
	router.GET("/debug/pprof/symbol", adapter(pprof.Symbol))
	router.POST("/debug/pprof/symbol", adapter(pprof.Symbol))
	router.GET("/debug/pprof/trace", adapter(pprof.Trace))

	// 以下使用了pprof.Handler，它为不同的分析类型提供了处理器
	router.GET("/debug/pprof/heap", adapter(pprof.Handler("heap").ServeHTTP))
	router.GET("/debug/pprof/goroutine", adapter(pprof.Handler("goroutine").ServeHTTP))
	router.GET("/debug/pprof/block", adapter(pprof.Handler("block").ServeHTTP))
	router.GET("/debug/pprof/mutex", adapter(pprof.Handler("mutex").ServeHTTP))
	router.GET("/debug/pprof/allocs", adapter(pprof.Handler("allocs").ServeHTTP))
	router.GET("/debug/pprof/threadcreate", adapter(pprof.Handler("threadcreate").ServeHTTP))
	// 注意：'cmdline' 和 'symbol' 有特定的处理函数，不需要使用pprof.Handler
}

// Use 适用于非web程序
func Use(port string) {
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("http.ListenAndServe :", err)
	}
}
