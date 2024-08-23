package main

import (
	"fmt"  // 标准输出包
	"sync" // 同步包，提供同步原语如互斥锁
	"time" // 时间处理包，用于本例中的延迟等待
)

// count 是一个共享的整数变量
var count int

// mutex 是一个互斥锁，用于保护 count 的访问
var mutex sync.Mutex

// readValue 用于读取 count 的值
func readValue() {
	// 获取锁
	mutex.Lock()
	// 读取 count 的值
	value := count
	// 延迟一秒，模拟读取操作
	time.Sleep(time.Second)
	// 释放锁
	mutex.Unlock()

	// 打印读取到的值
	fmt.Printf("Read value: %d\n", value)
}

// writeValue 用于设置 count 的新值
func writeValue(newValue int) {
	// 获取锁
	mutex.Lock()
	// 设置 count 的新值
	count = newValue
	// 释放锁
	mutex.Unlock()

	// 打印写入的新值
	fmt.Printf("Wrote value: %d\n", newValue)
}

func main() {
	// 创建多个 goroutine 来读取和写入 count 的值
	go readValue()
	go readValue()
	go writeValue(10) // 写入新值 10
	go readValue()
	go readValue()

	// 等待一秒，确保所有 goroutine 都已完成
	time.Sleep(time.Second * 5)
}
