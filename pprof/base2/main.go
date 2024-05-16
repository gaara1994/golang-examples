package main

import (
	"fmt"
	"golang-examples/pprof/base2/cpu"
	"math"
	"net/http"
	_ "net/http/pprof" // 导入pprof包，但因为是匿名导入，所以不会自动注册处理器
	"time"
)

func main() {

	go func() {
		fmt.Println("Start Calculated PI")
		startTime := time.Now()
		//这个值可以根据需要调整，值越大结果越精确，但计算时间也越长
		//iterations := 10000000000 // 9秒
		iterations := 1000000000000 //16m
		pi := cpu.CalculatePi(iterations)
		duration := time.Since(startTime)

		fmt.Printf("Calculated PI: %.15f\n", pi)
		fmt.Printf("Math library PI: %.15f\n", math.Pi)
		fmt.Printf("Time taken: %s\n", duration)
	}()


	if err := http.ListenAndServe(":8080", nil); err != nil {
		// 处理错误
		fmt.Println("http.ListenAndServe :", err)
	}
	// http://localhost:8080/debug/pprof/
}
