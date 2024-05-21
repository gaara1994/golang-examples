package main

import (
	"github.com/gaara1994/demo/pprof"
	_ "net/http/pprof" // 导入pprof包，但因为是匿名导入，所以不会自动注册处理器
)

func main() {
	pprof.Use("7070")
}