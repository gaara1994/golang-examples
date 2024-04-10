# 1.介绍

Delve 是 Go 编程语言的调试器。该项目的目标是为 Go 提供一个简单、功能齐全的调试工具。

项目地址： https://github.com/go-delve/delve

使用方法： https://github.com/go-delve/delve/tree/master/Documentation/usage

# 2.安装

```shell
go install github.com/go-delve/delve/cmd/dlv@latest
```



> [!TIP]
>
> 这会把dlv安装到 $GOPATH/bin 下，如果切换 root 用户则不能运行 dlv 。可以把dlv添加软链接到 /usr/local/bin 目录下解决。
>
> ```shell
> sudo ln -s $GOPATH/bin/dlv /usr/local/bin/dlv
> ```



# 3.编写程序

```shell
go mod init

vim main.go        
```

```go
package main

import "fmt"

func main() {
	var a = 10
	var b = 7
	var c = a + b

	fmt.Println(c)
}
```



# 4.指定目标并使用默认[终端接口](https://github.com/go-delve/delve/blob/master/Documentation/cli/README.md)开始调试

## 1. [dlv debug](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_debug.md) 

编译并开始调试当前目录或指定的包。

```shell
dlv debug [package] [flags]
```

常用的

| 功能                     | 命令     | 缩写 |
| ------------------------ | -------- | ---- |
| 设置断点                 | break    | b    |
| 继续运行，直到断点处停止 | continue | c    |
| 单步执行，运行下一行     | next     | n    |
| 查看变量值               | print    | p    |
| 查看函数调用栈           | stack    |      |
| 退出调试器               | quit     | q    |
| 打印局部变量             | locals   |      |
| 查看全局变量             | vars     |      |
| 打印函数input            | args     |      |
|                          |          |      |
|                          |          |      |

全部的调试命令：https://github.com/go-delve/delve/blob/master/Documentation/cli/README.md



```shell
yantao@ubuntu20:~/go/src/delve_learn$ dlv debug delve_learn
Type 'help' for list of commands.
(dlv) break main.go:6
Breakpoint 1 set at 0x495bd2 for main.main() ./main.go:6
(dlv) break main.go:7
Breakpoint 2 set at 0x495bdb for main.main() ./main.go:7
(dlv) break main.go:8
Breakpoint 3 set at 0x495be4 for main.main() ./main.go:8
(dlv) c
> main.main() ./main.go:6 (hits goroutine(1):1 total:1) (PC: 0x495bd2)
     1: package main
     2:
     3: import "fmt"
     4:
     5: func main() {
=>   6:         var a = 10
     7:         var b = 7
     8:         var c = a + b
     9:
    10:         fmt.Println(c)
    11:
(dlv) n
> main.main() ./main.go:7 (hits goroutine(1):1 total:1) (PC: 0x495bdb)
     2:
     3: import "fmt"
     4:
     5: func main() {
     6:         var a = 10
=>   7:         var b = 7
     8:         var c = a + b
     9:
    10:         fmt.Println(c)
    11:
    12: }
(dlv) n
> main.main() ./main.go:8 (hits goroutine(1):1 total:1) (PC: 0x495be4)
     3: import "fmt"
     4:
     5: func main() {
     6:         var a = 10
     7:         var b = 7
=>   8:         var c = a + b
     9:
    10:         fmt.Println(c)
    11:
    12: }
(dlv) print a
10
(dlv) print b
7
(dlv) print c
Command failed: could not find symbol value for c
(dlv) n
> main.main() ./main.go:10 (PC: 0x495bed)
     5: func main() {
     6:         var a = 10
     7:         var b = 7
     8:         var c = a + b
     9:
=>  10:         fmt.Println(c)
    11:
    12: }
(dlv) p c
17
(dlv) stack
0  0x0000000000495bed in main.main
   at ./main.go:10
1  0x000000000043a107 in runtime.main
   at /usr/local/go/src/runtime/proc.go:267
2  0x0000000000465781 in runtime.goexit
   at /usr/local/go/src/runtime/asm_amd64.s:1650
(dlv) n
17
> main.main() ./main.go:12 (PC: 0x495c6d)
     7:         var b = 7
     8:         var c = a + b
     9:
    10:         fmt.Println(c)
    11:
=>  12: }
(dlv) q
```





## 2. [dlv exec](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_exec.md) 

执行预编译的二进制文件，并开始调试会话。

由于编译会优化代码和删除行数，所在在编译的时候需要添加参数 `-gcflags "all=-N -l"` 。

> [!WARNING]
>
> ```shell
> go build -gcflags "all=-N -l" -o main main.go
> ```



