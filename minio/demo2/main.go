/**
 * @author yantao
 * @date 2024/4/1
 * @description minio接口-操作存储桶
 */
package main

import (
	"fmt"
	"github.com/minio/minio-go"
	"log"
)

func main() {
	endpoint := "127.0.0.1:9000"
	accessKeyID := "ROOTNAME"
	secretAccessKey := "CHANGEME123"
	useSSL := false
	// 初使化 minio client对象。
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}

	//2.列出所有的存储桶。
	buckets, err := minioClient.ListBuckets()
	if err != nil {
		log.Fatal(err)
	}

	for _, bucket := range buckets {
		fmt.Println(bucket)
	}

	// Make a new bucket called testbucket.
	bucketName := "testbucket2"
	location := "us-east-1" //分布式minio的所在区域

	//3.检查存储桶是否存在。
	found, err := minioClient.BucketExists(bucketName)
	if err != nil {
		log.Fatalln(err)
	}

	if found {
		log.Fatalln(bucketName, "已经存在")
	}

	//创建bucket
	err = minioClient.MakeBucket(bucketName, location)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("创建成功", bucketName)

	//5.删除bucket
	err = minioClient.RemoveBucket(bucketName)
	if err != nil {
		log.Println(err)
	}
	log.Println("创建删除", bucketName)
}
