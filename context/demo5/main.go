package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 设置超时时间为2秒
	parentCtx := context.Background()

	ctx := context.WithValue(parentCtx,"user_id","10010")

	getUserName(ctx)

	// 让main函数等待一段时间，确保上面的goroutine有足够时间执行
	time.Sleep(5 * time.Second)
}

func getUserName(ctx context.Context) {
	id := ctx.Value("user_id")
	fmt.Println("用户的id是：",id)
}