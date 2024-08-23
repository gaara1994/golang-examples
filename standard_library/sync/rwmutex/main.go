package main

import (
	"fmt"
	"sync"
	"time"
)

var count int
var rwMutex sync.RWMutex

func readValue() {
	// 获取读锁
	rwMutex.RLock()
	// 读取 count 的值
	value := count
	// 模拟读取操作的延迟
	time.Sleep(time.Second)
	// 释放读锁
	rwMutex.RUnlock()

	// 打印读取到的值
	fmt.Printf("Read value: %d\n", value)
}

func writeValue(newValue int) {
	// 获取写锁
	rwMutex.Lock()
	// 设置 count 的新值
	count = newValue
	// 释放写锁
	rwMutex.Unlock()

	// 打印写入的新值
	fmt.Printf("Wrote value: %d\n", newValue)
}

func main() {
	go readValue()
	go readValue()
	go writeValue(10)
	go readValue()
	go readValue()

	// 等待所有 goroutine 完成
	time.Sleep(time.Second * 5)
}
