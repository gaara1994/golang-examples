# pprof基本使用

## 1.安装

```shell
go install github.com/google/pprof@latest
```



## 2.引入程序

```go
package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof" // 导入pprof包，但因为是匿名导入，所以不会自动注册处理器
)

func main() {
	// 正确的地址格式是 ":8081"
	if err := http.ListenAndServe(":8081", nil); err != nil {
		// 处理错误
		fmt.Println("http.ListenAndServe :",err)
	}
	// http://localhost:8081/debug/pprof/
}
```



## 3.启动程序

```shell
go run main.go
```



打开页面：http://localhost:8081/debug/pprof/

```go
/debug/pprof/	#主路由
Set debug=1 as a query parameter to export in legacy text format


Types of profiles available:
//数量   类型
Count	Profile
5	allocs		//显示内存分配情况的剖析数据，反映程序中分配了多少内存以及分配的位置。
0	block		//显示 goroutine 阻塞情况，需要先通过runtime.SetBlockProfileRate设置采样率。
0	cmdline		//提供当前程序的命令行参数信息。
6	goroutine	//显示当前所有goroutine的堆栈信息，有助于分析goroutine的状态和数量。
5	heap		//提供堆内存使用情况的剖析数据，展示内存分配和存活对象的信息。
0	mutex		//显示互斥锁的竞争情况，需要先通过runtime.SetMutexProfileFraction设置采样率。
0	profile		//用于获取CPU使用情况的剖析数据。可以通过查询参数seconds指定采样时长。默认需等待采样30s
8	threadcreate//显示线程创建情况的剖析数据，有助于理解程序的线程模型和负载
0	trace		//提供执行追踪数据，可以用来分析程序的执行流程和阻塞情况，可通过查询参数seconds指定追踪时长。
full goroutine stack dump
```







## 4.获取文件

这里获取的是堆内存的文件。

```shell
curl http://localhost:8080/debug/pprof/heap > heap.out
```

这里获取30s内cpu的文件。

```shell
curl http://localhost:8080/debug/pprof/profile?seconds=30 > cpu.out
```



## 5.分析文件

### top

内存使用情况排序

```shell
➜  base git:(main) ✗ go tool pprof heap.out
File: main
Type: inuse_space
Time: May 16, 2024 at 11:40am (CST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top 10
Showing nodes accounting for 2060.48kB, 100% of 2060.48kB total
Showing top 10 nodes out of 16
      flat  flat%   sum%        cum   cum%
  524.09kB 25.44% 25.44%   524.09kB 25.44%  regexp/syntax.(*compiler).inst
  512.31kB 24.86% 50.30%   512.31kB 24.86%  regexp.onePassCopy
  512.05kB 24.85% 75.15%   512.05kB 24.85%  regexp/syntax.(*Regexp).Simplify
  512.02kB 24.85%   100%   512.02kB 24.85%  gopkg.in/yaml%2ev3.longTag (inline)
         0     0%   100%   512.05kB 24.85%  github.com/go-playground/validator/v10.init
         0     0%   100%   512.31kB 24.86%  github.com/go-playground/validator/v10.init.0
         0     0%   100%   512.02kB 24.85%  gopkg.in/yaml%2ev3.init.1
         0     0%   100%  1024.37kB 49.72%  regexp.Compile (inline)
         0     0%   100%  1024.37kB 49.72%  regexp.MustCompile
         0     0%   100%  1024.37kB 49.72%  regexp.compile

```



### list

查看指定包的的使用情况

```shell
(pprof) list regexp.onePassCopy
Total: 2.01MB
ROUTINE ======================== regexp.onePassCopy in /usr/local/go/src/regexp/onepass.go
  512.31kB   512.31kB (flat, cum) 24.86% of Total
         .          .    222:func onePassCopy(prog *syntax.Prog) *onePassProg {
         .          .    223:   p := &onePassProg{
         .          .    224:           Start:  prog.Start,
         .          .    225:           NumCap: prog.NumCap,
  512.31kB   512.31kB    226:           Inst:   make([]onePassInst, len(prog.Inst)),
         .          .    227:   }
         .          .    228:   for i, inst := range prog.Inst {
         .          .    229:           p.Inst[i] = onePassInst{Inst: inst}
         .          .    230:   }
         .          .    231:
(pprof) 
```



### web

生成连线图，并在浏览器查看图片

```
(pprof) web
[24142:24142:0100/000000.189911:ERROR:zygote_linux.cc(672)] write: 断开的管道 (32)

```



保存图片

```

```



go tool pprof http://localhost:8081/debug/pprof/profile

go tool pprof -http :8083  /home/yantao/pprof/pprof.main.samples.cpu.007.pb.gz







