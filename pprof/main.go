package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"sync"
	"time"
)

func main() {
	go func() {
		err := http.ListenAndServe("localhost:6060", nil)
		if err != nil {
			return
		}
	}()
	go cpuRun()
	go memoryRun()
	time.Sleep(60 * time.Second)
}

func cpuRun() {
	var wg sync.WaitGroup
	maxGoroutines := runtime.NumCPU() - 1 // // 设置为CPU的核心数

	for i := 1; i <= maxGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("启动协程 #%d...\n", id)
			for {
				// 可以在这里添加一些轻量级操作，但注意即使是空循环也会消耗CPU资源
			}
		}(i) // 注意这里直接传入i的当前值，避免goroutine共享同一个变量的问题
	}

	// 等待所有goroutine启动完毕
	wg.Wait()
	fmt.Println("所有协程都已启动完毕！")
}

func memoryRun() {
	const MB = 1 << 20 // 1 MB = 1 << 20 bytes
	var wg sync.WaitGroup

	for i := 0; ; i++ { // 循环进行分配和释放操作
		wg.Add(1)
		go func(iteration int) {
			defer wg.Done()
			allocateMemory(500 * MB) // 分配500MB的内存
			releaseMemory()          // 尝试释放内存（依赖于GC）
		}(i)

		if i%10 == 0 { // 每10次循环后等待所有goroutine完成
			wg.Wait()
			fmt.Println("本轮的所有内存分配任务都已完成。")
			time.Sleep(time.Second) // 等待一段时间再继续下一轮，以便观察效果
		}
	}

	select {} // 阻塞主线程，使得程序不会结束
}

func allocateMemory(size int) {
	b := make([]byte, size)
	fmt.Printf("已分配%dMb内存\n", len(b)/1024/1024)
}

func releaseMemory() {
	runtime.GC() // 强制执行垃圾回收以尝试释放内存
}
