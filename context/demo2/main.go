package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//创建一个根上下文，通常在main函数或者初始化函数中创建
	ctx := context.Background()

	//使用WithCancel从父上下文创建一个可以取消的上下文
	ctx,cancle := context.WithCancel(ctx)

	//启动一个goroutine，把新的上下文传递进去，让goroutine监听这个ctx，并相应的退出
	go doSomething(ctx)

	//等待3秒之后关闭上下文
	time.Sleep(3 * time.Second)
	cancle()

	// 等待一段时间让goroutine有机会退出
	time.Sleep(2 * time.Second)
}

func doSomething(ctx context.Context)  {
	//通过select 监听上下文的Done通道
	for{
		select {
		case <-ctx.Done():
			//当上下文取消时，这里会接收到信号
			fmt.Println("接收到信号停止运行")
			return
		default:
			fmt.Println("工作中...")
			time.Sleep(1 * time.Second)
		}
	}
}