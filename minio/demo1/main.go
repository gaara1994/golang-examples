/**
 * @author yantao
 * @date 2024/4/1
 * @description minio接口-操作存储桶
 */
package main

import (
	"log"

	"github.com/minio/minio-go"
)

var Client minio.Client

func init() {
	endpoint := "192.168.68.112:9000"
	accessKeyID := "admin"
	secretAccessKey := "adminehpv2"
	useSSL := false
	// 初使化 minio client对象。
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}
	Client = *minioClient
	log.Println("minioClient init success: ", Client)
}
func main() {
	// 判断存储桶是否存在
	// BucketExists()
	// ObjectExists()
	ListObjects()
}

func BucketExists() {
	var bucketName = "ehpdata"
	exists, err := Client.BucketExists(bucketName)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("bucket exists:", exists)
	//ehpdata/ehp/zytest01/db_v2_edit/onemap/33633217/557040086
}

func ObjectExists() {
	// 检查的桶名和对象名
	bucketName := "ehpdata"
	objectName := "ehp/zytest01/db_v2_edit/onemap/33633217/557040086/CNode.geojson"

	// 检查对象是否存在
	_, err := Client.StatObject(bucketName, objectName, minio.StatObjectOptions{})
	if err == nil {
		log.Printf("Object '%s' exists in bucket '%s'.\n", objectName, bucketName)
	} else if minio.ToErrorResponse(err).Code == "NoSuchKey" {
		log.Printf("Object '%s' does not exist in bucket '%s'.\n", objectName, bucketName)
	} else {
		log.Fatalf("Failed to check object: %v\n", err)
	}
}

func ListObjects() {
	// 检查的桶名和对象名
	bucketName := "ehpdata"
	objectName := "ehp/zytest01/db_v2_edit/onemap/33633217/557040086/CNode.geojson"
	objectsCh := Client.ListObjects(bucketName, objectName, true, nil)

	// 检查是否存在任何对象
	var found bool
	for object := range objectsCh {
		if object.Err != nil {
			log.Fatalln(object.Err)
		}
		found = true
		break
	}
	log.Default().Println("found ==", found)
}
