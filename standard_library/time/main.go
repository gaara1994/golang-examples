package main

import (
	"fmt"
	"time"
)

func main() {
	// 获取当前时间
	now := time.Now()
	fmt.Println("Current Time:", now) //2024-09-11 14:04:26.14300447 +0800 CST m=+0.000013387

	// 格式化时间
	formattedTime := now.Format("2006-01-02 15:04:05")
	fmt.Println("Formatted Time:", formattedTime) //2024-09-11 14:38:58

	fmt.Println("Current Time:", now.Format(time.RFC3339Nano)) // 人类可读时间格式
	fmt.Println("Current Time:", now.Format(time.DateTime))    // 人类可读时间格式

	fmt.Println(time.Now().Unix()) //1726036320

}
