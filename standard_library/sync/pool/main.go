package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type FileHandlerManager struct {
	pool *sync.Pool
}

func NewFileHandlerManager(filename string) *FileHandlerManager {
	return &FileHandlerManager{
		pool: &sync.Pool{
			New: func() interface{} {
				file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
				if err != nil {
					fmt.Println("Error opening file:", err)
					return nil
				}
				return file
			},
		},
	}
}

// GetFileHandle 从池中获取一个文件句柄
func (f *FileHandlerManager) GetFileHandle() *os.File {
	return f.pool.Get().(*os.File)
}

// ReleaseFileHandle 将文件句柄放回池中
func (f *FileHandlerManager) ReleaseFileHandle(file *os.File) {
	if file != nil {
		f.pool.Put(file)
	}
}

func main() {
	filename := "test.log"
	manager := NewFileHandlerManager(filename)

	//写入文件
	for i := 0; i < 10; i++ {
		file := manager.GetFileHandle()
		defer manager.ReleaseFileHandle(file)
		str := fmt.Sprintf("Hello, world %d!\n", i)
		_, err := file.WriteString(str)
		if err != nil {
			fmt.Println("Error writing file:", err)
			return
		}
	}

	//读取文件
	go func() {
		file := manager.GetFileHandle()
		defer manager.ReleaseFileHandle(file)
		buffer := make([]byte, 1024)
		n, err := file.Read(buffer)
		if err != nil {
			fmt.Println("Error reading from file:", err)
			return
		}
		content := buffer[:n]                        // 截取实际读取到的字节数
		fmt.Printf("Content: %s\n", string(content)) // 使用 string() 将字节切片转换为字符串
	}()

	time.Sleep(2 * time.Second)
}
