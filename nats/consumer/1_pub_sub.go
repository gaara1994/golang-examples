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

	// 订阅主题
	_, err = nc.Subscribe("updates", func(m *nats.Msg) {
		log.Printf("Received a message: %s\n", string(m.Data))
	})
	if err != nil {
		log.Fatal(err)
	}

	// 等待接收消息
	select {}
}
