package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

// 启动多个该消费者
/**
在NATS中，当多个消费者订阅同一个主题，并且他们都属于同一个队列组时，NATS服务器会确保每条消息只被队列组中的一个消费者接收。

具体来说，当一条消息到达服务器时，服务器会查看所有订阅了该消息主题的队列组，并从每个队列组中选择一个消费者来接收该消息。这个选择过程通常是轮询的，
这意味着如果有多个消费者在同一个队列组中，他们会依次接收到消息，从而实现负载均衡。

这种机制允许你在多个消费者之间分发消息负载，同时确保每条消息只被处理一次。这对于需要大规模并行处理的应用来说非常有用，例如大数据处理、微服务架构等。
如果你有任何问题，或者需要进一步的帮助，随时告诉我！
*/
func main() {
	// 连接到 NATS 服务器
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// 订阅主题
	_, err = nc.QueueSubscribe("updates", "worker_group", func(m *nats.Msg) {
		log.Printf("Received a message: %s\n", string(m.Data))
	})
	if err != nil {
		log.Fatal(err)
	}

	// 等待接收消息
	select {}
}
