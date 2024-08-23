package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var ready = false //状态是否准备好的标志
	var data int      // 生产的数据

	var mutex sync.Mutex            // 互斥锁
	var cond = sync.NewCond(&mutex) // 条件变量

	// 生产者 goroutine
	go func() {
		mutex.Lock()                // 获取锁
		time.Sleep(3 * time.Second) // 模拟数据准备时间
		data = 99                   // 生产数据
		fmt.Println("数据已准备好")
		ready = true   // 通知消费者数据已准备好
		cond.Signal()  // 发送通知
		mutex.Unlock() // 解锁
	}()

	// 消费者 goroutine
	go func() {
		mutex.Lock() // 获取锁
		for ready == false {
			fmt.Println("等待数据...")
			cond.Wait() // 等待通知，由于cond.Wait()，导致goroutine阻塞在此,直到cond.Signal()被调用,goroutine才会继续执行
		}
		fmt.Printf("数据是: %d\n", data)
		mutex.Unlock() // 解锁
	}()

	// 主 goroutine 等待一段时间以确保其他 goroutines 已经完成
	time.Sleep(5 * time.Second)
}
