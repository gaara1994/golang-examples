package main

import (
	"github.com/nats-io/nats.go"
	"log"
)

/*
NATS JetStream是一种基于发布/订阅模式的消息传递系统，核心原理是提供持久化的NATS服务。它通过发布者发布消息到主题(Topic)，然后订阅者订阅感兴趣的主题，从而实现消息的传递和处理。

在NATS JetStream中，消息是通过流(Stream)的方式进行处理的。流是一个持久化的消息日志，可以存储大量的消息，并且支持消息的顺序订阅和回溯。
发布者将消息发布到特定的流中，订阅者可以从流中订阅消息，并按照发布的顺序接收和处理消息。这种方式可以确保消息的可靠传递和顺序处理。

与传统的消息队列系统不同，NATS JetStream具有更高的性能和可扩展性。它使用流来处理和存储消息，可以轻松地实现高吞吐量和可扩展性。
此外，NATS JetStream还提供了持久化和可靠的消息传递，即使在服务器故障的情况下也能保证消息的不丢失。
*/

/*
NATS JetStream模式具有以下特点：

1.高性能：NATS JetStream具有低延迟和高吞吐量的特点，可以满足实时数据处理的需求，并能够处理大量的消息，同时保持良好的性能表现。
2.可扩展性：NATS JetStream可以通过增加节点来扩展系统的容量和吞吐量，以满足不断增长的业务需求。
3.持久性和可靠性：NATS JetStream支持持久化和可靠的消息传递，即使在服务器故障的情况下也能保证消息的不丢失。
4.顺序消息：NATS JetStream支持消息的顺序订阅和回溯，可以按照发布的顺序接收和处理消息。
5.灵活的订阅方式：NATS JetStream支持多种订阅方式，包括拉取（pull）和推送（push）订阅，可以满足不同业务场景下的需求。
6.轻量级：NATS JetStream是一种轻量级的消息传递系统，相对于传统的消息队列系统来说，它的资源占用更少，更易于部署和管理。
*/
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
