package main

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

type Order struct {
	Id     string `json:"id"`     // ID
	Sn     string `json:"sn"`     // 订单编号
	Price  string `json:"price"`  // 订单价格
	Date   string `json:"date"`   // 日期
	Status string `json:"status"` // 状态
}

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	order := Order{
		Id:     "1001",
		Sn:     "sale_2023_0603_1458_8888",
		Price:  "999",
		Date:   "1685772594",
		Status: "2",
	}

	orderJson, err := json.Marshal(order)
	if err != nil {
		fmt.Println(err)
		return
	}

	msg := &sarama.ProducerMessage{
		Topic:     "order",
		Key:       nil,
		Value:     sarama.StringEncoder(orderJson),
		Headers:   nil,
		Metadata:  nil,
		Offset:    0,
		Partition: 0,
		Timestamp: time.Time{},
	}

	client, err := sarama.NewSyncProducer([]string{"192.168.70.31:31919"}, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}
	defer client.Close()

	//发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err:", err)
		return
	}
	fmt.Printf("生产者：pid:%v offset:%v\n", pid, offset)
}
