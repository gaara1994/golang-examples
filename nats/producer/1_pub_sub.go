package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/nats-io/nats.go"
)

func main() {
	// 连接到 NATS 服务器
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// 发布消息
	for i := 0; i < 100; i++ {
		str1 := "这是第"
		str2 := "条消息"

		var builder strings.Builder
		builder.WriteString(str1)
		builder.WriteString(strconv.Itoa(i))
		builder.WriteString(str2)
		result := builder.String()

		err = nc.Publish("updates", []byte(result))
		if err != nil {
			log.Fatal(err)
		}
	}
}
