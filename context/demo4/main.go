package main

import (
	"context"
	"fmt"
	"time"
)

func simulateTask(ctx context.Context) {
	for {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("运行中...")
		case <-ctx.Done():
			if ctx.Err() == context.DeadlineExceeded {
				fmt.Println("由于超时而结束.")
			} else {
				fmt.Println("任务被关闭:", ctx.Err())
			}
			return
		}
	}
}

func main() {
	// 设置超时时间为2秒
	parentCtx := context.Background()
	ctx, cancel := context.WithTimeout(parentCtx, 2*time.Second)
	defer cancel() // 确保上下文在函数结束时被取消

	go simulateTask(ctx)

	// 让main函数等待一段时间，确保上面的goroutine有足够时间执行
	time.Sleep(5 * time.Second)
}