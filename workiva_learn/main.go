package main

import (
	"fmt"
	"github.com/Workiva/go-datastructures/queue"
)
func main() {
	// 创建一个新的队列
	q := queue.New(10)

	err := q.Put(9)
	if err != nil {
		return
	}
	err = q.Put(9, 5, 2, 7)
	if err != nil {
		return
	}


	for q.Len() > 0 {
		n,err := q.Get(1)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(n)
	}
}