```shell
yantao@ubuntu20:~/go/src/golang-examples/delve_learn$ dlv exec main
Type 'help' for list of commands.
(dlv) b main.go:6
Breakpoint 1 set at 0x495bd2 for main.main() ./main.go:6
(dlv) b main.go:8
Breakpoint 2 set at 0x495be4 for main.main() ./main.go:8
(dlv) b main.go:9
Breakpoint 3 set at 0x495bed for main.main() ./main.go:9
(dlv) c
> main.main() ./main.go:6 (hits goroutine(1):1 total:1) (PC: 0x495bd2)
     1: package main
     2:
     3: import "fmt"
     4:
     5: func main() {
=>   6:         var a = 10
     7:         var b = 7
     8:         var c = a + b
     9:         fmt.Println(c)
    10: }
    11:
(dlv) n
> main.main() ./main.go:7 (PC: 0x495bdb)
     2:
     3: import "fmt"
     4:
     5: func main() {
     6:         var a = 10
=>   7:         var b = 7
     8:         var c = a + b
     9:         fmt.Println(c)
    10: }
    11:
    12: func Add(a int64, b int64) int64 {
(dlv) n
> main.main() ./main.go:8 (hits goroutine(1):1 total:1) (PC: 0x495be4)
     3: import "fmt"
     4:
     5: func main() {
     6:         var a = 10
     7:         var b = 7
=>   8:         var c = a + b
     9:         fmt.Println(c)
    10: }
    11:
    12: func Add(a int64, b int64) int64 {
    13:         return a + b
(dlv) p b
7
(dlv) n
> main.main() ./main.go:9 (hits goroutine(1):1 total:1) (PC: 0x495bed)
     4:
     5: func main() {
     6:         var a = 10
     7:         var b = 7
     8:         var c = a + b
=>   9:         fmt.Println(c)
    10: }
    11:
    12: func Add(a int64, b int64) int64 {
    13:         return a + b
    14: }
(dlv) n
17
> main.main() ./main.go:10 (PC: 0x495c6d)
     5: func main() {
     6:         var a = 10
     7:         var b = 7
     8:         var c = a + b
     9:         fmt.Println(c)
=>  10: }
    11:
    12: func Add(a int64, b int64) int64 {
    13:         return a + b
    14: }
(dlv) q
```



## 3. [dlv attach](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_attach.md) 

附加到正在运行的进程并开始调试。

首先用 gin 写一个http服务用来测试

```go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

```



编译

```shell
yantao@ubuntu20:~/go/src/gin_demo$ go build -gcflags "all=-N -l"
yantao@ubuntu20:~/go/src/gin_demo$ ls
gin_demo  go.mod  go.sum  main.go
```

启动

```shell
yantao@ubuntu20:~/go/src/gin_demo$ ./gin_demo 
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /                         --> main.main.func1 (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Environment variable PORT is undefined. Using port :8080 by default
[GIN-debug] Listening and serving HTTP on :8080
```



新建终端查询 gin_demo 的 pid，然后调试。

```shell
yantao@ubuntu20:~/go/src/gin_demo$ pidof gin_demo
13925
yantao@ubuntu20:~/go/src/gin_demo$ sudo dlv attach 13925 #需要使用root权限
Type 'help' for list of commands.
(dlv) b main.go:14
Breakpoint 1 set at 0x9210e5 for main.main() ./main.go:14
(dlv) q
Would you like to kill the process? [Y/n] n
```



## 4. [dlv connect](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_connect.md) 

使用终端客户端连接到无外设调试服务器。

前提条件是：启动 Delve 调试会话并监听 9999 端口

```shell
dlv debug --listen=:9999 --headless=true your_program.go
```

这个命令的各个部分有特定的含义：

- `dlv debug`: 指示 Delve 调试器启动一个调试会话。
- `--listen=:9999`: 指定 Delve 调试器应该在哪个端口上监听调试客户端的连接。`:9999` 表示 Delve 将在所有可用的网络接口上监听 9999 端口。如果你只想在本地接口上监听，可以使用 `localhost:9999` 或 `127.0.0.1:9999`。
- `--headless=true`: 指示 Delve 以无头模式运行，即不启动任何图形用户界面（GUI）。
- `your_program.go`: 是你想要调试的 Go 程序的源文件。

它告诉 Delve 编译并运行 `your_program.go` 中的 Go 程序，同时监听 9999 端口以便调试客户端可以连接并进行调试操作。由于没有启动 GUI，所有的调试操作都需要通过命令行进行。

