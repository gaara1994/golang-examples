package main

import (
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//创建一个信号通道
	quitCh := make(chan os.Signal, 1)
	//注册信号处理
	signal.Notify(quitCh, os.Interrupt, syscall.SIGTERM)

	//阻塞等待信号
	<-quitCh

	//处理信号
	sig := <-quitCh
	switch sig {
	case os.Interrupt:
		println("用户终止了程序")
	default:
		println("程序异常终止")
	}
}
