/**
 * @author yantao
 * @date 2024/4/1
 * @description minio接口-操作存储桶
 */
package main

import (
	"context"
	"fmt"
	"github.com/minio/minio-go"
	"log"
	"os"
	"time"
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

	//1.创建一个存储桶。
	//err = minioClient.MakeBucket("mybucket", "us-east-1")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println("mybucket创建成功")

	//2.列出所有的存储桶。
	buckets, err := minioClient.ListBuckets()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, bucket := range buckets {
		fmt.Println(bucket)
	}

	//3.检查存储桶是否存在。
	found, err := minioClient.BucketExists("mybucket")
	if err != nil {
		fmt.Println(err)
		return
	}
	if found {
		fmt.Println("Bucket 存在")
	}

	//4.删除一个存储桶，存储桶里面必须为空才能被成功删除。
	//if found {
	//	err = minioClient.RemoveBucket("mybucket")
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	} else {
	//		fmt.Println("mybucket 已经删除成功")
	//	}
	//}

	/*
		5.上传文件
		当对象小于64MiB时，直接在一次PUT请求里进行上传。
		当大于64MiB时，根据文件的实际大小，PutObject会自动地将对象进行拆分成64MiB一块或更大一些进行上传。
		对象的最大大小是5TB。
	*/
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	file, err := os.Open("../minio.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fileStat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	n, err := minioClient.PutObjectWithContext(ctx, "mybucket", "a/b/minio.txt", file, fileStat.Size(), minio.PutObjectOptions{
		ContentType: "application/octet-stream",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("上传成功 bytes: ", n)

	//6.列举存储桶中的对象
	// Create a done channel to control 'ListObjectsV2' go routine.
	doneCh := make(chan struct{})
	// Indicate to our routine to exit cleanly upon return.
	defer close(doneCh)
	isRecursive := true
	objectCh := minioClient.ListObjectsV2("mybucket", "", isRecursive, doneCh)
	for object := range objectCh {
		if object.Err != nil {
			fmt.Println(object.Err)
			return
		}
		fmt.Println(object.Size)
		fmt.Println(object.Key)
	}

	//7.下载并将文件保存到本地文件系统。
	err = minioClient.FGetObject("mybucket", "a/b/minio.txt", "a.txt", minio.GetObjectOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}

}
