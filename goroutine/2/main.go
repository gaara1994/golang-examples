package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	//给你3秒钟，数数1-100
	go fun1(done)

	time.Sleep(time.Second * 3)
	if <-done == true {
		fmt.Println("真棒！")
	}
}

func fun1(done chan bool) {
	for i := 1; i <= 100; i++ {
		fmt.Println(i)
	}
	done <- true
}
