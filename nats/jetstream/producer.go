package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	// 连接到 NATS 服务器
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// 创建 JetStream 上下文
	js, err := nc.JetStream()
	if err != nil {
		log.Fatal(err)
	}

	// 创建流
	_, err = js.AddStream(&nats.StreamConfig{
		Name:     "MY_STREAM",
		Subjects: []string{"ORDERS.*"},
	})
	if err != nil {
		log.Fatal(err)
	}

	// 发布消息
	ack, err := js.Publish("ORDERS.scratch", []byte("order-123"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(ack)
}
