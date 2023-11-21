package main

import (
	"github.com/nats-io/nats.go"
	"log"
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

	// 创建消费者
	_, err = js.AddConsumer("ORDERS", &nats.ConsumerConfig{
		Durable: "my_consumer",
	})
	if err != nil {
		log.Fatal(err)
	}

	// 订阅消息
	sub, err := js.SubscribeSync("ORDERS.scratch")
	if err != nil {
		log.Fatal(err)
	}

	// 处理消息
	msg, err := sub.NextMsg(nats.DefaultTimeout)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Received a message: %s\n", string(msg.Data))
}
