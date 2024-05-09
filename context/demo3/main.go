package main

import (
	"context"
	"fmt"
	"time"
)

// simulateLongRunningOperation 模拟一个可能运行时间较长的操作
func simulateLongRunningOperation(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Println("操作正在进行...")
		case <-ctx.Done():
			if ctx.Err() == context.DeadlineExceeded {
				fmt.Println("由于超过最后期限，操作被取消.",ctx.Err())
			} else if ctx.Err() == context.Canceled {
				fmt.Println("操作已取消:", ctx.Err())
			}
			return
		}
		// 注意：这里不再直接比较i和duration
	}
	// fmt.Println("Operation completed successfully.") 这行现在是多余的，因为我们通过context控制退出了
}

func main() {
	// 设置一个具体的截止时间，比如当前时间后3秒
	deadline := time.Now().Add(3 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	go simulateLongRunningOperation(ctx)

	// 让main函数等待一段时间，确保上面的goroutine有足够时间执行
	time.Sleep(10 * time.Second)
}