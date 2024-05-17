package main

import (
	"fmt"
	"golang-examples/pprof/base2/cpu"
	"golang-examples/pprof/base2/goroutine"
	"golang-examples/pprof/base2/memory"
	"net/http"
	_ "net/http/pprof" // 导入pprof包，但因为是匿名导入，所以不会自动注册处理器
)

func main() {
	//cpu
	go func() {
		//这个值可以根据需要调整，值越大结果越精确，但计算时间也越长
		//iterations := 10000000000 // 9秒
		iterations := 1000000000000 //16分钟
		pi := cpu.CalculatePi(iterations)
		fmt.Println(pi)
	}()

	//内存测试
	memory := new(memory.Memory)
	memory.MakeOneGbSlice()

	//goroutine
	goroutine := new(goroutine.Goroutine)
	goroutine.Run()

	if err := http.ListenAndServe(":8080", nil); err != nil {
		// 处理错误
		fmt.Println("http.ListenAndServe :", err)
	}
	// http://localhost:8080/debug/pprof/
}
