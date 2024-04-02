/**
 * @author yantao
 * @date 2024/4/1
 * @description minio接口-连接
 */
package main

import (
	"github.com/minio/minio-go"
	"log"
)

func main() {
	endpoint := "10.233.63.88:9000"
	accessKeyID := "admin"
	secretAccessKey := "password"
	useSSL := false
	// 初使化 minio client对象。
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}
	exists, err := minioClient.BucketExists("my-bucket")
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("my-bucket的状态%#v\n", exists) // minioClient初使化成功
}
