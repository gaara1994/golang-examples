package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go fun1(ch)

	ch <- 0
	time.Sleep(time.Second * 2)

	ch <- 1
	time.Sleep(time.Second * 2)

	ch <- 2
	ch <- 2
	ch <- 2
	ch <- 2
	ch <- 2
	time.Sleep(time.Second * 2)

	ch <- 3
	time.Sleep(time.Second * 2)

	ch <- 4
	time.Sleep(time.Second * 2)

	ch <- 5
	time.Sleep(time.Second * 2)

	ch <- 6
	time.Sleep(time.Second * 2)

}

func fun1(ch chan int) {
	for {
		tmp := <-ch
		if tmp == 6 {
			fmt.Println("收到的是===", tmp)
		} else {
			fmt.Println("收到的是：", tmp)
		}
	}
}
