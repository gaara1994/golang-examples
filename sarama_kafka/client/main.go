package main

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
)

type Order struct {
	Id     string `json:"id"`     // ID
	Sn     string `json:"sn"`     // 订单编号
	Price  string `json:"price"`  // 订单价格
	Date   string `json:"date"`   // 日期
	Status string `json:"status"` // 状态
}

func main() {
	consumer, err := sarama.NewConsumer([]string{"192.168.70.31:31919"}, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return
	}

	partitionList, err := consumer.Partitions("order") // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}
	fmt.Println(partitionList)

	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition("order", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return
		}
		defer pc.Close()

		// 异步从每个分区消费信息

		for msg := range pc.Messages() {
			fmt.Printf("消费者：Partition:%d Offset:%d Key:%v Value:%v\n", msg.Partition, msg.Offset, msg.Key, string(msg.Value))
			var order Order
			json.Unmarshal(msg.Value, &order)
			fmt.Println(order)
		}

	}
	select {}
}
