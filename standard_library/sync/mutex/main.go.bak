package main

import (
	"fmt"
	"sync"
	"time"
)

func func1() {
	var count int
	for i := 0; i < 1000; i++ {
		count++
	}
	fmt.Println("count1:", count)
}

func func2() {
	var count int
	for i := 0; i < 1000; i++ {
		go func() {
			count++
		}()
	}

	time.Sleep(time.Second * 3)
	fmt.Println("count2:", count)
}

func func3() {
	var count int
	var mu sync.Mutex
	for i := 0; i < 1000; i++ {
		go func() {
			mu.Lock()
			count++
			mu.Unlock()
		}()
	}
	time.Sleep(time.Second * 3)
	fmt.Println("count3:", count)
}
func main() {
	func1()

	func2()

	func3()

}
