package main

import (
	"context"
	"fmt"
	"log"

	"github.com/olivere/elastic/v7"
)

func main() {
	// 创建一个新的 Elasticsearch 客户端
	client, err := elastic.NewClient(elastic.SetURL("http://192.168.68.117:30200"), elastic.SetSniff(false))
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// 检查 Elasticsearch 服务器的状态
	info, code, err := client.Ping("http://192.168.68.117:30200").Do(context.Background())
	if err != nil {
		log.Fatalf("Error pinging Elasticsearch server: %s", err)
	}
	fmt.Printf("Elasticsearch 连接成功， code %d and version %s\n", code, info.Version.Number)

	// 获取一个文档，其中 kubernetes.pod.name = abc
	// 假设你的索引名称为 "k8slog*"
	searchResult, err := client.Search().
		Index("k8slog*").
		Query(elastic.NewTermQuery("kubernetes.pod.name", "mxschejob-vm6v9-fdwx4")).
		From(0).Size(10).
		Pretty(true).
		Do(context.Background())

	if err != nil {
		log.Fatalf("Error getting the document: %s", err)
	}

	// 输出搜索结果中的命中数和文档信息
	totalHits := searchResult.Hits.TotalHits.Value
	fmt.Printf("Found a total of %d documents:\n", totalHits)
	for _, hit := range searchResult.Hits.Hits {
		fmt.Println(hit.Source)
	}

}
