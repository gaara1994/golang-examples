package main

func B() func() int {
	var x = 10 // 这个x就是闭包要捕获的外部局部变量
	return func() int {
		x = x + 5 // 闭包内部必须修改x才会发生逃逸
		return x
	}
}

func main() {
	resB := B() //外部函数B已经执行完毕
	_ = resB    //x仍然保持有效
}
