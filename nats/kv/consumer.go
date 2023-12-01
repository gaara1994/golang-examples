package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
)

func main() {
	// 连接到NATS服务器
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}

	// 创建JetStream上下文
	js, err := nc.JetStream()
	if err != nil {
		log.Fatal(err)
	}

	// 创建一个键/值存储桶
	kv, err := js.CreateKeyValue(&nats.KeyValueConfig{
		Bucket: "MY_BUCKET",
	})
	if err != nil {
		log.Fatal(err)
	}

	// 在桶中放置一个键
	_, err = kv.PutString("foo", "bar")
	if err != nil {
		log.Fatal(err)
	}

	// 获取键的值
	val, err := kv.Get("foo")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Got value:", string(val.Value()))

	// 删除键
	err = kv.Delete("foo")
	if err != nil {
		log.Fatal(err)
	}
}
