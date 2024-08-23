package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func echo() {
	fmt.Println("只打印一次")
}
func main() {
	once.Do(echo)
	once.Do(echo)
	once.Do(echo)
}
