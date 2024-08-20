package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("咱家要出门郊游啦")
	var wg = sync.WaitGroup{}
	wg.Add(3)

	go func() {
		defer wg.Done()
		fmt.Println("大儿子关水完毕")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("二儿子关电完毕")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("三儿子关燃气完毕")
	}()

	wg.Wait()

	fmt.Println("出发啦~~")
}
