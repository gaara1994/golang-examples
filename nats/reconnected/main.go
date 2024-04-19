package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

// 示例回调函数，处理接收到的消息
func msgHandler(msg *nats.Msg) {
	fmt.Printf("Received a message: %s\n", string(msg.Data))
}

func main() {
	// 创建一个连接选项，设置重连策略
	opts := nats.Options{
		Url:            "nats://192.168.246.128:8222", // 替换为实际的NATS服务器地址
		AllowReconnect: true,
		MaxReconnect:   10,
		ReconnectWait:  5 * time.Second, // 设置重连间隔
		Timeout:        2 * time.Second, // 设置连接超时时间

		DisconnectedCB: func(conn *nats.Conn) {
			log.Println("Disconnected from server, trying to reconnect...")
		},
		ReconnectedCB: func(conn *nats.Conn) {
			log.Println("Reconnected to server.")
			// 在这里可以重新订阅之前的主题，具体取决于业务需求
		},
	}

	// 建立与NATS服务器的连接
	nc, err := opts.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to NATS server: %v", err)
	}
	defer nc.Close()

	// 订阅主题并设置消息处理器
	sub, err := nc.Subscribe("my.subject", msgHandler)
	if err != nil {
		log.Fatalf("Failed to subscribe: %v", err)
	}
	defer sub.Unsubscribe() // 确保在退出前取消订阅

	// 在主循环中等待，让客户端保持运行
	for {
		time.Sleep(time.Minute)
	}

	// 如果网络不稳定或客户端意外断开连接，由于设置了ReconnectedCB回调，
	// 当客户端重新连接时，可以根据需要重新订阅主题或清理资源。
}