一旦这个命令被执行，Delve 就会开始运行你的 Go 程序，并等待调试客户端的连接。你可以使用另一个命令行窗口运行 `dlv connect 127.0.0.1:9999` 来连接到这个调试会话，并进行调试操作。

下面是操作：

```shell
yantao@ubuntu20:~/go/src/gin_demo$ dlv debug --listen=:9999 --headless=true gin_demo 
API server listening at: [::]:9999
2024-04-17T14:30:11+08:00 warning layer=rpc Listening for remote connections (connections are not authenticated nor encrypted)
```



```shell
yantao@ubuntu20:~$ dlv connect 127.0.0.1:9999
Type 'help' for list of commands.
(dlv) b main.go:14
Breakpoint 1 set at 0x9210e5 for main.main() ./go/src/gin_demo/main.go:14
(dlv) 
```



## 5. [dlv core](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_core.md) 

检查核心转储。



## 6. [dlv DAP](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_dap.md) 

启动通过调试适配器协议 （DAP） 进行通信的无外设 TCP 服务器。



## 7. [dlv replay](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_replay.md) 

重放 rr 跟踪。



## 8. [dlv test](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_test.md)  

编译测试二进制文件并开始调试程序。



## 9. [dlv trace](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_trace.md)  

编译并开始跟踪程序。



## 10. [dlv version](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_version.md)  

打印版本。



## 11. [dlv log](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_log.md)  

 有关日志记录标志的帮助



## 12. [dlv backend](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_backend.md)  







# 跟踪目标程序执行

## 1. [trace ](https://github.com/go-delve/delve/blob/master/Documentation/cli/README.md#trace) 





# 仅启动无头后端服务器

- dlv **--headless** <command> <target> <args>
  - 启动服务器，进入指定目标的调试会话，并等待通过 JSON-RPC 或 DAP 接受客户端连接
  - `<command>`可以是 、 、 、 或 中的任何一项`debug``test``exec``attach``core``replay`
  - 如果未指定标志，则默认[终端客户端](https://github.com/go-delve/delve/blob/master/Documentation/cli/README.md)将自动启动`--headless`
  - 兼容 [dlv connect](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_connect.md)、[VS Code Go](https://github.com/golang/vscode-go/blob/master/docs/debugging.md#remote-debugging)、[GoLand](https://www.jetbrains.com/help/go/attach-to-running-go-processes-with-debugger.html#attach-to-a-process-on-a-remote-machine)
- DLV DAP
  - 启动仅限 DAP 的服务器，并等待 DAP 客户端连接以指定目标和参数
  - 与 [VS Code Go](https://github.com/golang/vscode-go/blob/master/docs/debugging.md#remote-debugging) 兼容
  - 与 [dlv connect](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_connect.md)、[GoLand](https://www.jetbrains.com/help/go/attach-to-running-go-processes-with-debugger.html#attach-to-a-process-on-a-remote-machine) 不兼容
- DLV 连接<ADDR>
  - 启动[终端接口客户端](https://github.com/go-delve/delve/blob/master/Documentation/cli/README.md)，并通过 JSON-RPC 将其连接到正在运行的无头服务器







# 帮助信息

- [dlv help [command\]](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv.md)
- [dlv log](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_log.md)
- [dlv backend](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_backend.md)
- [dlv redirect](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_redirect.md)
- [dlv version](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_version.md)



| 指令    | 用处                                                         | 实操  |
| ------- | ------------------------------------------------------------ | ----- |
| attach  | 这个命令将使Delve控制一个已经运行的进程，并开始一个新的调试会话。 当退出调试会话时，你可以选择让该进程继续运行或杀死它。 | case1 |
| exec    | 这个命令将使Delve执行二进制文件，并立即附加到它，开始一个新的调试会话。请注意，如果二进制文件在编译时没有关闭优化功能，可能很难正确地调试它。请考虑在Go 1.10或更高版本上用-gcflags="all=-N -l "编译调试二进制文件，在Go的早期版本上用-gcflags="-N -l"。 | case2 |
| help    | 使用手册                                                     | case3 |
| debug   | 默认情况下，没有参数，Delve将编译当前目录下的 "main "包，并开始调试。或者，你可以指定一个包的名字，Delve将编译该包，并开始一个新的调试会话。 | case4 |
| test    | test命令允许你在单元测试的背景下开始一个新的调试会话。默认情况下，Delve将调试当前目录下的测试。另外，你可以指定一个包的名称，Delve将在该包中调试测试。双破折号`--`可以用来传递参数给测试程序。 | case5 |
| version | 查看dlv版本                                                  | case6 |





## 2.dlv test

编译测试二进制文件并开始调试程序。