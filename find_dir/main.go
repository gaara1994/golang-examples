package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	// 输出的文件名
	var OutPutFile = "dir_list_by_go_16-17.txt"

	// 修改后的开始和结束时间
	var StartTime = time.Date(2024, 3, 16, 0, 0, 0, 0, time.UTC)
	var EndTime = time.Date(2024, 3, 20, 0, 0, 0, 0, time.UTC)

	cmd := exec.Command("pwd")  // 创建一个表示"pwd"命令的对象
	output, err := cmd.Output() // 执行命令并获取输出

	if err != nil {
		fmt.Println("Error executing command:", err)
		return
	}

	// 由于输出是字节切片，需要转换为字符串并去除末尾的换行符
	pwd := strings.TrimSpace(string(output))
	fmt.Println("查找指定时间范围内最后一次修改的文件所在的目录")
	fmt.Println("当前目录为:", pwd)
	fmt.Println("文件最后一次修改的时间范围：", StartTime, "到", EndTime)

	// 初始化 checkedDirs map
	checkedDirs := make(map[string]bool) // 使用map存储已检查过的目录

	// 创建输出文件并准备写入
	outputFile, err := os.Create(OutPutFile)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 检查是否为目录并且已经找到了符合条件的文件
		if info.IsDir() && checkedDirs[path] {
			return filepath.SkipDir // 跳过当前目录下的其他文件和子目录
		}

		if !info.Mode().IsRegular() { // 只处理普通文件
			return nil
		}

		modTime := info.ModTime()
		if modTime.After(StartTime) && modTime.Before(EndTime) {
			// 将符合条件的目录路径写入输出文件
			if _, writeErr := outputFile.WriteString(filepath.Dir(path) + "\n"); writeErr != nil {
				fmt.Println("Error writing to output file:", writeErr)
				return writeErr
			}

			checkedDirs[filepath.Dir(path)] = true // 标记该目录已检查并找到符合条件的文件
			return filepath.SkipDir                // 跳过当前目录下的其他文件和子目录
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error walking the path:", err)
	}
	fmt.Println("执行结束，输出文件为：", OutPutFile)
}
