/**
 * @author yantao
 * @date 2024/4/1
 * @description minio接口-连接
 * 安装方式 https://min.io/docs/minio/container/index.html
 * 官方文档有错误 https://min.io/docs/minio/linux/developers/go/minio-go.html
 */
package main

import (
	"github.com/minio/minio-go"
	"log"
)

func main() {
	endpoint := "127.0.0.1:9000"
	accessKeyID := "ROOTNAME"
	secretAccessKey := "CHANGEME123"
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%#v\n", minioClient) // minioClient is now set up
}
