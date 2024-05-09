package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 直接使用 Background 作为上下文执行操作
	baseCtx := context.Background()

	// 假设我们有一个函数需要一个上下文参数，但不需要取消或超时功能，也不需要传递值
	go simpleTask(baseCtx)

	// 等待一段时间后程序结束，这里仅为演示，实际中可能有更复杂的流程控制
	time.Sleep(2 * time.Second)
}

func simpleTask(ctx context.Context) {
	// 这里我们的任务很简单，只是打印一条消息
	fmt.Println("正在执行一个简单的任务...")
	// 假设有一些耗时操作
	time.Sleep(1 * time.Second)
	fmt.Println("任务完成")
}