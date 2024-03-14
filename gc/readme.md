# 内存逃逸

参考文章

fmt 系列的打印函数会让参数“逃逸”   https://juejin.cn/post/7067084509132357645 

通过实例理解go逃逸分析  https://tonybai.com/2021/05/24/understand-go-escape-analysis-by-example/

如何避免逃逸发生  https://www.zhihu.com/question/592999770?write

在Go语言中，内存逃逸（Memory Escape）通常发生在 **变量生命周期超出了其定义的作用域**，导致变量被分配到堆上而不是栈上。这通常发生在将局部变量传递给闭包、跨函数边界传递指针、或者当局部变量被外部引用时。

可以在编译时候 添加参数 `-gcflags="-m"`来查看。

```shell
go build -gcflags="-l -m" main.go
```

使用 `go build -gcflags="-l -m" main.go` 命令编译 Go 程序时，您将同时实现两个目标：

- `-l` 参数会禁用函数内联优化。这意味着编译器不会尝试将函数体插入到调用它的位置，而是保留正常的函数调用。也可以不适用`-l`，但是我们为了减少无关的信息输出，使用了`-l`。

- `-m` 参数会输出逃逸分析信息和一些编译器的中间表示决策，包括为什么某些函数没有被内联的原因等。

这样编译程序后，您不仅关闭了内联优化，还能查看编译过程中有关变量逃逸、内联决定等方面的详细信息。这对于理解编译器如何处理代码以及调试性能问题非常有用。

以下是Go语言中可能导致内存逃逸的一些常见场景：

## 1.返回局部变量的指针

`返回局部变量的指针`强调的是**指针**，那么我们先试一试 **返回值**会不会逃逸。

 **返回值：**

备注：由于声明变量必须使用，而fmt.系列的打印会导致内存逃逸，干扰结果。ps:研究了一下午为什么`a`逃逸了（TAT）

```go
package main

func fun1() int {
	x := 10
	return x
}

func main() {
	var a = 5
	_ = a

	resX := fun1()
	_ = resX
}

```

没有输出，说明没有逃逸。

```shell
yantao@ubuntu20:~/go/src/golang-examples/gc$ go build -gcflags="-l -m" test1.go 
yantao@ubuntu20:~/go/src/golang-examples/gc$ 
```



**指针：**

```go
package main

func fun2() *int {
	x := 10
	return &x
}

func main() {
	resX := fun2()
	_ = resX
}
```

```shell
yantao@ubuntu20:~/go/src/golang-examples/gc$ go build -gcflags="-l -m" test2.go 
# command-line-arguments
./test2.go:4:2: moved to heap: x
```

看到x被移动到堆内存。

函数 `fun2()` 返回了局部变量 `x` 的地址。由于 `x` 是一个局部变量，在函数返回后其作用域将结束，但是通过返回指针的方式，该变量的生命周期需要延长到函数调用者可以访问它的时候。因此，Go 编译器会执行逃逸分析，并决定将变量 `x` 从栈上移动到堆上以确保即使在 `fun2()` 函数返回后，`resX` 指向的内容依然有效。



**怎么避免？**

返回局部变量的指针	改为：

返回全局变量的指针

```go
package main

var x3 int

func fun3() *int {
	x3 = 3
	return &x3
}

func main() {
	resX := fun3()
	_ = resX
}

```

```shell
yantao@ubuntu20:~/go/src/golang-examples/gc$ go build -gcflags="-l -m" test3.go 
yantao@ubuntu20:~/go/src/golang-examples/gc$ 
```

函数 `fun3()` 直接修改全局变量 `x3` 的值，并返回它的地址。在这种情况下，**由于没有涉及局部变量的存储位置改变或生命周期延长**，所以编译时不会触发内存逃逸。





## 2.闭包引用





## 3.大对象或可变大小对象





## 4.动态类型





## 5.并发编程