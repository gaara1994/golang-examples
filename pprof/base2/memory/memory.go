package memory

import "fmt"

type Memory struct {
}

func (m *Memory) MakeOneGbSlice() {
	fmt.Println("开始创建切片")
	size := 1024 * 1024 * 1024
	slice := make([]byte, size)

	for i := range slice {
		slice[i] = 1
	}
	fmt.Println("创建切片结束")
}
