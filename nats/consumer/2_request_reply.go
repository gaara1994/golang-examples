package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

/*
请求-回复模式在以下场景中通常会被使用：

远程过程调用（RPC）：在这种场景中，客户端发送一个请求到服务器，请求执行某个操作，并返回结果。请求-回复模式允许客户端等待并接收服务器的响应。

数据查询：如果你的应用需要查询数据库或其他数据源，你可以使用请求-回复模式。客户端发送包含查询参数的请求，然后等待并接收包含查询结果的回复。

任务分发：如果你的应用需要在多个工作节点之间分发任务，你可以使用请求-回复模式。客户端（通常是一个调度器或协调器）发送包含任务详情的请求，
然后等待并接收工作节点完成任务后的回复。

服务间通信：在微服务架构中，服务之间经常需要进行通信。请求-回复模式允许一个服务向另一个服务发送请求，并等待回复。
*/
func main() {
	// 连接到 NATS 服务器
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// 订阅主题并回复所有请求
	_, err = nc.Subscribe("help", func(m *nats.Msg) {
		err := nc.Publish(m.Reply, []byte("收到"))
		if err != nil {
			log.Fatal(err)
		}
	})
	if err != nil {
		log.Fatal(err)
	}

	// 等待接收请求
	select {}
}
