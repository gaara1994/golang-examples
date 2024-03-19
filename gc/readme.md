# 内存逃逸

参考文章

fmt 系列的打印函数会让参数“逃逸”   https://juejin.cn/post/7067084509132357645 

通过实例理解go逃逸分析  https://tonybai.com/2021/05/24/understand-go-escape-analysis-by-example/

https://juejin.cn/post/6992178559208914957

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

备注：由于声明变量必须使用，而`fmt`系列的打印会导致内存逃逸，干扰结果。ps:研究了一下午为什么`a`逃逸了（TAT）

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

**闭包是由函数和其相关的引用环境组合而成的实体**，当这个函数a被定义在另一个函数B内部，并且该内部函数a引用了外部函数B的局部变量时，即使外部函数B已经执行完毕，这些局部变量B.x的值仍然会被内部函数a（即闭包）所捕获并保持有效。

```go
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
	_ = resB    //B.x仍然保持有效
}
```

```shell
yantao@ubuntu20:~/go/src/golang-examples/gc$ go build -gcflags="-l -m" test4.go 
# command-line-arguments
./test4.go:4:6: moved to heap: x
./test4.go:5:9: func literal escapes to heap
```

在实际运行时，即使你没有立即调用`resB`，变量`x`也会因为闭包的存在而逃逸到堆上。



## 3.大对象或未知大小对象

1.尺寸可能超过函数栈空间限制的变量。

2.尺寸在编译时无法确定的对象。

```go
package main

func createLargeArray(size int) []int {
    // 这里创建了一个大小由运行时传入参数决定的切片
    return make([]int, size)
}

func main() {
    KnownSize := make([]int, 9) //可知大小：小（未逃逸）
    _ = KnownSize
    UnknownSize := make([]int, 999999) //可知大小：大（逃逸）
    _ = UnknownSize

    bigArray := createLargeArray(8) //不可知大小：无论大小（逃逸）
    _ = bigArray
}
```

```shell
yantao@ubuntu20:~/go/src/golang-examples/gc$ go build -gcflags="-l -m" test5.go 
# command-line-arguments
./test5.go:5:13: make([]int, size) escapes to heap
./test5.go:9:19: make([]int, 9) does not escape
./test5.go:11:21: make([]int, 999999) escapes to heap

```

另外，像字符串、字节数组等，如果它们的长度在编译期间未知并且实际创建时长度较大，也会被分配到堆上。



## 4.interface类型逃逸

在前面我们提到过`fmt`系列的打印会导致内存逃逸。现在就来看看为什么吧。

```go
package main

import "fmt"

func main() {
	str := "Hello"
	fmt.Println(str)
}
```

```shell
yantao@ubuntu20:~/go/src/golang-examples/gc$ go build -gcflags="-l -m" test6.go 
# command-line-arguments
./test6.go:7:13: ... argument does not escape
./test6.go:7:14: str escapes to heap #str逃逸到堆内存上
```

看一下`fmt.Println()`源代码

```go
func Println(a ...any) (n int, err error) {
	return Fprintln(os.Stdout, a...)
}
```

`Println`的形参是一个`interface{}`类型，

## 5.并发编程