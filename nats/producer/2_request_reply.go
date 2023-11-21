package main

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	// 连接到 NATS 服务器
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// 发送请求并等待回复
	// 如果在10秒内没有收到回复，nc.Request函数就会返回一个错误
	msg, err := nc.Request("help", []byte("辅助请跟我"), 10*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Received a reply: %s\n", string(msg.Data))
}
