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
	kv, err := js.CreateKeyValue(&nats.KeyValueConfig{
		Bucket: "my_bucket",
	})

	// 使用 Watch 方法来监听 "code" 上的新消息
	kw, err := kv.Watch("code")
	if err != nil {
		log.Fatalf("无法设置 Watch: %v", err)
	}
	updates := kw.Updates()

	for update := range updates {
		if update == nil {
			continue
		}
		// 调用 Value 方法来获取键值对的值，并检查是否为 nil
		valueBytes := update.Value()
		if valueBytes == nil {
			log.Println("Received update for key 'code' with nil value")
			continue // 如果 valueBytes 是 nil，则跳过此次循环
		}
		if len(valueBytes) > 0 {
			// 将字节切片转换为字符串（假设值是UTF-8编码的文本）
			valueStr := string(valueBytes)
			log.Printf("Received update for key 'code': %s", valueStr)
		} else {
			log.Println("Received update for key 'code' with empty value")
		}
	}

	// kv put my_bucket code 401
}
