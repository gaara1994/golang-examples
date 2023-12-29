package main

import (
	"context"
	"fmt"
	"log"

	"github.com/olivere/elastic/v7"
)

func main() {
	// 连接 Elasticsearch
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:92001"))
	if err != nil {
		panic(err)
	}

	// 创建一个新的索引
	// _, err = client.CreateIndex("my_index").Do(context.Background())
	// if err != nil {
	// 	panic(err)
	// }

	// 插入一条新的文档
	// _, err = client.Index().Index("my_index").BodyJson(map[string]interface{}{
	// 	"title":   "这是一条日志标题 blog",
	// 	"content": "这是一条日志的具体内容 blog.",
	// }).Do(context.Background())
	// if err != nil {
	// 	panic(err)
	// }

	// 查询文档
	res, err := client.Search().Index("my_index").Query(elastic.NewMatchQuery("title", "blog")).Do(context.Background())
	if err != nil {
		log.Println("err=", err.Error())
	}

	// 输出查询结果
	fmt.Printf("Query took %d milliseconds\n", res.TookInMillis)
	fmt.Printf("Found a total of %d documents\n", res.TotalHits())
	for _, hit := range res.Hits.Hits {
		fmt.Printf(" * %s\n", hit.Source)
	}
}
