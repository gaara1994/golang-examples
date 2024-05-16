package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof" // 导入pprof包，但因为是匿名导入，所以不会自动注册处理器
)

func main() {
	if err := http.ListenAndServe(":8080", nil); err != nil {
		// 处理错误
		fmt.Println("http.ListenAndServe :", err)
	}
	// http://localhost:8080/debug/pprof/
}
